import hello_pb2 as _hello_pb2
from author import author_pb2 as _author_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class BookRequest(_message.Message):
    __slots__ = ('title', 'price', 'author')
    TITLE_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    AUTHOR_FIELD_NUMBER: _ClassVar[int]
    title: str
    price: int
    author: _author_pb2.Author

    def __init__(self, title: _Optional[str]=..., price: _Optional[int]=..., author: _Optional[_Union[_author_pb2.Author, _Mapping]]=...) -> None:
        ...

class BookResponse(_message.Message):
    __slots__ = ('reply',)
    REPLY_FIELD_NUMBER: _ClassVar[int]
    reply: str

    def __init__(self, reply: _Optional[str]=...) -> None:
        ...