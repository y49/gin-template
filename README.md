# Go Gin Template
## Features

- Gin
- RESTful API
- Gorm
- Swagger
- Logging
- Graceful restart or stop 
- App configurable
- Cron
- multi-language
- Error warning（Email notification）
- Jaeger

## Conf

项目 configs/config.yaml（利用viper实现热更新）
```config
Server:
    RunMode:  debug
    HttpPort: 9090
    ReadTimeout: 60
    WriteTimeout: 60
App:
    ServiceName: gin-template
    DefaultPageSize: 10
    MaxPageSize: 100
    DefaultContextTimeout: 60
    LogSavePath: storage/logs
    LogFileName: app
    LogFileExt: .log
    UploadSavePath: storage/uploads
    UploadServerUrl: http://127.0.0.1:9090/static
    UploadImageMaxSize: 5 # MB
    UploadImageAllowExts:
        - .jpg
        - .jpeg
        - .png
Email:
    Host: smtp.qq.com
    Port: 465
    UserName: xxx@qq.com
    Password: xxxx
    IsSSL: true
    From: xxx@qq.com
    To:
        -  xxx@qq.com
JWT:
    Secret: test123
    Issuer: test-service
    Expire: 7200
Database:
    DBType: mysql
    UserName: root
    Password: 123456
    Host: 127.0.0.1:3306
    DBName: test
    TablePrefix: test_
    Charset: utf8
    ParseTime: True
    MaxIdleConns: 10
    MaxOpenConns: 30

Jaeger:
    Host: 127.0.0.1:6831
```

- JWT: gin-template\internal\model\auth.go //可修改表和字段


## Run
指定端口模式和配置
```bash
 go run main.go  -port=8001 -mode=release -config=configs/ go run main.go  -port=8080 -mode=release -config=configs    
```
带有执行信息方式：
```bash
go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0" //buildTime 运行时间 buildVersion 版本 
```


