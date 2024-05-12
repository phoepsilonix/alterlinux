# AlterISOをビルドするためのDockerfile

# Buildステージ
FROM golang:1.22.3 AS build

WORKDIR /alteriso5

COPY ./alteriso5 .

RUN go mod download
RUN go build -o alteriso5

# 実行ステージ
# FROM archlinux:base-20240101.0.204074 AS run

# COPY --from=build /alteriso5/alteriso5 /usr/bin/alteriso5

# ENTRYPOINT ["/usr/bin/alteriso5"]

# CMD ["build", "./profile"]