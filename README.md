# Testing with Golang

### Installing and Requirements
Go Version 1.6

##### Frameworks 
* github.com/jinzhu/gorm
* github.com/ant0ine/go-json-rest/rest
* github.com/shengkehua/xlog4go

##### Run 
```
cd $GOPATH/src
git clone https://github.com/viniciusfpe/go-gorm.git
go get
go run main.go

```


### Tests
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


