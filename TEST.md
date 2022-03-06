# エンドポイントテスト
* バリデーションはOK
* CompanyのバリデーションもOK
* 同じメールアドレス、uidでerr返ってくる → 同じものを挿入したときidは進む


## 正しいパターン

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxxx","name": "tomoki", "email": "tt@gmail.com", "company": {"name": "yy_company"}}'
```


```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxyy","name": "tomoki", "email": "yy@gmail.com", "company": {"name": "yy_company"}}'
```

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxzz","name": "tomoki", "email": "zz@gmail.com", "company": {"name": "yy_company"}}'
```

## uidなし

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"name": "tomoki", "email": "tt@gmail.com", "company": {"name": "yy_company"}}'
```

## 不正email

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxxx","name": "tomoki", "email": "ttaidewdwem", "company": {"name": "yy_company"}}'
```

## nameなし

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxzz", "email": "zz@gmail.com", "company": {"name": "yy_company"}}'
```

## companyなし(パターン1)

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxzz","name": "tomoki", "email": "zz@gmail.com", "company": {}}'
```

## comapanyなし(パターン2)

```
curl -X POST http://localhost:9000/user/create \
-H 'Content-Type: application/json' \
-d '{"uid":"xxxxxxxxxxxxxxxxxxxxxxxxxxee","name": "tomoki", "email": "rrr@gmail.com"}'
```
