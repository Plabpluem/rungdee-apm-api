FROM golang:1.25-bookworm AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -tags netgo -ldflags "-s -w" -o /out/app

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
        chromium \
        fonts-thai-tlwg \
        ca-certificates \
        tzdata \
    && rm -rf /var/lib/apt/lists/*

ENV CHROME_PATH=/usr/bin/chromium \
    TZ=Asia/Bangkok

WORKDIR /app
COPY --from=builder /out/app ./app
COPY internal/adapters/invoice/template ./internal/adapters/invoice/template

EXPOSE 8080
CMD ["./app"]
