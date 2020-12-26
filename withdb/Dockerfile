FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

ENV ELASTIC_APM_SERVER_URL=http://apm-server:8200

ENV ELASTIC_APM_SERVICE_NAME=withdb

ENV ELASTIC_APM_SECRET_TOKEN=xxVpmQB2HMzCL9PgBHVrnxjNXXw5J7bd79DFm6sjBJR5HPXDhcF8MSb3vv4bpg44

CMD [ "go","run","withdb.go" ]