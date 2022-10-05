FROM golang:1.18

WORKDIR $GOPATH/src/github.com/Natasha-H-S/Go-Book-API

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 10000

CMD go run . 
