FROM golang:1.12
ADD . /opt/zy-o2o-service
WORKDIR /opt/zy-o2o-service
VOLUME ["/config","/config"]
RUN go build .


ENTRYPOINT ["./zy-o2o-service"]
EXPOSE 8080