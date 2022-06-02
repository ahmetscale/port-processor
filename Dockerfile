FROM golang:1.18.0-alpine3.15 as builder

WORKDIR /src/go

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -tags postgres  -installsuffix cgo -o processor cmd/main.go

FROM scratch

WORKDIR /bin/processor

COPY --from=builder /src/go/processor .

CMD [ "./processor" ]


