FROM golang:latest

LABEL version="1.0.0"

LABEL "maintainer"="Mridhul Pax <mridhuljospax@gmail.com>"
LABEL "repository"="https://github.com/mridhul/cfn-deploy"

LABEL com.github.actions.name="Cloudformation Github Deploy"
LABEL com.github.actions.description="Cloudformation Github Deploy"
LABEL com.github.actions.icon="upload-cloud"
LABEL com.github.actions.color="orange"

RUN apt-get update && apt-get install -y awscli git build-essential python3-pip 

RUN ls -l 
RUN rm -rf /var/lib/apt/lists/*

RUN pip3 install cloudformation-cli cloudformation-cli-java-plugin cloudformation-cli-go-plugin cloudformation-cli-python-plugin
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
