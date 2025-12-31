FROM golang:1.25.1-alpine as builder

# GOの補助ツールインストール
RUN go install github.com/go-delve/delve/cmd/dlv@v1.25 && \
    go install honnef.co/go/tools/cmd/staticcheck@2025.1.1 && \
    go install golang.org/x/tools/gopls@latest

FROM golang:1.25.1

# SSL証明書検証エラー対応
COPY cert/* /usr/local/share/ca-certificates/
RUN update-ca-certificates
ENV REQUESTS_CA_BUNDLE=/etc/ssl/certs/

# 関連ツールをインストール
RUN apt update && \
    apt install -y nodejs npm python3 python3-pip python3-venv && \
    rm -rf /var/lib/apt/lists/*
RUN python3 -m venv /root/venv
ENV PATH=/root/venv/bin:$PATH

# online-judge-tools,atcoder-cliをインストール
RUN pip3 install setuptools && \
    pip3 install online-judge-tools && \
    pip3 cache purge && \
    npm install -g atcoder-cli && \
    npm cache clean --force

COPY --from=builder /go/bin/dlv /go/bin/dlv
COPY --from=builder /go/bin/staticcheck /go/bin/staticcheck
COPY --from=builder /go/bin/gopls /go/bin/gopls
