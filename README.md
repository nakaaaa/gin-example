# gin-example
ginを色々と使ってみる

# 確認
`curl`を使用する  
```bash
curl localhost:8080/ping
```

responseは以下のとおり
```
{"message":"pong"}
```

# CRUD
## 登録・更新
- `/user`
- Method: Post
## 取得
- `/user/:id`
- Method: Get
## 削除
- `/user/:id`