
#note that you should have the protoc-gen-grpc-swift and protoc-gen-swift installed in your system - PATH
# learn about building these plugins https://github.com/grpc/grpc-swift
protoc *.proto \
     --proto_path=. \
     --grpc-swift_opt=Client=true,Server=false \
     --grpc-swift_out=./swift \
     --swift_opt=Visibility=Public \
     --swift_out=./swift
