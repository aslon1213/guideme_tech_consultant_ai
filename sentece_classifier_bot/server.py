from main_pb2 import TrainResponse, ActionFull, GeneralAnswer, ChatID, TrainResponse
import main_pb2_grpc
import grpc
from concurrent import futures
import logging
import math
import time
import chromadb
from chromadb.db.base import UniqueConstraintError
from chromadb.utils import embedding_functions
from google.protobuf.json_format import MessageToDict
import orjson
from load_model import load_model
from main import main
from uuid import uuid4
import initializers
from chatbot import Chatbot, TrainonDocuments

# chat_id = [actions, chatbot]
opened_chats = {}

model = load_model()
print("Model loaded!")


def handle_chat_query(chatbot: Chatbot, query):
    results = chatbot.Query(query)
    return chatbot.FormatForGeneralAnswer(results)


class ToClassifierServicer(main_pb2_grpc.ToClassifierServicer):
    """Provides methods that implement functionality of route guide server."""

    def __init__(self):
        self.text = "Hello"

    def SaveDocuments(self, request, context):
        username = ""
        username = request.username
        with open("./data_chatbot/" + username + "/" + request.filename, "w") as f:
            f.write(str(request.file_content))
        return TrainResponse(message="Documents saved")

    def TrainonSavedDocumentsJson(self, request, context):
        questions_answer = MessageToDict(request)["qa"]
        print(questions_answer)
        username = request.username
        training_ = TrainonDocuments()
        training_.SetUsername(username)
        training_.LoadDocuments()
        chroma_instance = training_.TrainDocumentsAlsoJson(questions_answer)
        # chroma_instance.add_documents(questions_answer)
        print("Training done")
        return TrainResponse(message="Training is being done")

    def TrainOnSavedDocuments(self, request, context):
        username = request.username
        training_ = TrainonDocuments()
        training_.SetUsername(username)
        training_.LoadDocuments()
        training_.Train()
        print("Training done")
        return TrainResponse(message="Training is being done")

    # def Train(self, request_iterator, context):
    #     username = ""
    #     counter = 0
    #     training_ = TrainonDocuments()
    #     for request in request_iterator:
    #         if counter == 0:
    #             username = request.username
    #             training_.SetUsername = username
    #         training_.AddDocument(request.document, request.file_type, request.filename)
    #         counter += 1
    #     trained_model = training_.Train()
    #     trained_model.save("./data_chatbot" + username)

    #     return TrainResponse(message="Training done")

    def TrainActions(self, request_iterator, context):
        username = ""
        counter = 0
        training_ = initializers.TrainActionsBot()
        for request in request_iterator:
            if counter == 0:
                username = request.username
            action_full = MessageToDict(request)
            print(action_full)
            training_.AddAction(
                {
                    "properties": action_full.get("properties", []),
                    "name": action_full["name"],
                }
            )

            counter += 1
        training_.SetUsername(username)
        training_.TrainandSave()
        return TrainResponse(message="Training done")

    def OpenChat(self, request, context):
        username = request.username
        chat_id = str(uuid4())
        # get actions from chroma
        actions_formatter = initializers.OpenedActionsFormatter()
        try:
            actions_formatter.GetChroma(username)
        except ValueError:
            pass

        # collection = actions_formatter.collection
        # get chatbot
        chatbot = Chatbot()
        chatbot.LoadVectorstoreandChatbot(username)
        opened_chats[chat_id] = [actions_formatter, chatbot]
        return ChatID(chat_id=chat_id)

    def CloseChat(self, request, context):
        chat_id = request.chat_id

        try:
            del opened_chats[chat_id]
        except KeyError:
            return GeneralAnswer(
                answer=orjson.dumps({"message": "Chat already closed"})
            )
        answer = orjson.dumps(
            {"chat_id": chat_id, "message": "Chat closed --- take care !!!!"}
        )
        return GeneralAnswer(answer=answer)

    def ClassifyAndAnswer(self, request, context):
        print(request)
        q = request.query
        p = main(model, q)
        # types = {"0": " ", "question": 1, "statement": 2, "command": 3}
        print(p)
        if p[0] == 3:
            # go to query actions
            print("I am goint with actions")
            try:
                actions_formatter = opened_chats[request.chat_id][0]
            except KeyError:
                return GeneralAnswer(
                    answer=orjson.dumps({"message": "Chat With this chat_id not found"})
                )
            results = actions_formatter.collection.query(
                query_texts=[request.query], n_results=1
            )
            formatted_results = actions_formatter.FormatToGeneralAnswer(
                results, request.query
            )
            # parse results to ActionFull
            print("Got results", formatted_results)
            response_2 = GeneralAnswer(
                answer=orjson.dumps({"message": formatted_results, "type": "actions"})
            )
        else:
            #  go to chatbot
            print("I am going to chatbot")
            try:
                chatbot = opened_chats[request.chat_id][1]
            except KeyError:
                return GeneralAnswer(
                    answer=orjson.dumps({"message": "Chat With this chat_id not found"})
                )
            results = handle_chat_query(chatbot, q)
            print("Got results", results)
            response_2 = GeneralAnswer(answer=results)
        return response_2

    def QueryActions(self, request, context):
        print(request)
        chroma = chromadb.PersistentClient("./data_actions/" + request.username)
        collection = chroma.get_collection("demo_collection")

        # results = collection.query(query_texts=["hello"], n_results=1)["documents"]
        results = collection.query(query_texts=[request.query], n_results=1)
        # parse results to ActionFull
        o = orjson.loads(results["documents"][0][0])
        print(o)
        return ActionFull(
            properties=o["properties"],
            username=request.username,
            name=o["name"],
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    main_pb2_grpc.add_ToClassifierServicer_to_server(ToClassifierServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Server started at port 50051")
    server.wait_for_termination()


if __name__ == "__main__":
    print("Starting server")
    initializers.load_envs()
    logging.basicConfig()
    serve()
