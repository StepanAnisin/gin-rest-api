FROM golang:latest
ENV GOPATH=/
COPY . .
RUN apt-get update
RUN apt-get -y install postgresql-client
RUN go mod download
RUN go build -o gin-rest-api ./cmd/main.go
CMD ["./gin-rest-api"]