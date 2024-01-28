# -*- encoding: utf-8 -*-
"""
Copyright (c) 2019 - present AppSeed.us
"""

from django.urls import path, re_path
from apps.home import views

urlpatterns = [
    # The home page
    path("", views.index, name="home"),
    # Matches any html file
    # re_path(, views.pages, name="pages"),
    path("actions/train", views.actionstrain, name="actions_train"),
    path("documents/upload", views.uploaddocument, name="upload_document"),
    path("documents/train", views.traindocuments, name="tune"),
    # chat paths
    path("chat/open", views.open_chat, name="open_chat"),
    path("chat/query", views.query_chat, name="query_chat"),
    path("chat/close", views.close_chat, name="close_chat"),
]
