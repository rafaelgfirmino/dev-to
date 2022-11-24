
# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sse_example

# final stage
FROM scratch
COPY --from=builder /app/sse_example /app/
EXPOSE 8080
ENTRYPOINT ["/app/sse_example"]