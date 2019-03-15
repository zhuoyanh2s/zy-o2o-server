FROM golang
ADD . $GOPATH/opt/zy-o2o-service
WORKDIR $GOPATH/opt/zy-o2o-service
RUN go build .


ENTRYPOINT ["./zy-o2o-service"]
EXPOSE 8080