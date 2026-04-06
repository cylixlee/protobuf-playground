"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings
from . import book_pb2 as book__pb2
from . import hello_pb2 as hello__pb2
GRPC_GENERATED_VERSION = '1.71.2'
GRPC_VERSION = grpc.__version__
_version_not_supported = False
try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True
if _version_not_supported:
    raise RuntimeError(f'The grpc package installed is at version {GRPC_VERSION},' + f' but the generated code in book_pb2_grpc.py depends on' + f' grpcio>={GRPC_GENERATED_VERSION}.' + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}' + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.')

class BookStub(object):
    """Book service
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Hello = channel.unary_unary('/proto.Book/Hello', request_serializer=hello__pb2.HelloRequest.SerializeToString, response_deserializer=hello__pb2.HelloResponse.FromString, _registered_method=True)
        self.Buy = channel.unary_unary('/proto.Book/Buy', request_serializer=book__pb2.BookRequest.SerializeToString, response_deserializer=book__pb2.BookResponse.FromString, _registered_method=True)

class BookServicer(object):
    """Book service
    """

    def Hello(self, request, context):
        """Say hello
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Buy(self, request, context):
        """Buy a book
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_BookServicer_to_server(servicer, server):
    rpc_method_handlers = {'Hello': grpc.unary_unary_rpc_method_handler(servicer.Hello, request_deserializer=hello__pb2.HelloRequest.FromString, response_serializer=hello__pb2.HelloResponse.SerializeToString), 'Buy': grpc.unary_unary_rpc_method_handler(servicer.Buy, request_deserializer=book__pb2.BookRequest.FromString, response_serializer=book__pb2.BookResponse.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('proto.Book', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('proto.Book', rpc_method_handlers)

class Book(object):
    """Book service
    """

    @staticmethod
    def Hello(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Book/Hello', hello__pb2.HelloRequest.SerializeToString, hello__pb2.HelloResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata, _registered_method=True)

    @staticmethod
    def Buy(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.Book/Buy', book__pb2.BookRequest.SerializeToString, book__pb2.BookResponse.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata, _registered_method=True)