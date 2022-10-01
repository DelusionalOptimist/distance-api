FROM golang:1.19 AS build

WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /distance-api main.go

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch

ARG API_KEY
ENV API_KEY=${API_KEY}
ARG PORT
ENV PORT=${PORT}

COPY --from=build /distance-api /distance-api
COPY --from=build /app/static /static

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD [ "/distance-api" ]
