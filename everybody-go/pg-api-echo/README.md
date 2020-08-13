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