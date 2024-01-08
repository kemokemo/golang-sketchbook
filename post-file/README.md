# post file

`Form`でファイルをPOSTしたらローカルに保存するWebAPIを作ってみるテスト。

## 確認

`httpie`を使った動作確認方法は以下。

```sh
http -f POST localhost:8080/upload file@./README.md
```
