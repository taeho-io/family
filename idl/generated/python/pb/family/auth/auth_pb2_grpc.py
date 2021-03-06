# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from pb.family.auth import auth_pb2 as pb_dot_family_dot_auth_dot_auth__pb2


class AuthServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Auth = channel.unary_unary(
        '/pb.family.auth.AuthService/Auth',
        request_serializer=pb_dot_family_dot_auth_dot_auth__pb2.AuthRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.AuthResponse.FromString,
        )
    self.Validate = channel.unary_unary(
        '/pb.family.auth.AuthService/Validate',
        request_serializer=pb_dot_family_dot_auth_dot_auth__pb2.ValidateRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.ValidateResponse.FromString,
        )
    self.Refresh = channel.unary_unary(
        '/pb.family.auth.AuthService/Refresh',
        request_serializer=pb_dot_family_dot_auth_dot_auth__pb2.RefreshRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.RefreshResponse.FromString,
        )
    self.Parse = channel.unary_unary(
        '/pb.family.auth.AuthService/Parse',
        request_serializer=pb_dot_family_dot_auth_dot_auth__pb2.ParseRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.ParseResponse.FromString,
        )


class AuthServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Auth(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Validate(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Refresh(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Parse(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AuthServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Auth': grpc.unary_unary_rpc_method_handler(
          servicer.Auth,
          request_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.AuthRequest.FromString,
          response_serializer=pb_dot_family_dot_auth_dot_auth__pb2.AuthResponse.SerializeToString,
      ),
      'Validate': grpc.unary_unary_rpc_method_handler(
          servicer.Validate,
          request_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.ValidateRequest.FromString,
          response_serializer=pb_dot_family_dot_auth_dot_auth__pb2.ValidateResponse.SerializeToString,
      ),
      'Refresh': grpc.unary_unary_rpc_method_handler(
          servicer.Refresh,
          request_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.RefreshRequest.FromString,
          response_serializer=pb_dot_family_dot_auth_dot_auth__pb2.RefreshResponse.SerializeToString,
      ),
      'Parse': grpc.unary_unary_rpc_method_handler(
          servicer.Parse,
          request_deserializer=pb_dot_family_dot_auth_dot_auth__pb2.ParseRequest.FromString,
          response_serializer=pb_dot_family_dot_auth_dot_auth__pb2.ParseResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'pb.family.auth.AuthService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
