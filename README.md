# Go + VSCode + Dev Container でAtCoder
VSCodeのDev ContainerとGoでAtCoderの環境セットを構築します。

コンテスト用のコードはsubmoduleで[atcoder-go](https://github.com/kazmrks/atcoder-go)へ分離しています。
## 備忘録

1. 始めるとき

    ```bash
    # コンテスト用のフォルダへ移動
    cd atcoder-go
    # コンテストの全問題をダウンロード(ABC289など)
    acc new ABC289
    # 解きたい問題のフォルダに移動
    cd ABC289/a
    # テンプレートファイルをコピー
    cp ../../../template.go main.go
    ```

1. 実装してローカルテスト(VSCodeのTaskあり: ojgo)

    ```bash
    oj t -c "go run main.go" -d tests/
    ```

1. 提出するとき

    ```bash
    acc s main.go
    ```

## Runtime Error(RE)対応

1. bufio.Scannerのデフォルトサイズエラー

    読み込める文字の最大サイズが`MaxScanTokenSize = 64 * 1024`で約6万5千文字のため、Scan時に超える場合エラー。

    `(*Scanner).Buffer`でサイズ指定する。

    ```go
    sc.Buffer([]byte{}, math.MaxInt64)
    ```