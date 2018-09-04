# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from pb.family.notes import notes_pb2 as pb_dot_family_dot_notes_dot_notes__pb2


class NotesServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.CreateNote = channel.unary_unary(
        '/pb.family.notes.NotesService/CreateNote',
        request_serializer=pb_dot_family_dot_notes_dot_notes__pb2.CreateNoteRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.CreateNoteResponse.FromString,
        )
    self.GetNote = channel.unary_unary(
        '/pb.family.notes.NotesService/GetNote',
        request_serializer=pb_dot_family_dot_notes_dot_notes__pb2.GetNoteRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.GetNoteResponse.FromString,
        )
    self.ListNotes = channel.unary_unary(
        '/pb.family.notes.NotesService/ListNotes',
        request_serializer=pb_dot_family_dot_notes_dot_notes__pb2.ListNotesRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.ListNotesResponse.FromString,
        )
    self.UpdateNote = channel.unary_unary(
        '/pb.family.notes.NotesService/UpdateNote',
        request_serializer=pb_dot_family_dot_notes_dot_notes__pb2.UpdateNoteRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.UpdateNoteResponse.FromString,
        )
    self.DeleteNote = channel.unary_unary(
        '/pb.family.notes.NotesService/DeleteNote',
        request_serializer=pb_dot_family_dot_notes_dot_notes__pb2.DeleteNoteRequest.SerializeToString,
        response_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.DeleteNoteResponse.FromString,
        )


class NotesServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def CreateNote(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetNote(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListNotes(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateNote(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteNote(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_NotesServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'CreateNote': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNote,
          request_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.CreateNoteRequest.FromString,
          response_serializer=pb_dot_family_dot_notes_dot_notes__pb2.CreateNoteResponse.SerializeToString,
      ),
      'GetNote': grpc.unary_unary_rpc_method_handler(
          servicer.GetNote,
          request_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.GetNoteRequest.FromString,
          response_serializer=pb_dot_family_dot_notes_dot_notes__pb2.GetNoteResponse.SerializeToString,
      ),
      'ListNotes': grpc.unary_unary_rpc_method_handler(
          servicer.ListNotes,
          request_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.ListNotesRequest.FromString,
          response_serializer=pb_dot_family_dot_notes_dot_notes__pb2.ListNotesResponse.SerializeToString,
      ),
      'UpdateNote': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateNote,
          request_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.UpdateNoteRequest.FromString,
          response_serializer=pb_dot_family_dot_notes_dot_notes__pb2.UpdateNoteResponse.SerializeToString,
      ),
      'DeleteNote': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteNote,
          request_deserializer=pb_dot_family_dot_notes_dot_notes__pb2.DeleteNoteRequest.FromString,
          response_serializer=pb_dot_family_dot_notes_dot_notes__pb2.DeleteNoteResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'pb.family.notes.NotesService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
