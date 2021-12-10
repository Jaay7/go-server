FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go mod download

ENV MONGODB_URI=mongodb://localhost:27017
ENV APP_SECRET=1234567890

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test .

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080