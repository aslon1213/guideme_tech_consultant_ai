# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import main_pb2 as main__pb2


class ToClassifierStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.TrainOnSavedDocuments = channel.unary_unary(
                '/toclassifier.ToClassifier/TrainOnSavedDocuments',
                request_serializer=main__pb2.Username.SerializeToString,
                response_deserializer=main__pb2.TrainResponse.FromString,
                )
        self.TrainActions = channel.stream_unary(
                '/toclassifier.ToClassifier/TrainActions',
                request_serializer=main__pb2.ActionFull.SerializeToString,
                response_deserializer=main__pb2.TrainResponse.FromString,
                )
        self.TrainonSavedDocumentsJson = channel.unary_unary(
                '/toclassifier.ToClassifier/TrainonSavedDocumentsJson',
                request_serializer=main__pb2.JsonData.SerializeToString,
                response_deserializer=main__pb2.TrainResponse.FromString,
                )
        self.QueryActions = channel.unary_unary(
                '/toclassifier.ToClassifier/QueryActions',
                request_serializer=main__pb2.Query.SerializeToString,
                response_deserializer=main__pb2.ActionFull.FromString,
                )
        self.SaveDocuments = channel.unary_unary(
                '/toclassifier.ToClassifier/SaveDocuments',
                request_serializer=main__pb2.Documents.SerializeToString,
                response_deserializer=main__pb2.GeneralAnswer.FromString,
                )
        self.ClassifyAndAnswer = channel.unary_unary(
                '/toclassifier.ToClassifier/ClassifyAndAnswer',
                request_serializer=main__pb2.Query.SerializeToString,
                response_deserializer=main__pb2.GeneralAnswer.FromString,
                )
        self.OpenChat = channel.unary_unary(
                '/toclassifier.ToClassifier/OpenChat',
                request_serializer=main__pb2.Query.SerializeToString,
                response_deserializer=main__pb2.ChatID.FromString,
                )
        self.CloseChat = channel.unary_unary(
                '/toclassifier.ToClassifier/CloseChat',
                request_serializer=main__pb2.ChatID.SerializeToString,
                response_deserializer=main__pb2.GeneralAnswer.FromString,
                )


class ToClassifierServicer(object):
    """Missing associated documentation comment in .proto file."""

    def TrainOnSavedDocuments(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TrainActions(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TrainonSavedDocumentsJson(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryActions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SaveDocuments(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ClassifyAndAnswer(self, request, context):
        """to classifier service
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def OpenChat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CloseChat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ToClassifierServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'TrainOnSavedDocuments': grpc.unary_unary_rpc_method_handler(
                    servicer.TrainOnSavedDocuments,
                    request_deserializer=main__pb2.Username.FromString,
                    response_serializer=main__pb2.TrainResponse.SerializeToString,
            ),
            'TrainActions': grpc.stream_unary_rpc_method_handler(
                    servicer.TrainActions,
                    request_deserializer=main__pb2.ActionFull.FromString,
                    response_serializer=main__pb2.TrainResponse.SerializeToString,
            ),
            'TrainonSavedDocumentsJson': grpc.unary_unary_rpc_method_handler(
                    servicer.TrainonSavedDocumentsJson,
                    request_deserializer=main__pb2.JsonData.FromString,
                    response_serializer=main__pb2.TrainResponse.SerializeToString,
            ),
            'QueryActions': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryActions,
                    request_deserializer=main__pb2.Query.FromString,
                    response_serializer=main__pb2.ActionFull.SerializeToString,
            ),
            'SaveDocuments': grpc.unary_unary_rpc_method_handler(
                    servicer.SaveDocuments,
                    request_deserializer=main__pb2.Documents.FromString,
                    response_serializer=main__pb2.GeneralAnswer.SerializeToString,
            ),
            'ClassifyAndAnswer': grpc.unary_unary_rpc_method_handler(
                    servicer.ClassifyAndAnswer,
                    request_deserializer=main__pb2.Query.FromString,
                    response_serializer=main__pb2.GeneralAnswer.SerializeToString,
            ),
            'OpenChat': grpc.unary_unary_rpc_method_handler(
                    servicer.OpenChat,
                    request_deserializer=main__pb2.Query.FromString,
                    response_serializer=main__pb2.ChatID.SerializeToString,
            ),
            'CloseChat': grpc.unary_unary_rpc_method_handler(
                    servicer.CloseChat,
                    request_deserializer=main__pb2.ChatID.FromString,
                    response_serializer=main__pb2.GeneralAnswer.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'toclassifier.ToClassifier', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ToClassifier(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def TrainOnSavedDocuments(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/TrainOnSavedDocuments',
            main__pb2.Username.SerializeToString,
            main__pb2.TrainResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TrainActions(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/toclassifier.ToClassifier/TrainActions',
            main__pb2.ActionFull.SerializeToString,
            main__pb2.TrainResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TrainonSavedDocumentsJson(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/TrainonSavedDocumentsJson',
            main__pb2.JsonData.SerializeToString,
            main__pb2.TrainResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryActions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/QueryActions',
            main__pb2.Query.SerializeToString,
            main__pb2.ActionFull.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SaveDocuments(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/SaveDocuments',
            main__pb2.Documents.SerializeToString,
            main__pb2.GeneralAnswer.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ClassifyAndAnswer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/ClassifyAndAnswer',
            main__pb2.Query.SerializeToString,
            main__pb2.GeneralAnswer.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def OpenChat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/OpenChat',
            main__pb2.Query.SerializeToString,
            main__pb2.ChatID.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CloseChat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/toclassifier.ToClassifier/CloseChat',
            main__pb2.ChatID.SerializeToString,
            main__pb2.GeneralAnswer.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)