FROM ubuntu:22.04

# SSL証明書検証エラー対応
RUN apt update && \
    apt install -y ca-certificates software-properties-common
COPY cert/* /usr/local/share/ca-certificates/
RUN update-ca-certificates

# GOインストール用にPPA追加
RUN add-apt-repository ppa:longsleep/golang-backports

# Python，関連ツールをインストール
RUN apt update && \
    apt install -y git curl python3 python3-pip nodejs npm

# コンテスト補助アプリケーションをインストール
RUN pip install online-judge-tools
RUN npm install -g atcoder-cli

# GOをインストール
RUN apt install -y  golang-go

# GOの補助ツールインストール
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest

ENV PATH $PATH:/root/go/bin
ENV REQUESTS_CA_BUNDLE /etc/ssl/certs/