FROM golang:alpine AS build

LABEL summary="PiHole Exporter Docker image" \
      description="Prometheus Exporter for PiHole" \
      name="github.com/projectbadger/pihole_exporter" \
      version="dev" \
      url="https://github.com/projectbadger/pihole_exporter"

RUN apk add --no-cache alpine-sdk bash
WORKDIR /go/src/github.com/nlamirault/pihole_exporter
COPY . .
RUN go build -o /app/pihole_exporter pihole_exporter.go && chmod +x /app/pihole_exporter

FROM alpine:latest
COPY --from=build /app/pihole_exporter /app/pihole_exporter
WORKDIR /app
ENTRYPOINT ["./pihole_exporter"]
CMD ["-h"]
EXPOSE 9311
