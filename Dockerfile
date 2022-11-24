FROM golang:1.16 as base

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY . /opt/app/api/tmp

WORKDIR /opt/app/api
RUN chmod -R 777 /opt/app/api/tmp
CMD ["air"]