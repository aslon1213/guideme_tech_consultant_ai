import chromadb
import json
import orjson
from format_with_chatgpt import format_actions_sequence
from chromadb.utils import embedding_functions
from chromadb.db.base import UniqueConstraintError
import dotenv


def load_envs():
    dotenv.load_dotenv(verbose=True)


class OpenedActionsFormatter:
    def __init__(
        self,
    ):
        self.collection = None

    def GetChroma(self, username):
        chroma = chromadb.PersistentClient("./data_actions/" + username)
        collection = chroma.get_collection("demo_collection")
        self.collection = collection

    def Query(self, query_string):
        results = self.collection.query(query_texts=[query_string], n_results=1)
        return results

    def Format(self, results):
        results = results["documents"][0]
        results = orjson.dumps(results)
        return results

    def FormatToGeneralAnswer(self, results, user_message):
        o = orjson.loads(results["documents"][0][0])
        message, response = format_actions_sequence(o, user_message=user_message)
        if "action" in message.content[:10]:
            return orjson.loads(message.content[8:])
        return orjson.loads(message.content)


class TrainActionsBot:
    def __init__(self):
        self.actions = []
        self.username = None

    def SetActions(self, actions):
        self.actions = actions

    def AddAction(self, action):
        print("Got action", action)
        self.actions.append(action)

    def SetUsername(self, username):
        self.username = username

    def TrainandSave(self, sentence_transformer_ef):
        chroma_client = chromadb.PersistentClient("./data_actions/" + self.username)

        try:
            collection = chroma_client.create_collection(
                "demo_collection", embedding_function=sentence_transformer_ef
            )

        except UniqueConstraintError:
            print("UniqueConstraintError error occured - making new collection")
            chroma_client.delete_collection("demo_collection")
            collection = chroma_client.create_collection(
                "demo_collection", embedding_function=sentence_transformer_ef
            )
            print("Deleted data from collection")
            # da = collection.query(query_texts=["hello"], n_results=1000)["documents"]
            # print(da)
        print("Added data to collection")
        for i in range(len(self.actions)):
            properties_dict = {x: "" for x in self.actions[i]["properties"]}
            self.actions[i]["properties"] = properties_dict
        print(self.actions)
        ids = [str(i) for i in range(len(self.actions))]

        collection.add(documents=[json.dumps(i) for i in self.actions], ids=ids)
        return "done"
