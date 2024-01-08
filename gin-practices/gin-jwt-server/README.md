# gin-jwt-server

`gin`のミドルウェア[gin-jwt](https://github.com/appleboy/gin-jwt)を使って、認証機能などを実装してみるテスト。
詳しくは[公式サイトのデモ](https://github.com/appleboy/gin-jwt#demo)などを参照。

## メモ

認証を行って`token`の発行を行う。（[httpie](https://httpie.io/docs/cli/pypi)を使うと便利）

```sh
http -v --json POST localhost:8000/login username=admin password=admin
```

発行された`token`を使って、認証した人だけみれる内容を閲覧する。（`token`は有効期限が1時間ぐらいなので注意。）

```sh
http localhost:8000/auth/hello 'Cookie:jwt={発行されたtoken}'
```
