# you should be inside the file where main.proto is located
protoc *.proto --go_out=../ --go-grpc_out=../ --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc
python3.11 -m grpc_tools.protoc --proto_path=. \
  --python_out=../../sentence-classification/sentece_classifier_bot/ --grpc_python_out=../../sentence-classification/sentece_classifier_bot/ \
  ./main.proto


  #/Users/aslonkhamidov/Desktop/code/tasks/customer_support_bot/sentence-classification/sentece_classifier_bot