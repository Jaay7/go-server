FROM golang:1.14-alpine AS build
RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux go build -o ./out/api .

ENV MONGODB_URI=<your_mongodb_atlas_uri>
ENV APP_SECRET=<your_app_secret>


FROM alpine:latest

RUN apk add ca-certificates

COPY --from=build /tmp/app/out/api /app/api

WORKDIR "/app"

EXPOSE 5000

CMD ["./api"]