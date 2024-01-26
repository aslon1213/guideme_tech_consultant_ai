import tensorflow as tf
from keras.models import Sequential, model_from_json
from sentence_types import encode_data, import_embedding
from sentence_types import load_encoded_data
from keras.preprocessing.sequence import pad_sequences

maxlen = 500
batch_size = 64


def preprocess_sentence(sss):
    test_comments = [sss]
    test_comments_category = ["command"]
    embedding_name = "data/default"
    pos_tags_flag = True
    x_test, _, y_test, _ = encode_data(
        test_comments,
        test_comments_category,
        data_split=1.0,
        embedding_name=embedding_name,
        add_pos_tags_flag=pos_tags_flag,
    )
    x_test = pad_sequences(x_test, maxlen=maxlen)
    return x_test


def load_model():
    model_name = "cnn"
    print("Loading model!")

    # load json and create model
    json_file = open("./trained/" + model_name + ".json", "r")
    loaded_model_json = json_file.read()
    json_file.close()
    model = model_from_json(loaded_model_json)

    # load weights into new model
    model.load_weights("./trained/" + model_name + ".h5")
    print("Loaded model from disk")

    # evaluate loaded model on test data
    model.compile(
        loss="categorical_crossentropy", optimizer="adam", metrics=["accuracy"]
    )
    print("Model compiled!")
    return model
