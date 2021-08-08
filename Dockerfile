
FROM golang:1.16 as builder

RUN cd /tmp && \
    GO111MODULE=on GOPRIVATE=github.com/mattn/go-sqlite3 \
    go get -u github.com/rubenv/sql-migrate/...

WORKDIR /src
COPY . .
RUN go get 

FROM alpine:3
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.32-r0/glibc-2.32-r0.apk && \
    apk add glibc-2.32-r0.apk
COPY --from=builder /go/bin/sql-migrate /usr/bin
COPY --from=builder /src/database/migrations /database/migrations