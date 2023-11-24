FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go get github.com/josepht96/learning/projects/scout
RUN go build -o /scout

RUN chmod +x /scout

CMD [ "/scout" ]

FROM golang:1.21-alpine
RUN apk add curl
COPY --from=build /scout /bin/scout
EXPOSE 8080
CMD ["/bin/scout"]