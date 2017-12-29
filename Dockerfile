FROM ubuntu:latest

RUN apt-get update && apt-get install -y hugo git


COPY bin/boop boop
COPY templates templates
COPY deploy.sh deploy.sh

CMD ./boop
