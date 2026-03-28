FROM golang:1.25-bullseye

# ติดตั้ง chromium
RUN apt-get update && apt-get install -y chromium

WORKDIR /app
COPY . .

RUN go build -tags netgo -ldflags "-s -w" -o app

CMD ["./app"]