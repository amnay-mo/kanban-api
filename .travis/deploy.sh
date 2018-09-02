#!/bin/bash

set -ex

if [ ! -e ${HOME}/google-cloud-sdk/bin/gcloud ]; then
   curl https://sdk.cloud.google.com | bash >> /dev/null;
   source ${HOME}/google-cloud-sdk/path.bash.inc
fi

openssl aes-256-cbc -K $encrypted_e3c80bef8e36_key -iv $encrypted_e3c80bef8e36_iv -in .travis/myfoobarproject-9a043adac874.json.enc -out myfoobarproject-9a043adac874.json -d
gcloud auth activate-service-account --key-file=myfoobarproject-9a043adac874.json

gcloud auth configure-docker

docker build -t eu.gcr.io/myfoobarproject/kanban-api:${TRAVIS_TAG} -t eu.gcr.io/myfoobarproject/kanban-api:latest .

docker push eu.gcr.io/myfoobarproject/kanban-api:${TRAVIS_TAG}

docker push eu.gcr.io/myfoobarproject/kanban-api:latest
