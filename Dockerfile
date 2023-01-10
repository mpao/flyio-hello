FROM golang:1.19 AS build

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o /app/hello

FROM scratch
COPY --from=build /app/hello /hello
EXPOSE 8080
CMD ["/hello"]