FROM golang:1.15 as build

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o fejkApi

FROM debian:buster

WORKDIR /app
COPY --from=build /app/fejkApi /app
COPY --from=build /app/.env /app

EXPOSE 8124 8124

CMD ["./fejkApi"]
