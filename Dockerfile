FROM golang:1.18-alpine as builder

#DECLARE ENVIRONMENT VARIABLES HERE
ARG DB_HOST
ENV DB_HOST=${DB_HOST}
ARG DB_PORT
ENV DB_PORT=${DB_PORT}
ARG DB_USER
ENV DB_USER=${DB_USER}
ARG DB_PASSWORD
ENV DB_PASSWORD=${DB_PASSWORD}
ARG DB_NAME
ENV DB_NAME=${DB_NAME}
ARG JWT_SECRET_KEY
ENV JWT_SECRET_KEY=${JWT_SECRET_KEY}
ARG JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT
ENV JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=${JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT}
ARG GO_ENV
ENV GO_ENV=${PRODUCTION}
RUN echo $GO_ENV

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -v -o bin/backend
RUN echo $(ls)
EXPOSE 3000


# PENDING NEED TO TRY AGAIN
# FROM ubuntu:22.04
# RUN set -x && apt-get update && apt-get install -y \
#     ca-certificates && \
#     rm -rf /var/lib/apt/lists/*

FROM alpine
# Install any required dependencies.
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN echo $(ls)
COPY --from=builder /app/bin/backend /usr/local/bin/

CMD ["backend"]