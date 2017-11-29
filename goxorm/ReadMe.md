
## 目录

1. 根据课程要求，使用`xorm`改写课程资料上使用的“entity - dao - service” 层次结构模型（（xorm下不再需要编写复杂的——"dao"服务）

2. 创建数据服务并进行`单元测试`

3. `travis `在线测试

4. web 服务 `curl` 测试

5. `ab` 性能测试

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
## 1. testing 单元测试
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
## 2. 添加 .travis.yml 在线测试

```
Cloning into 'jmFang/Cloudgo'...
remote: Counting objects: 85, done.
remote: Compressing objects: 100% (73/73), done.
remote: Total 85 (delta 21), reused 53 (delta 7), pack-reused 0
Unpacking objects: 100% (85/85), done.
$ cd jmFang/Cloudgo
$ git checkout -qf 520f9d74c72bb29b5c6d18571fed550e671ef6f6
Updating gimme
144.24s$ GIMME_OUTPUT="$(gimme tip | tee -a $HOME/.bashrc)" && eval "$GIMME_OUTPUT"
go version devel +21672b36eb Wed Nov 29 05:17:03 2017 +0000 linux/amd64
$ export GOPATH=$HOME/gopath
$ export PATH=$HOME/gopath/bin:$PATH
$ mkdir -p $HOME/gopath/src/github.com/jmFang/Cloudgo
$ rsync -az ${TRAVIS_BUILD_DIR}/ $HOME/gopath/src/github.com/jmFang/Cloudgo/
$ export TRAVIS_BUILD_DIR=$HOME/gopath/src/github.com/jmFang/Cloudgo
$ cd $HOME/gopath/src/github.com/jmFang/Cloudgo
0.01s
$ gimme version
v1.2.0
$ go version
go version devel +21672b36eb Wed Nov 29 05:17:03 2017 +0000 linux/amd64
...
...
```
## 3. curl web服务测试

```
  sysuygm@localhost:~$ curl -d "username==ooo&departname=1" http://localhost:8080/service/userinfo
  {
    "UID": 47,
    "UserName": "=ooo",
    "DepartName": "1",
    "CreateAt": "2017-11-29T23:17:03.421626349+08:00"
  }
  sysuygm@localhost:~$ curl http://localhost:8080/service/userinfo?userid=47
  {
    "UID": 47,
    "UserName": "=ooo",
    "DepartName": "1",
    "CreateAt": "2017-11-29T00:00:00+08:00"
  }
  sysuygm@localhost:~$ 

```
## 4. ab 压力测试

### 4.1 使用XORM下，数据库并发查询的压力测试

```
sysuygm@localhost:~$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=47
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=47
Document Length:        102 bytes

Concurrency Level:      100
Time taken for tests:   0.313 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      235000 bytes
HTML transferred:       102000 bytes
Requests per second:    3199.68 [#/sec] (mean)
Time per request:       31.253 [ms] (mean)
Time per request:       0.313 [ms] (mean, across all concurrent requests)
Transfer rate:          734.30 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.6      0       6
Processing:     0   29  36.4     19     164
Waiting:        0   29  36.3     18     163
Total:          0   30  37.6     19     170

Percentage of the requests served within a certain time (ms)
  50%     19
  66%     25
  75%     29
  80%     32
  90%     84
  95%    144
  98%    153
  99%    160
 100%    170 (longest request)
sysuygm@localhost:~$ 

```

### 4.2 使用XORM下，数据库并发存储的压力测试


```
sysuygm@localhost:~$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?"username==ooo&departname=1"
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?username==ooo&departname=1
Document Length:        3537 bytes

Concurrency Level:      100
Time taken for tests:   0.352 seconds
Complete requests:      1000
Failed requests:        15
   (Connect: 0, Receive: 0, Length: 15, Exceptions: 0)
Non-2xx responses:      1000
Total transferred:      3652985 bytes
HTML transferred:       3536985 bytes
Requests per second:    2837.04 [#/sec] (mean)
Time per request:       35.248 [ms] (mean)
Time per request:       0.352 [ms] (mean, across all concurrent requests)
Transfer rate:          10120.77 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.8      0       4
Processing:     0   34  38.3     11     202
Waiting:        0   34  38.1     11     202
Total:          0   35  38.3     11     203
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     43
  75%     78
  80%     82
  90%     95
  95%     97
  98%    107
  99%    113
 100%    203 (longest request)
sysuygm@localhost:~$ 

```
