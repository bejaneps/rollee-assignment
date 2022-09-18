FROM golang:1.19-bullseye as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux CGO_ENABLED=0 \
    go build -ldflags "-s -w" -o app ./cmd

# 11aeeec130d652b38b022507f722f9a826eff9a5 is latest-amd64 version
FROM gcr.io/distroless/static-debian11:11aeeec130d652b38b022507f722f9a826eff9a5 as base
COPY --from=builder /go/src/app /app

WORKDIR /app

EXPOSE 7171

ENTRYPOINT ["./app"]
