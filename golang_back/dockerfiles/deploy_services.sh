#!/bin/bash
#backend - golang
docker system prune -a
cd golang_back
pwd
docker container stop consultant_ai_back && docker container rm consultant_ai_back
docker build -f dockerfiles/dockerfile.golang -t aslon1213/consultant_ai_back .
docker run -d -p 9000:9000 --net cs_network --name consultant_ai_back aslon1213/consultant_ai_back
# classifier - python 
cd ../sentece_classifier_bot
pwd
docker container stop classifier && docker container rm classifier
docker build -f dockerfile.python -t aslon1213/classifier .
docker run -d -p 50050:50051 --net cs_network --name classifier aslon1213/classifier
cd ../