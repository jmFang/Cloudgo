# Cloudgo

## web application practice

## 基本功能

1.支持静态文件服务

2.支持简单 js 访问

3.提交表单，并输出一个表格

4.对 /unknown 给出开发中的提示，返回码 5xx

## 基本的service

```
mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
mx.HandleFunc("/result", submit(formatter)).Methods("POST")
mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
mx.HandleFunc("/api/data", apiFetchDataHandler(formatter)).Methods("GET")
mx.PathPrefix("/api/").Handler(NotImplementedHandler())
```

## 静态文件访问

```
  sysuygm@localhost:~$ curl localhost:8080/static/css/main.css
  body {
    background-color: #fff7df;
  }

  #moretext{
    font-size: 20px;
    margin-top: 20px;

  }
  .span {
    width: 100px;
    margin-right: 20px;
  }
  sysuygm@localhost:~$ curl localhost:8080/static/
  <pre>
  <a href="css/">css/</a>
  <a href="images/">images/</a>
  <a href="js/">js/</a>
  </pre>

```
## 支持简单的js

```
$(document).ready(function() {
    $("#moretext").hide()
});
```

## 简单的表单提交和返回

![](https://github.com/jmFang/Cloudgo/blob/master/cloundgo-templates/assets/images/index.png)

![](https://github.com/jmFang/Cloudgo/blob/master/cloundgo-templates/assets/images/result.png)

## 状态码

![](https://github.com/jmFang/Cloudgo/blob/master/cloundgo-templates/assets/images/unkwon.png)

```
func NotImplementedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "501 Not Implemented.")
	})
}

```
