# -*- encoding: utf-8 -*-
"""
Copyright (c) 2019 - present AppSeed.us
"""

from django import template
from django.contrib.auth.decorators import login_required
from django.http import HttpResponse, HttpResponseRedirect, HttpRequest, JsonResponse
from django.template import loader
from django.urls import reverse

# import message
from django.contrib import messages


def home_page(request):
    context = {}
    template = loader.get_template("home/index.html")
    return HttpResponse(template.render(context=context))


def open_chat(request):
    username = request.GET.get("username")
    print(username)
    return HttpResponseRedirect(f"http://localhost:9000/chat/open?username={username}")


def query_chat(request):
    q = request.GET.get("q")
    chat_id = request.GET.get("chat_id")
    return HttpResponseRedirect(
        f"http://localhost:9000/chat/query?", params={"q": q, "chat_id": chat_id}
    )


def close_chat(request):
    chat_id = request.GET.get("chat_id")
    return HttpResponseRedirect(f"http://localhost:9000/chat/close?chat_id={chat_id}")


@login_required(login_url="/login/")
def index(request):
    context = {"segment": "index"}

    html_template = loader.get_template("home/index.html")
    return HttpResponse(html_template.render(context, request))


@login_required(login_url="/login/")
def pages(request):
    context = {}
    # All resource paths end in .html.
    # Pick out the html file name from the url. And load that template.
    try:
        load_template = request.path.split("/")[-1]

        if load_template == "admin":
            return HttpResponseRedirect(reverse("admin:index"))
        context["segment"] = load_template

        html_template = loader.get_template("home/" + load_template)
        return HttpResponse(html_template.render(context, request))

    except template.TemplateDoesNotExist:
        html_template = loader.get_template("home/page-404.html")
        return HttpResponse(html_template.render(context, request))

    except:
        html_template = loader.get_template("home/page-500.html")
        return HttpResponse(html_template.render(context, request))


import json

# from requests import Request
import requests
from django import forms


class myForm(forms.Form):
    jsonfield = forms.JSONField()
    username = forms.CharField()

    def clean_jsonfield(self):
        jdata = self.cleaned_data["jsonfield"]
        try:
            json_data = json.loads(jdata)  # loads string as json
            # validate json_data
        except:
            raise forms.ValidationError("Invalid data in jsonfield")
        # if json data not valid:
        # raise forms.ValidationError("Invalid data in jsonfield")
        return jdata


def actionstrain(request: HttpRequest):
    context = {}

    if request.method == "POST":
        form = myForm(request.POST)
        # get username query param
        username = form.data["username"]
        data = form.data["jsonfield"]
        try:
            user_input = json.loads(data)
        except:
            messages.warning(request, "Invalid data in jsonfield")
            return HttpResponseRedirect(reverse("actions"))

        res = requests.put(
            f"http://localhost:9000/actions/train?username={username}",
            data=data,
            headers={"Content-Type": "application/json"},
        )
        print(res.content)
    form = myForm()
    context["form"] = form
    html_template = loader.get_template("home/ui-tables.html")
    return HttpResponse(html_template.render(context, request))


class FormForFineTUning(forms.Form):
    username = forms.CharField()


def traindocuments(request: HttpRequest):
    context = {}
    if request.method == "POST":
        form = FormForFineTUning(request.POST)
        # get username query param
        username = form.data["username"]
        res = requests.get(
            f"http://localhost:9000/documents/train?username={username}",
            headers={"Content-Type": "application/json"},
        )
        print(res.content)
        messages.success(request, res.content.decode("utf-8"))
        return JsonResponse(res.content.decode("utf-8"), safe=False)
    form = FormForFineTUning()
    context["form"] = form
    html_template = loader.get_template("home/ui-tables.html")
    return HttpResponse(html_template.render(context, request))


def uploaddocument(request: HttpRequest):
    context = {}
    # print("actions")
    # print(request.headers)
    if request.method == "POST":
        # get file from request
        print(request.FILES)
        try:
            file = request.FILES["myfile"]
            # get form data
            username = request.POST["username"]
        except:
            messages.warning(request, "Invalid data in form")
            return HttpResponseRedirect(reverse("documents"))

        print(file.name)
        print(username)

        response = requests.post(
            f"http://localhost:9000/documents/upload",
            files={"file": file},
            params={"username": username},
        )
        print(response.content)
        messages.success(request, response.content.decode("utf-8"))

    html_template = loader.get_template("home/ui-typography.html")
    return HttpResponse(html_template.render(context, request))
