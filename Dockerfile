FROM ubuntu:latest

RUN apt-get update && apt-get install -y wget git python-pip

RUN pip install awscli
RUN wget https://github.com/gohugoio/hugo/releases/download/v0.28/hugo_0.28_Linux-64bit.deb && dpkg -i hugo_0.28_Linux-64bit.deb

COPY bin/boop boop
COPY templates templates
COPY deploy.sh deploy.sh

CMD ./boop
