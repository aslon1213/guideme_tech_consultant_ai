import json

with open("test.txt", "r") as f:
    text = f.read()
    couples = text.split("\n\n")
    print(couples)
    print(len(couples))
    questions = []
    for couple in couples:
        q = couple.split("\n")[0]
        a = "".join(couple.split("\n")[1:])
        questions.append({"question": q, "answer": a})
    print(questions)
    with open("test.json", "w") as f1:
        json.dump(questions, f1)
