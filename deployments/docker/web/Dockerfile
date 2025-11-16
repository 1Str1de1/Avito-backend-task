FROM golang:alpine AS builder

WORKDIR /Avito-backend-task

COPY  go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/Avito-backend-task ./cmd/server

FROM alpine

COPY --from=builder /bin/Avito-backend-task /bin/Avito-backend-task

EXPOSE 80

ENTRYPOINT ["/bin/Avito-backend-task"]
