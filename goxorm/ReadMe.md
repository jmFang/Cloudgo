


```
sysuygm@localhost:~/golang-workspace/src/Cloudgo/goxorm$ go test -v
[xorm] [warn]  2017/11/29 22:03:12.390745 Table user_info Column username db default is , struct default is null
[xorm] [warn]  2017/11/29 22:03:12.390787 Table user_info Column departname db default is , struct default is null
[xorm] [warn]  2017/11/29 22:03:12.390796 Table user_info Column created db default is , struct default is null
[xorm] [warn]  2017/11/29 22:03:12.390805 Table user_info has column u_i_d but struct has not related field
[xorm] [warn]  2017/11/29 22:03:12.390812 Table user_info has column user_name but struct has not related field
[xorm] [warn]  2017/11/29 22:03:12.390818 Table user_info has column depart_name but struct has not related field
[xorm] [warn]  2017/11/29 22:03:12.390824 Table user_info has column create_at but struct has not related field
=== RUN   TestService
=== Test 'Save' Successfully!
=== Test 'FindAll' Successfully!. Finding result is : {42 sysu sdcs 2017-11-29 00:00:00 +0800 CST} 
=== Test 'FindById' Successfully!.--- PASS: TestService (0.09s)
PASS
ok  	Cloudgo/goxorm	0.098s
sysuygm@localhost:~/golang-workspace/src/Cloudgo/goxorm$ go test -v
[xorm] [warn]  2017/11/29 22:05:05.464789 Table user_info Column username db default is , struct default is null
[xorm] [warn]  2017/11/29 22:05:05.464821 Table user_info Column departname db default is , struct default is null
[xorm] [warn]  2017/11/29 22:05:05.464824 Table user_info Column created db default is , struct default is null
[xorm] [warn]  2017/11/29 22:05:05.464828 Table user_info has column u_i_d but struct has not related field
[xorm] [warn]  2017/11/29 22:05:05.464831 Table user_info has column user_name but struct has not related field
[xorm] [warn]  2017/11/29 22:05:05.464834 Table user_info has column depart_name but struct has not related field
[xorm] [warn]  2017/11/29 22:05:05.464836 Table user_info has column create_at but struct has not related field
=== RUN   TestService
=== Test 'Save' Successfully!
=== Test 'FindAll' Successfully!. Finding result is : {43 sysu sdcs 2017-11-29 00:00:00 +0800 CST} 
=== Test 'FindById' Successfully!.--- PASS: TestService (0.04s)
PASS
ok  	Cloudgo/goxorm	0.048s
sysuygm@localhost:~/golang-workspace/src/Cloudgo/goxorm$ 


```
