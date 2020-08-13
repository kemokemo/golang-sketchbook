# pgsample

`PostgreSQL`データベースに接続する君。

## 必要ライブラリ

公式の`postgresql`接続用ライブラリを取得しておく。

```sh
go get github.com/lib/pq
```

## 使い方

### PostgreSQLを起動

`Postgres`データベースを`docker`で起動するコマンドは以下。

```sh
docker run --name pgsampledb -e POSTGRES_USER=posuser -e POSTGRES_PASSWORD=pospass -e POSTGRES_INITDB_ARGS="--encoding=UTF-8 --locale=C" -p 5432:5432 -d postgres:latest
```

上記では以下の設定を使った。変更する場合は、この後述の`DSN`設定にも反映させる必要がある。

- Username: posuser
- Password: pospass
- Port: 5432

### 接続用のDSNを設定

プログラムから接続するための`DSN(=Data source name)`を、以下のように環境変数で設定する(fish shell用)。

```sh
set -x DSN "postgres://posuser:pospass@localhost:5432/postgres?sslmode=disable"
```

#### 補足

`DSN`のフォーマットは以下。

```sh
postgres://{username}:{password}@{hostname}/{db_name}{option}
```

うまくいかない場合は、StackOverflowの以下のスレッドが参考になるかも。

- [SSL is not enabled on the server](https://stackoverflow.com/questions/21959148/ssl-is-not-enabled-on-the-server)

### クライアントから接続

テーブルを作るため、まずはmacOSで`psql`クライアントを使える状態にする。

```sh
brew install postgres
```

接続する。

```sh
psql $DSN
```

### `users`テーブルを作成

```sql
CREATE TABLE users (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
age INTEGER NOT NULL);
```