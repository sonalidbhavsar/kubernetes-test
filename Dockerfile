FROM golang:1.16
RUN mkdir hello_world
COPY . hello_world
WORKDIR hello_world
RUN go build cmd/hello_world.go 
ENTRYPOINT ["./hello_world"]