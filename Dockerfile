FROM golang:1.16 as base

COPY . /opt/app/api/tmp

WORKDIR /opt/app/api

CMD ["go", "test", "./..."]