# gRPC and Remote Procedure Call (With Protobuf) - Chat Server using go

What is gRPC and it is advantages

gRPC is a Free, open-source framework developed by google. Has a high-performance, feature-rich RPC framework. Also considers language-independent and supports many languages.

How does gRPC work

First, the client generares stub that provides the same methods as the server.
Second, the client calls the gRPC framework under the hood to exchange information over the network.
Third, the client and the server use stubs to interact with each other.
Finally, Separation of concern. The server only cares about the core service logic.

What are Protobuffers 
It is Google's mature open-source mechanism for serializing structured data,Also the protocol buffer data is structured as messages
