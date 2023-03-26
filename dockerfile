FROM golang:1.18

RUN mkdir -p /usr/src/app/
workdir /usr/src/app/

COPY . /usr/src/app/

# install psql for wait-for-postgres.sh
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# install make to run migrates
RUN apt-get -y install make

# install migrates
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate

# build go app
RUN go mod download
RUN go build -o crud-books cmd/app/main.go

CMD ["./crud-books"]

