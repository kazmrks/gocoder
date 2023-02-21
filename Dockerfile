FROM ubuntu:22.04

# Pythonと関連ツールをインストール
RUN apt-get update && \
    apt-get install -y software-properties-common git curl python3 python3-pip nodejs npm

# GOインストール用にPPA追加
RUN add-apt-repository ppa:longsleep/golang-backports

# GOインストール
RUN apt-get update && \
    apt-get install -y golang

# コンテスト補助アプリケーションをインストール
RUN pip install online-judge-tools
RUN npm install -g atcoder-cli

# GOの補助ツールインストール
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest

ENV PATH $PATH:/root/go/bin