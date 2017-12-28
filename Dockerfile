FROM ubuntu:latest

RUN apt-get update && apt-get install -y golang-go hugo


COPY bin/boop boop
COPY templates templates
COPY deploy.sh deploy.sh

CMD ./boop
