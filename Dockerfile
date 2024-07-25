FROM golang:1.17 as build
WORKDIR /app
COPY . .
RUN go build -o /server .

FROM scratch
COPY --from=build /server /server
EXPOSE 3002
CMD ["/server"]
