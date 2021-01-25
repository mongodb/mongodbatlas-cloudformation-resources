FROM golang:1.14-alpine

RUN apk --no-cache add py3-pip make git zip

RUN pip3 install cloudformation-cli-go-plugin

COPY . /build

WORKDIR /build

RUN cd project && \
    make && \
    cd bin && \
    zip -X ../handler.zip ./handler && \
    cd .. && \
    cp mongodb-atlas-project.json schema.json && \
    zip -X ../mongodb-atlas-project.zip ./handler.zip ./schema.json ./.rpdk-config 
    
RUN cd cluster && \
    make && \
    cd bin && \
    zip -X ../handler.zip ./handler && \
    cd .. && \
    cp mongodb-atlas-cluster.json schema.json && \
    zip -X ../mongodb-atlas-cluster.zip ./handler.zip ./schema.json ./.rpdk-config 
 
RUN cd database-user && \
    make && \
    cd bin && \
    zip -X ../handler.zip ./handler && \
    cd .. && \
    cp mongodb-atlas-databaseuser.json schema.json && \
    zip -X ../mongodb-atlas-databaseuser.zip ./handler.zip ./schema.json ./.rpdk-config 

RUN cd project-ip-access-list && \
    make && \
    cd bin && \
    zip -X ../handler.zip ./handler && \
    cd .. && \
    cp mongodb-atlas-projectipaccesslist.json schema.json && \
    zip -X ../mongodb-atlas-projectipaccesslist.zip ./handler.zip ./schema.json ./.rpdk-config 
    
RUN cd network-peering && \
    make && \
    cd bin && \
    zip -X ../handler.zip ./handler && \
    cd .. && \
    cp mongodb-atlas-networkpeering.json schema.json && \
    zip -X ../mongodb-atlas-networkpeering.zip ./handler.zip ./schema.json ./.rpdk-config 
    
CMD mkdir -p /output/ && mv /build/*.zip /output/