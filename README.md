# go-gorm
Teste com Gorm e Go-Json-Rest

```
Insert
curl -i -H 'Content-Type: application/json' \
    -d '{"Message":"this is a test"}' http://127.0.0.1:8080/reminders

Get specific
curl -i http://127.0.0.1:8080/reminders/1

Get All
curl -i http://127.0.0.1:8080/reminders

Update
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Message":"is updated"}' http://127.0.0.1:8080/reminders/1

Delete
curl -i -X DELETE http://127.0.0.1:8080/reminders/1
```


