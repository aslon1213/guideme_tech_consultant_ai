import os

from load_model import load_model

from typing import Union

from fastapi import FastAPI

app = FastAPI()
model = load_model()
print("Model loaded!")


@app.get("/")
def read_root():
    return {"Hello": "World"}


from main import main


@app.get("/classify")
def read_item(q: Union[str, None] = None):
    p = main(model, q)
    types = {"0": " ", "question": 1, "statement": 2, "command": 3}
    for t in types.keys():
        if p[0] == types[t]:
            print("Text: ", q)
            print("Predicted type: ", t)
            return {"text": q, "type": t}
    return {"text": q, "type": "unknown"}
