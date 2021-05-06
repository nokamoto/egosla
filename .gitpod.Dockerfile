FROM gitpod/workspace-mysql

RUN sudo apt-get update && sudo apt-get install -y clang-format protobuf-compiler

RUN curl -sSL https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64 > protoc-gen-grpc-web
RUN chmod +x protoc-gen-grpc-web
RUN sudo mv protoc-gen-grpc-web /usr/local/bin
