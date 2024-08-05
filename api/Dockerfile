FROM golang:1.22.5-alpine AS builder
WORKDIR /mahlzeit
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o mahlzeit

FROM scratch
COPY --from=builder /mahlzeit/mahlzeit ./mahlzeit
CMD ["./mahlzeit"]
