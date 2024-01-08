# pg-api-echo

PostgreSQLを扱うWebAPIをechoで実装した君。

## ライブラリ

### echo

- [Echo - High performance, minimalist Go web framework](https://echo.labstack.com/)

### Validator

- [validator.v9 - gopkg.in/go-playground/validator.v9](https://gopkg.in/go-playground/validator.v9)

```sh
go get gopkg.in/go-playground/validator.v9
```

## テスト

httpieでJSONをPOSTするテスト。

```sh
http POST http://localhost:8080/api/comments text=test
```

## メモ

PostgreSQLを手元で起動して接続確認する場合、以下のようにDockerで起動すると楽。

```sh
docker run -e POSTGRES_USER={YOUR_USER} -e POSTGRES_PASSWORD={YOUR_PASSWORD} -p 5432:5432 postgres
```

いつものように`{YOUR_USER}`と`{YOUR_PASSWORD}`の箇所は、適宜環境に合わせて書き換える。

WebAPIサービスからの接続には`DSN`という環境変数を使っているので、以下のように設定しておいてからサービスを起動する。

```sh
export DSN=postgres://{YOUR_USER}:{YOUR_PASSWORD}@localhost:5432/postgres?sslmode=disable
```
