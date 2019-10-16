FROM golang:1.12.7

WORKDIR /go/src/tarpit
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["tarpit"]