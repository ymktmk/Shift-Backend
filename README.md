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

## 環境構築

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

## 開発中
* 余力があったらホットリロードしたい
https://zenn.dev/tomi/articles/2020-10-14-go-docker

```
docker-compose exec app go run main.go
```

## エンドポイント
認証付きは -H 'Authorization: Bearer xxxxxxxxxx' をつけてリクエスト

* Firebase Authmのuid、名前、メールアドレス、会社名

```
curl -X POST http://localhost:9000/users/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxxx","name": "tomoki", "email": "tt@gmail.com", "company": {"name": "yy_company"}}'
```

```
curl http://localhost:9000/users/1
```

リクエストを送信した人の会社に属している人たちを取得
```
curl http://localhost:9000/company/users
```