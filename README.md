##  基于go简单的上传图片转base64

```shell
    go run main.go
```

上传地址localhost:8080/upload

报文格式，名称file
```
POST /upload HTTP/1.1
Host: localhost:8080
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW
User-Agent: PostmanRuntime/7.17.1
Accept: */*
Cache-Control: no-cache
Postman-Token: 9a0874a7-06e6-4812-8a92-8feaeda98785,0b3fe5b1-9b6d-4bec-98cc-e4fb3135f099
Host: localhost:8080
Accept-Encoding: gzip, deflate
Content-Length: 3865
Connection: keep-alive
cache-control: no-cache


Content-Disposition: form-data; name="file"; filename="/C:/Users/Pictures/_编组_.png


------WebKitFormBoundary7MA4YWxkTrZu0gW--
```