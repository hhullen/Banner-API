FROM golang:1.22.1-alpine3.19 AS builder_stage

WORKDIR /app
COPY . .
RUN go build -o ./server .

FROM alpine:3.19

ARG UID=1001
ARG GID=1001
ARG USER=Kakashi
ARG DIR=/home/${USER}

RUN \
    addgroup -g ${GID} ${USER} && \
    adduser -D -G ${USER} -u ${UID} ${USER}

WORKDIR ${DIR}

COPY --from=builder_stage /app/server .

RUN chown ${USER}:${USER} server
USER ${USER}

ENTRYPOINT [ "./server" ]
