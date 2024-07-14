FROM golang:1.22.5 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temppc ./cmd/

FROM scratch
WORKDIR /app
COPY --from=build /app/temppc .
COPY --from=build /app/cmd/.env .
ENTRYPOINT ["./temppc"] 