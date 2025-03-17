FROM golang:latest

COPY ./ ./


RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh


RUN make build

#CMD ["./apishelfshare"]