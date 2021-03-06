# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: message_service.proto for package 'messaging'

require 'grpc'
require 'message_service_pb'

module Messaging
  module MessageService
    # MessageService is our message microservice.
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'messaging.MessageService'

      # Send sends a new message to a user.
      rpc :Send, AuthedMessage, Thread
      # Reply replies to a thread.
      rpc :Reply, ThreadReply, Thread
      # Delete deletes a message.
      rpc :Delete, AuthedMessage, Thread
    end

    Stub = Service.rpc_stub_class
  end
end
