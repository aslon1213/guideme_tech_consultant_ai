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

    def FormatToGeneralAnswer(self, results):
        o = orjson.loads(results["documents"][0][0])
        if o.get("can_be_formatted", False):
            message, response = format_actions_sequence(o, o["user_message"])
            o["actions"] = orjson.loads(message.content[14:])
        o["type"] = "actions"
        return o


class TrainActionsBot:
    def __init__(self):
        self.actions = []
        self.username = None

    def SetActions(self, actions):
        self.actions = actions

    def AddAction(self, action):
        self.actions.append(action)

    def SetUsername(self, username):
        self.username = username

    def TrainandSave(self):
        chroma_client = chromadb.PersistentClient("./data_actions/" + self.username)
        sentence_transformer_ef = (
            embedding_functions.SentenceTransformerEmbeddingFunction(
                model_name="all-MiniLM-L6-v2"
            )
        )
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
            ids = [str(i) for i in range(len(self.actions))]
            collection.add(documents=self.actions, ids=ids)
            print("Added data to collection")
            # da = collection.query(query_texts=["hello"], n_results=1000)["documents"]
            # print(da)

        collection.add(
            documents=self.actions, ids=[str(i) for i in range(len(self.actions))]
        )
        return "done"
