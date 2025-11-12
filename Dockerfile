# ---- Build stage ----
FROM golang:1.25-alpine AS build

RUN apk add --no-cache make git gcc musl-dev

ENV SERVICE=graphql-auth
WORKDIR /src

# Copy Go modules first (better for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -o ${SERVICE} ./cmd/${SERVICE}

# ---- Runtime stage ----
FROM alpine:latest

ENV SERVICE=graphql-auth
WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy the built binary
COPY --from=build /src/${SERVICE} /app/${SERVICE}

ENTRYPOINT ["/app/graphql-auth"]
