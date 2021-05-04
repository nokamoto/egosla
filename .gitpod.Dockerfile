FROM gitpod/workspace-mysql

RUN sudo apt-get update && sudo apt-get install -y clang-format protobuf-compiler
