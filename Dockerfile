
# --------- building environment -------
FROM golang:1.23.7-alpine AS builder

WORKDIR /workdir/

COPY . .

RUN go mod download

RUN go build -o /tiny-go


# ---------- deploying environment ------
FROM alpine:3.21

COPY --from=builder tiny-go .
COPY --from=builder /workdir/templates ./templates

RUN ls

RUN ls /

EXPOSE 8080

CMD ["/tiny-go"]