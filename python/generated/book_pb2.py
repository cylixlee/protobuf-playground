"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(_runtime_version.Domain.PUBLIC, 5, 29, 0, '', 'book.proto')
_sym_db = _symbol_database.Default()
from . import hello_pb2 as hello__pb2
from .author import author_pb2 as author_dot_author__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nbook.proto\x12\x05proto\x1a\x0bhello.proto\x1a\x13author/author.proto"K\n\x0bBookRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12\r\n\x05price\x18\x02 \x01(\x05\x12\x1e\n\x06author\x18\x03 \x01(\x0b2\x0e.author.Author"\x1d\n\x0cBookResponse\x12\r\n\x05reply\x18\x01 \x01(\t2n\n\x04Book\x124\n\x05Hello\x12\x13.proto.HelloRequest\x1a\x14.proto.HelloResponse"\x00\x120\n\x03Buy\x12\x12.proto.BookRequest\x1a\x13.proto.BookResponse"\x00B8Z6github.com/cylixlee/protobuf-playground/internal/protob\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'book_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
    _globals['DESCRIPTOR']._loaded_options = None
    _globals['DESCRIPTOR']._serialized_options = b'Z6github.com/cylixlee/protobuf-playground/internal/proto'
    _globals['_BOOKREQUEST']._serialized_start = 55
    _globals['_BOOKREQUEST']._serialized_end = 130
    _globals['_BOOKRESPONSE']._serialized_start = 132
    _globals['_BOOKRESPONSE']._serialized_end = 161
    _globals['_BOOK']._serialized_start = 163
    _globals['_BOOK']._serialized_end = 273