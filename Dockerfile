FROM golang:1.11-alpine3.10 AS build
WORKDIR /go/src/asw-go-demo
RUN apk --no-cache add git;go get -d -v github.com/go-sql-driver/mysql
ADD ./src .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o asw-go-demo .

FROM scratch AS prod
EXPOSE 8080
COPY --from=build /go/src/asw-go-demo/asw-go-demo .
CMD ["./asw-go-demo"]

