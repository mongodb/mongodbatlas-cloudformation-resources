##
## BASE dockerfile - can build and one resource type

FROM golang:latest as base


# Install python-pip, zip
RUN apt-get update && apt-get install -y zip python3-pip

FROM base as awstools

RUN pip3 install cloudformation-cli cloudformation-cli-go-plugin cfn-lint taskcat 

FROM awstools as builder

COPY . /build/
#COPY requirements.txt Makefile .rpdk-config mongodb-atlas-project.json /build/

#COPY cmd/ /build/cmd/
#COPY ../go.mod /build/go.mod
WORKDIR /build
RUN pwd && ls -alh
CMD ./cfn-submit-helper.sh project cluster database-user
#RUN echo "build complete: " && pwd && ls -lah

#RUN mkdir -p /output/ && cp **/mongodb-atlas-*.zip /output/

#RUN echo "Build output:" && ls -lah /output
