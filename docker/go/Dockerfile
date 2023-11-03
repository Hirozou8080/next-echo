# ベースイメージを指定
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# コンテナのポートを公開
EXPOSE 9000

ARG UID=1000
RUN useradd -m -u ${UID} hiroeru
USER ${UID}
