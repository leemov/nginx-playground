FROM golang:alpine as build
WORKDIR /app
COPY main.go .
RUN go build main.go

FROM alpine:latest
WORKDIR /app
COPY . .
COPY --from=build /app/main .
# RUN chmod 755 script.sh
CMD ./main