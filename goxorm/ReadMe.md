
## Use xorm to build database engine service 

根据课程要求，使用xorm改写课程资料上使用的“entity - dao - service” 层次结构模型，同样地，创建数据服务并进行简单测试。xorm下不再需要编写复杂的——"dao"服务。

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
