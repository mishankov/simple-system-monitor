FROM node:22 AS buildjs
WORKDIR /app
COPY webapp .
RUN npm ci
RUN npm run build

FROM golang:1.23 AS buildgo
WORKDIR /app
COPY . .
COPY --from=buildjs cmd/server/build cmd/server/build
RUN CGO_ENABLED=0 go build -o ./build/simple-system-monitor ./cmd/server

FROM alpine:3
WORKDIR /app
COPY --from=buildgo /app/build/simple-system-monitor /app/simple-system-monitor
EXPOSE 4442
ENTRYPOINT ["/app/simple-system-monitor"]
