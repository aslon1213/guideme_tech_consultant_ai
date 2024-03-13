# you should be inside the file where main.proto is located
protoc *.proto --go_out=../ --go-grpc_out=../ --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc
python3.11 -m grpc_tools.protoc --proto_path=. \
  --python_out=../../sentece_classifier_bot/ --grpc_python_out=../../sentece_classifier_bot/ \
  ./classifier.proto
python3.11 -m grpc_tools.protoc --proto_path=. \
--python_out=../../tts/grpc/ --grpc_python_out=../../tts/grpc/ \
./tts.proto

python3.11 -m grpc_tools.protoc --proto_path=. \
--python_out=../../stt/ --grpc_python_out=../../stt \
./stt.proto

# swift 
protoc *.proto \
     --proto_path=. \
     --grpc-swift_opt=Client=true,Server=false \
     --grpc-swift_out=./swift \
     --swift_opt=Visibility=Public \
     --swift_out=./swift


  #/Users/aslonkhamidov/Desktop/code/tasks/customer_support_bot/sentence-classification/sentece_classifier_bot