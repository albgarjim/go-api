FROM alpine:3.9.4
RUN apk --no-cache add ca-certificates

WORKDIR /go/src/go-api
RUN mkdir logs

# Document that the service listens on port e.g.: 81.
ARG PORT
EXPOSE $PORT

COPY go-api .

CMD ["./go-api"]
