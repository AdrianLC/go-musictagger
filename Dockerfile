FROM golang

WORKDIR /app

RUN go-wrapper download \
    github.com/stretchr/testify/assert \
    github.com/mikkyang/id3-go
