FROM golang

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download -t
