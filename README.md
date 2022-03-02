# 参考

## クリーンアーキテクチャ
Clean ArchitectureでAPI Serverを構築してみる
https://qiita.com/hirotakan/items/698c1f5773a3cca6193e

Golang - EchoとGORMでClean Architecture APIを構築する
https://qiita.com/so-heee/items/0cca93008eae635c642a

## 認証

Go言語でEchoを用いて認証付きWebアプリの作成
https://qiita.com/x-color/items/24ff2491751f55e866cf

## MySQL

Sequel AceでMySQLに接続する
https://qiita.com/ucan-lab/items/b1304eee2157dbef7774

# 環境構築

コピーして.envファイルに記述
```
cp .env.example .env
```

```
docker-compose build
```

```
docker-compose up -d
```

```
docker exec -it shift bash
```

```
go run migrate/migrate.go
```

## エンドポイント
認証付きは -H 'Authorization: Bearer XXXXXXX' をつけてリクエスト

```
curl -X POST http://localhost:9000/signup -H 'Content-Type: application/json' -d '{"name": "tomoki", "email": "example@gmail.com", "password": "be3uiy31"}'
```

```
curl http://localhost:9000/users/1
```

```
curl http://localhost:9080/users
```

