FROM golang:1.11 AS build
WORKDIR /go/src/asw-go-demo
ADD ./src .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch AS prod
COPY --from=build /go/src/asw-go-demo/asw-go-demo .
CMD ["./asw-go-demo"]
