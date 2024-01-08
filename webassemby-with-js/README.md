# WebAssembly with js

[golang公式のWebAssembly解説](https://github.com/golang/go/wiki/WebAssembly)を読みつつ実践してみた系。
以下を試してみたい。

* WebAssemblyとしてビルドして、Webページからgolangの処理を実行する
* JavaScriptからgolangの処理へ情報を渡す
* golangの処理からJavaScriptへ情報返す

## Build方法

バイナリのビルドは以下の通り。

```sh
GOOS=js GOARCH=wasm go build -o main.wasm
```

ビルドしたのと同じバージョンの`go`から`wasm_exec.js`をコピーして使う必要がある。以下を実行。

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

## 確認方法

まず以下のコマンドで`goexec`をインストールする。

```sh
go install github.com/shurcooL/goexec
```

その後、以下を実行。（めっちゃ便利！`goexec`すごい！）

```sh
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

ブラウザで`http://localhost:8080/`を開いたら動作確認が可能。

## メモ

* [Go WebAssembly Tutorial - Building a Calculator](https://www.youtube.com/watch?v=4kBvvk2Bzis)
  * `golang`のメソッドをJavaScriptから呼び出すための登録方法がわかりやすい
  * 少し情報が古くて`syscall/js`の関数名が変わっているので、以下の記事も参照
* [Golang's syscall/js js.NewCallback is undefined](https://stackoverflow.com/questions/55800163/golangs-syscall-js-js-newcallback-is-undefined)
