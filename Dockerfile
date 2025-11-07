FROM golang:1-alpine AS build

RUN apk update && apk add make git gcc musl-dev

ARG SERVICE

RUN make ${SERVICE}

RUN mv ${SERVICE} /${SERVICE}

FROM alpine:latest

ARG SERVICE

ENV APP=${SERVICE}

RUN apk add --no-cache ca-certificates && mkdir /app
COPY --from=build /${SERVICE} /app/${SERVICE}
ENTRYPOINT /app/${APP}
