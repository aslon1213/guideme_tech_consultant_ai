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
from langchain_openai import ChatOpenAI
from langchain.prompts import PromptTemplate

# from langchain_core.documents import Document
from langchain_community.embeddings import HuggingFaceEmbeddings
import chromadb


class AnotherException(Exception):
    pass


# Define your custom prompt template
custom_prompt_template = """You are a helpful assistant that provides answers to user questions using the provided context and additional knowledge when necessary.

Context: {context}, {user_data}



Helpful Answer:"""

custom_prompt = PromptTemplate(
    input_variables=["context", "user_data" "question"],
    template=custom_prompt_template,
)


class Chatbot:
    def __init__(self):
        self.chatbot = None
        self.vectorstore = None

    def GetGreetingMessage(self):
        return "Hello, I am a smart assistant of Agrobank. I am created to be helpful to you. You have a payment for a loan that should be paid by the eleventh of this month for a amount of \u0022120000\u0022 soums. Do you have any question ?"

    def Query(self, query_string):
        results = self.chatbot.invoke(query_string)
        print("Got results from openai: ", results)
        return results

    def LoadVectorstoreandChatbot(self, username, user_data: dict):
        # if os.getenv("OPENAI_API_KEY") != None:
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
        user_data_str = json.dumps(user_data)
        user_data_str = user_data_str.replace("{", "")
        user_data_str = user_data_str.replace("}", "")
        custom_prompt_template = (
            """You are a helpful assistant that provides answers to user questions using the provided context and additional knowledge when necessary. Do not dictate the information about passport_number and id if question asks about it, just asnwer I cant this information. Try to be nice a bit and answer in a way that user can understand. If output contains date then convert date to word format like 11th November 2023 to eleventh November two thousand and twenty threeth year. If asked about who you are, then answer I am agrobank's smart assistant.
        
        Context: {context} ----- 
        user data =  + """
            + user_data_str
            + """
        ---------
        Question: {question}

        Helpful Answer:"""
        )

        custom_prompt = PromptTemplate(
            input_variables=["context", "question"],
            template=custom_prompt_template,
        )

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
        prompt = custom_prompt
        llm = ChatOpenAI(
            model_name="gpt-4",
            temperature=0,
        )
        # gpt-4-turbo-preview
        # gpt -3.5-turbo
        # gpt-4-0125-preview
        user_data_json = json.dumps(user_data)
        rag_chain = (
            {
                "context": retriever | format_docs,
                "question": RunnablePassthrough(),
            }
            | prompt
            | llm
            | StrOutputParser()
        )

        self.chatbot = rag_chain
        self.vectorstore = vectorstore

    def FormatForGeneralAnswer(self, result):
        # print("Message from openai: ", result)
        if "I'm sorry" in result:
            raise AnotherException()
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

    def TrainDocumentsAlsoJson(self, qa_list, sentence_transformer_ef):
        text_splitter = RecursiveCharacterTextSplitter(
            chunk_size=1000, chunk_overlap=200
        )
        splits = text_splitter.split_documents(documents=self.documents)
        print("Splits ----- ", splits)
        chromaaa = chromadb.PersistentClient("./data_chatbot/" + self.username)
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

    def Train(self, sentence_transformer_ef):
        text_splitter = RecursiveCharacterTextSplitter(
            chunk_size=1000, chunk_overlap=200
        )
        splits = text_splitter.split_documents(documents=self.documents)

        chromaaa = chromadb.PersistentClient("./data_chatbot/" + self.username)
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
