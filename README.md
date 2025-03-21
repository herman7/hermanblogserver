# Herman Blog Server with Go

## Dev
1. 启动本地swagger
   1. 安装swagger工具
   ```bash
   go install github.com/go-swagger/go-swagger/cmd/swagger@latest
   ```
   2. 本地运行
   ```bash
   swagger serve -F=swagger --no-open --port 65534 ./api/openapi/openapi.yaml
   ```
2. OpenApi文档中文：https://fishead.gitbook.io/openapi-specification-zhcn-translation/3.0.0.zhcn

## Tool

- air 热加载工具 https://github.com/air-verse/air
