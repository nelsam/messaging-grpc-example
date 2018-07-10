# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: instant_messaging_service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import user_pb2 as user__pb2
import message_pb2 as message__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='instant_messaging_service.proto',
  package='messaging',
  syntax='proto3',
  serialized_pb=_b('\n\x1finstant_messaging_service.proto\x12\tmessaging\x1a\nuser.proto\x1a\rmessage.proto\"n\n\x0b\x43onnectAuth\x12\x1d\n\x04\x61uth\x18\x01 \x01(\x0b\x32\x0f.messaging.Auth\x12 \n\x04\x66rom\x18\x02 \x01(\x0b\x32\x12.messaging.Profile\x12\x1e\n\x02to\x18\x03 \x01(\x0b\x32\x12.messaging.Profile\"N\n\x08\x41uthedIM\x12\x1d\n\x04\x61uth\x18\x01 \x01(\x0b\x32\x0f.messaging.Auth\x12#\n\x07message\x18\x02 \x01(\x0b\x32\x12.messaging.Message2\x85\x01\n\x15InstantMessageService\x12\x31\n\x04Send\x12\x13.messaging.AuthedIM\x1a\x12.messaging.Message\"\x00\x12\x39\n\x07\x43onnect\x12\x16.messaging.ConnectAuth\x1a\x12.messaging.Message\"\x00\x30\x01\x62\x06proto3')
  ,
  dependencies=[user__pb2.DESCRIPTOR,message__pb2.DESCRIPTOR,])




_CONNECTAUTH = _descriptor.Descriptor(
  name='ConnectAuth',
  full_name='messaging.ConnectAuth',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='auth', full_name='messaging.ConnectAuth.auth', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='from', full_name='messaging.ConnectAuth.from', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='to', full_name='messaging.ConnectAuth.to', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=73,
  serialized_end=183,
)


_AUTHEDIM = _descriptor.Descriptor(
  name='AuthedIM',
  full_name='messaging.AuthedIM',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='auth', full_name='messaging.AuthedIM.auth', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='messaging.AuthedIM.message', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=185,
  serialized_end=263,
)

_CONNECTAUTH.fields_by_name['auth'].message_type = user__pb2._AUTH
_CONNECTAUTH.fields_by_name['from'].message_type = user__pb2._PROFILE
_CONNECTAUTH.fields_by_name['to'].message_type = user__pb2._PROFILE
_AUTHEDIM.fields_by_name['auth'].message_type = user__pb2._AUTH
_AUTHEDIM.fields_by_name['message'].message_type = message__pb2._MESSAGE
DESCRIPTOR.message_types_by_name['ConnectAuth'] = _CONNECTAUTH
DESCRIPTOR.message_types_by_name['AuthedIM'] = _AUTHEDIM
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ConnectAuth = _reflection.GeneratedProtocolMessageType('ConnectAuth', (_message.Message,), dict(
  DESCRIPTOR = _CONNECTAUTH,
  __module__ = 'instant_messaging_service_pb2'
  # @@protoc_insertion_point(class_scope:messaging.ConnectAuth)
  ))
_sym_db.RegisterMessage(ConnectAuth)

AuthedIM = _reflection.GeneratedProtocolMessageType('AuthedIM', (_message.Message,), dict(
  DESCRIPTOR = _AUTHEDIM,
  __module__ = 'instant_messaging_service_pb2'
  # @@protoc_insertion_point(class_scope:messaging.AuthedIM)
  ))
_sym_db.RegisterMessage(AuthedIM)



_INSTANTMESSAGESERVICE = _descriptor.ServiceDescriptor(
  name='InstantMessageService',
  full_name='messaging.InstantMessageService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=266,
  serialized_end=399,
  methods=[
  _descriptor.MethodDescriptor(
    name='Send',
    full_name='messaging.InstantMessageService.Send',
    index=0,
    containing_service=None,
    input_type=_AUTHEDIM,
    output_type=message__pb2._MESSAGE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Connect',
    full_name='messaging.InstantMessageService.Connect',
    index=1,
    containing_service=None,
    input_type=_CONNECTAUTH,
    output_type=message__pb2._MESSAGE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_INSTANTMESSAGESERVICE)

DESCRIPTOR.services_by_name['InstantMessageService'] = _INSTANTMESSAGESERVICE

# @@protoc_insertion_point(module_scope)