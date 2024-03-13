from main_pb2 import (
    TrainResponse,
    ActionFull,
    GeneralAnswer,
    ChatID,
    TrainResponse,
    AudoWithText,
)
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
from chatbot import Chatbot, TrainonDocuments, AnotherException
from langchain_community.embeddings import HuggingFaceEmbeddings

sentence_transformer_ef_chroma = (
    embedding_functions.SentenceTransformerEmbeddingFunction(
        model_name="all-MiniLM-L6-v2"
    )
)
# chat_id = [actions, chatbot]
opened_chats = {}

model = load_model()
print("Model loaded!")

sentence_transformer_ef_huggingface = HuggingFaceEmbeddings(
    model_name="all-MiniLM-L6-v2"
)

############################################################################################################
############################################################################################################
# user for a testing purposes - it will be deleted ofc immidiately before production
user_data = {
    "id": {"$oid": "65bb9447fc13ae5136234b42"},
    "first_name": "Mekhriddin",
    "last_name": "Jumaev",
    "email": "mekhriddinjumaev1@bloglines.com",
    "gender": "Male",
    "passport_number": "AB1234567",
    "money": 1000000,
    "credit_left": {
        "amount": 120000,
        "currency": "soum",
        "should_be_paid": "11.02.2024",
    },
    "if_not_paid_on_time_policy": "You will be charged fine with the amount equal to 1 minimum wage in uzbekistan after 15 days of the deadline - also it can drastically damage your credit history",
    "credit_payment_history": [
        {"date": "11.11.2023", "amount": 100000, "currency": "soum"},
        {"date": "11.12.2024", "amount": 100000, "currency": "soum"},
        {"date": "11.01.2024", "amount": 100000, "currency": "soum"},
    ],
}
############################################################################################################
############################################################################################################


def handle_chat_query(chatbot: Chatbot, query, needs_formatting=True):
    results = chatbot.Query(query)
    if needs_formatting:
        return chatbot.FormatForGeneralAnswer(results)
    else:
        return results


class ToClassifierServicer(main_pb2_grpc.ToClassifierServicer):
    """Provides methods that implement functionality of route guide server."""

    def __init__(self):
        self.text = "Hello"

    def GiveAudioAnswerOrJustTextAnswer(self, request, context):
        # do something to get the asnwer for the query, question, sentence etc
        query = request.query
        # username = request.username
        chat_id = request.chat_id
        chatbot = opened_chats[chat_id][1]

        print("Doing query with audio or text, chat_id = ", chat_id)
        results = handle_chat_query(chatbot, query, needs_formatting=False)
        # convert
        return AudoWithText(text=results)

    def GetGreetingMessage(self, request, context):
        username = request.username
        chat_id = request.chat_id
        chatbot = opened_chats[chat_id][1]
        print("Getting greeting message")
        answer = chatbot.GetGreetingMessage()
        # convert answer to bytes

        return GeneralAnswer(answer=str.encode(answer))

    def SaveDocuments(self, request, context):
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
        chroma_instance = training_.TrainDocumentsAlsoJson(
            questions_answer,
            sentence_transformer_ef=sentence_transformer_ef_huggingface,
        )
        # chroma_instance.add_documents(questions_answer)
        print("Training done")
        return TrainResponse(message="Training is being done")

    def TrainOnSavedDocuments(self, request, context):
        username = request.username
        training_ = TrainonDocuments()
        training_.SetUsername(username)
        training_.LoadDocuments()
        training_.Train(sentence_transformer_ef=sentence_transformer_ef_huggingface)
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
                    "type": action_full["type"],
                    "description": action_full.get("description", ""),
                    "deeplink": action_full.get("deeplink", ""),
                }
            )
            counter += 1
        training_.SetUsername(username)
        training_.TrainandSave(sentence_transformer_ef=sentence_transformer_ef_chroma)
        return TrainResponse(message="Training done")

    def DeleteDocument(self, request, context):
        # yet to be implemented
        return super().DeleteDocument(request, context)

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
        chatbot.LoadVectorstoreandChatbot(username, user_data=user_data)
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
        p = [3]
        if "I want to" not in request.query:
            p = main(model, q)
        types = {"0": " ", "question": 1, "statement": 2, "command": 3}
        for t in types.keys():
            if types[t] == p[0]:
                print(
                    "Expected sentence classification ",
                    t,
                    " ---- with type number",
                    p,
                )
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
            formatted_results, open_ai_response = (
                actions_formatter.FormatToGeneralAnswer(results, request.query)
            )
            # parse results to ActionFull
            print("Got results", formatted_results)
            formatted_results["type"] = "actions"
            response_2 = GeneralAnswer(
                answer=orjson.dumps(formatted_results),
                input_token_length=open_ai_response.usage.prompt_tokens,
                output_token_length=open_ai_response.usage.completion_tokens,
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
            try:
                results = handle_chat_query(chatbot, q)
            except AnotherException:
                print(
                    "seems like question is actually is actions request - doing additional actions thing"
                )
                actions_formatter = opened_chats[request.chat_id][0]
                results_2 = actions_formatter.collection.query(
                    query_texts=[request.query], n_results=1
                )
                formatted_results, open_ai_response = (
                    actions_formatter.FormatToGeneralAnswer(results_2, request.query)
                )
                # parse results to ActionFull
                print("Got results", formatted_results)
                formatted_results["type"] = "actions"
                return GeneralAnswer(answer=orjson.dumps(formatted_results))
            # print("Got results", results)
            response_2 = GeneralAnswer(
                answer=results,
                input_token_length=open_ai_response.usage.prompt_tokens,
                output_token_length=open_ai_response.usage.completion_tokens,
            )
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
            description=o["description"],
            deeplink=o["deeplink"],
            username=request.username,
            type=o["type"],
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
