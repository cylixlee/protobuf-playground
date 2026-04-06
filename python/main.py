import grpc
import os
from generated import hello_pb2 as pb
from generated import hello_pb2_grpc as G


def main():
    port = os.getenv("PP_SERVER_PORT")
    with grpc.insecure_channel(f"localhost:{port}") as channel:
        stub = G.HelloStub(channel)
        resp = stub.Hello(pb.HelloRequest(name="CYLIX in Python"))
        print("Server>>> ", resp.reply)


if __name__ == "__main__":
    main()
