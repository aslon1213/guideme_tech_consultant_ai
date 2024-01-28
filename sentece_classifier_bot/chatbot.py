# import chromadb
# from chromadb.utils import embedding_functions
# from huggingface_hub import InferenceClient
# from sentence_transformers import SentenceTransformer
import os
import orjson
from sentence_transformers import SentenceTransformer
import json

# import bs4
from langchain import hub
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain_community.document_loaders import TextLoader
from langchain_community.document_loaders.directory import DirectoryLoader
from langchain_community.vectorstores import Chroma
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import RunnablePassthrough
from langchain_openai import ChatOpenAI, OpenAIEmbeddings
from langchain_core.documents import Document
from langchain_community.embeddings import HuggingFaceEmbeddings
import chromadb


class Chatbot:
    def __init__(self):
        self.chatbot = None
        self.vectorstore = None

    def Query(self, query_string):
        results = self.chatbot.invoke(query_string)
        print("Got results from openai: ", results)
        return results

    def LoadVectorstoreandChatbot(self, username):
        os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_KEY")
        os.environ["LANGCHAIN_TRACING_V2"] = "false"
        # os.environ["LANGCHAIN_API_KEY"] = ""

        # Load, chunk and index the contents of the blog.
        # loader = TextLoader(file_path="./data_uzbek.txt")
        # docs = loader.load()
        # text_splitter = RecursiveCharacterTextSplitter(
        #     chunk_size=1000, chunk_overlap=200
        # )
        # splits = text_splitter.split_documents(docs)

        def format_docs(docs):
            return "\n\n".join(doc.page_content for doc in docs)

        # load chroma
        # Retrieve and generate using the relevant snippets of the blog.
        chromaaa = chromadb.PersistentClient("./data_chatbot/" + username)
        vectorstore = Chroma(
            collection_name="demo_collection",
            client=chromaaa,
            embedding_function=HuggingFaceEmbeddings(model_name="all-MiniLM-L6-v2"),
        )
        retriever = vectorstore.as_retriever()
        prompt = hub.pull("rlm/rag-prompt")
        llm = ChatOpenAI(model_name="gpt-4", temperature=0)
        # gpt-4-turbo-preview
        # gpt -3.5-turbo
        # gpt-4-0125-preview
        rag_chain = (
            {"context": retriever | format_docs, "question": RunnablePassthrough()}
            | prompt
            | llm
            | StrOutputParser()
        )

        self.chatbot = rag_chain
        self.vectorstore = vectorstore

    def FormatForGeneralAnswer(self, result):
        print(result)
        return orjson.dumps({"type": "answer", "message": result})

    def LoadVectorStore(self):
        pass


class TrainonDocuments:
    def __init__(self) -> None:
        self.documents = []
        self.username = None
        self.file_types = []
        self.filenames = []

    def LoadDocuments(self):
        if self.username != None:
            loader = DirectoryLoader(
                "./data_chatbot/" + self.username,
                glob="**/*.txt",
            )
            print(loader.path)
            docs = loader.load()
            self.documents = docs
            # directory loader from langchain

    def SetUsername(self, username):
        self.username = username

    def TrainDocumentsAlsoJson(self, qa_list):
        text_splitter = RecursiveCharacterTextSplitter(
            chunk_size=1000, chunk_overlap=200
        )
        splits = text_splitter.split_documents(documents=self.documents)
        print("Splits ----- ", splits)
        chromaaa = chromadb.PersistentClient("./data_chatbot/" + self.username)
        sentence_transformer_ef = HuggingFaceEmbeddings(model_name="all-MiniLM-L6-v2")
        chroma_instance = Chroma(
            collection_name="demo_collection",
            client=chromaaa,
            embedding_function=sentence_transformer_ef,
        )
        print("adding documents")
        chroma_instance.add_documents(splits)
        qa_list_str = [json.dumps(i) for i in qa_list]
        print("QA LIST __------ ", qa_list)
        collection = chromaaa.get_collection("demo_collection")
        # get collection's ids
        n = collection.count()
        collection.add(
            documents=qa_list_str, ids=[str(i + n) for i in range(len(qa_list))]
        )
        return chroma_instance

    def Train(self):
        text_splitter = RecursiveCharacterTextSplitter(
            chunk_size=1000, chunk_overlap=200
        )
        splits = text_splitter.split_documents(documents=self.documents)

        chromaaa = chromadb.PersistentClient("./data_chatbot/" + self.username)
        sentence_transformer_ef = HuggingFaceEmbeddings(model_name="all-MiniLM-L6-v2")
        chroma_instance = Chroma(
            collection_name="demo_collection",
            client=chromaaa,
            embedding_function=sentence_transformer_ef,
        )
        print("adding documents")
        chroma_instance.add_documents(splits)

        return chroma_instance


# Path: sentence-classification/sentece_classifier_bot/initializers.p

# if __name__ == "__main__":
#     train = TrainonDocuments("test")
#     print("Training started")
#     train.SetUsername("test")

#     train.LoadDocuments()

#     os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_KEY")
#     os.environ["LANGCHAIN_TRACING_V2"] = "false"
#     print("Training is done")
#     # chroma_instance = Chroma(collection_name="demo_collection", client=chromaaaa)
#     retriever = train.Train()

#     def format_docs(docs):
#         return "\n\n".join(doc.page_content for doc in docs)

#     prompt = hub.pull("rlm/rag-prompt")
#     llm = ChatOpenAI(model_name="gpt-3.5-turbo", temperature=0)

#     rag_chain = (
#         {"context": retriever | format_docs, "question": RunnablePassthrough()}
#         | prompt
#         | llm
#         | StrOutputParser()
#     )
#     print("asking question")
#     response = rag_chain.invoke("What is this document about ?")
#     print(response)
#     pass
