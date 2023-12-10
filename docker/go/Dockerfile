# ベースイメージを指定
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# コンテナのポートを公開

# コンテナに永続的な環境変数を定義します。コピーする対象 コンテナ内のコピー先で指定できます
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8000

# ARG UID=1000
# RUN useradd -m -u ${UID} hiroeru
# USER ${UID}
