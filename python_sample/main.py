import grpc
import sys
import os

from helloworld.helloworld_pb2 import HelloRequest
from helloworld.helloworld_pb2 import HelloReply
from helloworld.helloworld_pb2_grpc import GreeterStub

os.environ['GRPC_VERBOSITY'] = 'debug'
os.environ['GRPC_TRACE'] = 'all'
def main():
    host = 'localhost'
    port = '8443'

    with open('../cert/server.crt', 'rb') as f:
        trusted_certs = f.read()

    credentials = grpc.ssl_channel_credentials(root_certificates=trusted_certs)
    channel = grpc.secure_channel('{}:{}'.format(host, port), credentials)

    stub = GreeterStub(channel)
    req = HelloRequest(name='hi prod!')
    res = stub.SayHello(req)
    print(res)

if __name__ == "__main__":
    main()