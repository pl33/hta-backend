# syntax=docker/dockerfile:1

FROM golang:1.19 AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/hta-server/*.go ./cmd/hta-server/
COPY models/*.go ./models/
COPY restapi/*.go ./restapi/
COPY restapi/operations/*.go ./restapi/operations/
COPY restapi/operations/category/*.go ./restapi/operations/category/
COPY restapi/operations/entry/*.go ./restapi/operations/entry/
COPY restapi/operations/login/*.go ./restapi/operations/login/
COPY schemas/*.go ./schemas/
RUN mkdir /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/hta-backend ./cmd/hta-server/main.go

FROM golang:1.19

WORKDIR /app

COPY --from=builder /build/hta-backend /app/hta-backend

VOLUME /sqlite
EXPOSE 8080
CMD [ "/app/hta-backend", "--scheme=http", "--host=0.0.0.0", "--port=8080" ]