FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o backend main.go

EXPOSE 8888

ENTRYPOINT ["/app/backend"]