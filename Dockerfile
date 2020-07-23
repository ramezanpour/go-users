FROM golang:1.14.6

WORKDIR /go/src/app

COPY . .

RUN go get -d -v
RUN go install -v

EXPOSE 3000

CMD ["users"]