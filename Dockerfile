FROM golang
ADD . /opt/zy-o2o-server
WORKDIR /opt/zy-o2o-server
RUN go install 

ENTRYPOINT /opt/zy-o2o-server
EXPOSE 8001