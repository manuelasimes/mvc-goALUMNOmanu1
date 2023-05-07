FROM golang:1.17

RUN mkdir /gin
ADD . /gin
WORKDIR /gin

RUN go mod init gin

RUN go mod tidy

RUN go build -o gin .

RUN chmod +x /gin/gin

ENTRYPOINT ["/gin/gin"]