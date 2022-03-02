# 参考

## クリーンアーキテクチャ
Clean ArchitectureでAPI Serverを構築してみる
https://qiita.com/hirotakan/items/698c1f5773a3cca6193e

Golang - EchoとGORMでClean Architecture APIを構築する
https://qiita.com/so-heee/items/0cca93008eae635c642a

## 認証

Go言語でEchoを用いて認証付きWebアプリの作成
https://qiita.com/x-color/items/24ff2491751f55e866cf

## エンドポイント
認証付きは -H 'Authorization: Bearer XXXXXXX' をつけてリクエスト

### Create
```
curl -X POST http://localhost:9000/users/create -H 'Content-Type: application/json' -d '{"name":"yyy"}'
```

### Show
```
curl http://localhost:9000/users/1
```

### Index
```
curl http://localhost:9080/users
```

