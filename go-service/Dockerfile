FROM golang:1.8

WORKDIR /go/src/server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Hot Reload
RUN go get github.com/githubnemo/CompileDaemon ./...
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ." -command="./server"

# Normal run.
#CMD ["server"]