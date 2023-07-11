FROM golang 

Label maintainer="Z01-Student"

COPY . /go/src/app

WORKDIR /go/src/app

RUN apt install wget

EXPOSE 80

EXPOSE 443

CMD go run .

