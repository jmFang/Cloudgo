
## Use xorm to build database engine service 

使用xorm创建数据库引擎处理数据CRUD

# Main parts

## 1. import

```
  import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
  )
```

## 2. initialize

```
  func init() {
    //https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
    engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
    //db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
    if err != nil {
      panic(err)
    }
    //mydb = db
    myengine = engine
    engine.Sync2(new(UserInfo))
  }
  
```

## 3. Transaction


```
  func (*UserInfoAtomicService) Save(u *UserInfo) error {

    session := myengine.NewSession()
    defer session.Clone()

    //add Begin() before any action
    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(u)

    if err != nil {
      session.Rollback()
      return err
    }
    // add commit after all actions
    err = session.Commit()
    checkErr(err)
    return nil
  }
```


# Testing

```
sysuygm@localhost:~/golang-workspace/src/Cloudgo/goxorm$ go test -v
=== RUN   TestService
=== Test 'Save' Successfully!
=== Test 'FindAll' Successfully!. Finding result is : {42 sysu sdcs 2017-11-29 00:00:00 +0800 CST} 
=== Test 'FindById' Successfully!.--- PASS: TestService (0.09s)
PASS
ok  	Cloudgo/goxorm	0.098s

sysuygm@localhost:~/golang-workspace/src/Cloudgo/goxorm$ go test -v
=== RUN   TestService
=== Test 'Save' Successfully!
=== Test 'FindAll' Successfully!. Finding result is : {43 sysu sdcs 2017-11-29 00:00:00 +0800 CST} 
=== Test 'FindById' Successfully!.--- PASS: TestService (0.04s)
PASS
ok  	Cloudgo/goxorm	0.048s
sysuygm@localhost:~/golang-workspace/src/Cloudgo/goxorm$ 


```
