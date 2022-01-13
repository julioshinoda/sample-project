FROM golang:1.14.3-alpine AS build
WORKDIR /src
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/api cmd/server/server.go

FROM scratch
ARG PORT
ARG SERVER_CLIENT
ARG SERVER_TOKEN
ARG SERVER_CLIENT_UUID
ENV PORT ${PORT}
ENV SERVER_CLIENT ${SERVER_CLIENT}
ENV SERVER_TOKEN ${SERVER_TOKEN}
ENV SERVER_CLIENT_UUID ${SERVER_CLIENT_UUID}
COPY --from=build /bin/api /bin/api
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/bin/api"]