FROM golang:1.16

LABEL maintainer="Rayon Lindraz Nunes <rayonnunes@hotmail.com>"

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8000

CMD ["./main"]