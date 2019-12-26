FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o hello-web.go .
CMD ["/app/hello-web"]
