import grpc
import sys
import os
from time import sleep
from helloworld.helloworld_pb2 import FeedResponse
from helloworld.helloworld_pb2 import Empty
from helloworld.helloworld_pb2_grpc import FeederStub

os.environ['GRPC_VERBOSITY'] = 'debug'
os.environ['GRPC_TRACE'] = 'all'
def main():
    host = 'localhost'
    port = '8443'
    with open('../cert/server.crt', 'rb') as f:
        trusted_certs = f.read()

    credentials = grpc.ssl_channel_credentials(root_certificates=trusted_certs)
    print('1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx')
    with grpc.secure_channel('{}:{}'.format(host, port), credentials) as channel:
        print('2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx')
        stub = FeederStub(channel)
        print('3xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx')
        while True:
            responses = stub.GetNewFeed(Empty())
            print('4xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx')
            #sleep(3)
            for response in responses: # ここでエラーになる
                print('5xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx')
                print(response.message)

if __name__ == "__main__":
    main()