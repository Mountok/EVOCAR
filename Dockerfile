FROM golang:latest

COPY ./ ./
RUN go build -o main cmd/main.go
CMD [ "./main" ]
