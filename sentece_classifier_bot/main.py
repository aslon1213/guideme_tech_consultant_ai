import time
from load_model import load_model, preprocess_sentence

batch_size = 64

model_being_loaded = load_model()


# def do_query_on_chroma(model, query):
#     preprocesed_sentece = preprocess_sentence(query)
#     predictions = model.predict(preprocesed_sentece, batch_size=batch_size, verbose=1)
#     test = []
#     for i in range(0, len(predictions)):
#         test.append(predictions[i].argmax(axis=0))
#     return test


def main(model, sentence: str):
    preprocesed_sentece = preprocess_sentence(sentence)
    predictions = model.predict(preprocesed_sentece, batch_size=batch_size, verbose=1)
    test = []
    for i in range(0, len(predictions)):
        test.append(predictions[i].argmax(axis=0))
    return test


# if __name__ == "__main__":
#     while True:
#         try:
#             i = input("Enter sentence: ").strip()
#             now = time.time()
#             p = main(model_being_loaded, i)
#             print(p)
#             types = {"0": " ", "question": 1, "statement": 2, "command": 3}
#             for t in types.keys():
#                 if p[0] == types[t]:
#                     print("Text: ", i)
#                     print("Predicted type: ", t)
#                     print("\nTime taken to run:", time.time() - now, "seconds")
#                     break
#         except Exception as e:
#             print(e)
#             continue
