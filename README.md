# Go + VSCode + Dev Container でAtCoder
VSCodeのDev ContainerとGoでAtCoderの環境セットを構築します。

## 備忘録

1. 始めるとき

    ```bash
    # コンテストの全問題をダウンロード
    acc new コンテストID
    # 解きたい問題のフォルダに移動
    cd コンテストID/問題
    ```

1. 実装してローカルテスト

    ```bash
    oj t -c "go run main.go" -d tests/
    ```

1. 提出するとき

    ```bash
    acc s main.go
    ```
