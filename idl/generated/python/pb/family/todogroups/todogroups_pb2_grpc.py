# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from pb.family.todogroups import todogroups_pb2 as pb_dot_family_dot_todogroups_dot_todogroups__pb2


class TodoGroupsServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateTodoGroup = channel.unary_unary(
        '/pb.family.todogroups.TodoGroupsService/CreateTodoGroup',
        request_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.CreateTodoGroupRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.CreateTodoGroupResponse.FromString,
        )
    self.ListTodoGroups = channel.unary_unary(
        '/pb.family.todogroups.TodoGroupsService/ListTodoGroups',
        request_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.ListTodoGroupsRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.ListTodoGroupsResponse.FromString,
        )
    self.GetTodoGroup = channel.unary_unary(
        '/pb.family.todogroups.TodoGroupsService/GetTodoGroup',
        request_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.GetTodoGroupRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.GetTodoGroupResponse.FromString,
        )
    self.UpdateTodoGroup = channel.unary_unary(
        '/pb.family.todogroups.TodoGroupsService/UpdateTodoGroup',
        request_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.UpdateTodoGroupRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.UpdateTodoGroupResponse.FromString,
        )
    self.DeleteTodoGroup = channel.unary_unary(
        '/pb.family.todogroups.TodoGroupsService/DeleteTodoGroup',
        request_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.DeleteTodoGroupRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.DeleteTodoGroupResponse.FromString,
        )


class TodoGroupsServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def CreateTodoGroup(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListTodoGroups(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetTodoGroup(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateTodoGroup(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteTodoGroup(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TodoGroupsServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateTodoGroup': grpc.unary_unary_rpc_method_handler(
          servicer.CreateTodoGroup,
          request_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.CreateTodoGroupRequest.FromString,
          response_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.CreateTodoGroupResponse.SerializeToString,
      ),
      'ListTodoGroups': grpc.unary_unary_rpc_method_handler(
          servicer.ListTodoGroups,
          request_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.ListTodoGroupsRequest.FromString,
          response_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.ListTodoGroupsResponse.SerializeToString,
      ),
      'GetTodoGroup': grpc.unary_unary_rpc_method_handler(
          servicer.GetTodoGroup,
          request_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.GetTodoGroupRequest.FromString,
          response_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.GetTodoGroupResponse.SerializeToString,
      ),
      'UpdateTodoGroup': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateTodoGroup,
          request_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.UpdateTodoGroupRequest.FromString,
          response_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.UpdateTodoGroupResponse.SerializeToString,
      ),
      'DeleteTodoGroup': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteTodoGroup,
          request_deserializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.DeleteTodoGroupRequest.FromString,
          response_serializer=pb_dot_family_dot_todogroups_dot_todogroups__pb2.DeleteTodoGroupResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'pb.family.todogroups.TodoGroupsService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
