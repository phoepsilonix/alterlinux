# AlterISOをビルドするためのDockerfile

# Buildステージ
FROM golang:1.22.3 AS build

WORKDIR /alteriso5

COPY ./alteriso5 .

RUN go mod download
RUN go build -o alteriso5

# 実行ステージ
FROM archlinux:base-20240101.0.204074 AS run

COPY . .
COPY --from=build /alteriso5/alteriso5 /usr/bin/alteriso5

WORKDIR /

RUN chmod +x ./usr/bin/alteriso5

RUN pacman -Sy --noconfirm
RUN pacman -S --noconfirm archiso arch-install-scripts

ENTRYPOINT ["/usr/bin/alteriso5"]
CMD ["build", "./profile"]