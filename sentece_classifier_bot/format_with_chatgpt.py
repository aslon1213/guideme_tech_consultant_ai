import openai
import os

openai.api_key = os.getenv("OPENAI_KEY")
import json


def format_actions_sequence(actions: dict, user_message: str = None):
    messages = [
        {
            "role": "system",
            "content": "You are bot and you have only one job: extract data from the user message and fill the empty values in the properties with the corresponding values you got from the user message. If properties cannot be filled - do not change anything -simply output. do not add any prefix to the output.",
        },
        {
            "role": "user",
            "content": user_message,
        },
    ]
    if user_message == None:
        raise Exception("User message is required")
    actions_string = ""
    actions_string = json.dumps(actions)
    messages.append({"role": "system", "content": "action: " + actions_string})

    response = openai.chat.completions.create(model="gpt-4", messages=messages)
    # usage = response.usage
    message = response.choices[0].message
    print("GET MESSAGE: ", message)
    return message, response


# import time

# if __name__ == "__main__":
#     now = time.time()
#     actions = {
#         "name": "send money to someone",
#         "actions": [
#             {"type": "click", "element": "transfers"},
#             {"type": "select", "element": "writing space"},
#             {"type": "input", "element": "card number", "value": ""},
#             {"type": "input", "element": "amount", "value": ""},
#             {"type": "click", "element": "next button"},
#             {"type": "choose", "element": "card from which to send"},
#             {"type": "click", "element": "send button"},
#         ],
#     }
#     user_message = "I want to send my 100000 som to card with number 8600122344552122"
#     message, response = format_actions_sequence(actions, user_message)
#     # message to message.json
#     with open("message.json", "w") as f:
#         f.write(message.content)
#     print(message.content)
#     print("\nTime taken to format actions:", time.time() - now, "seconds")
