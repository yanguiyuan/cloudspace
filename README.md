# cloudspace

## 构建相关命令

### cloufile

- 生成cloudfile rpc代码
 ```bash
 kitex -gen-path pkg -module github.com/yanguiyuan/cloudspace idl/cloudfile/cloudfile.thrift
 ```
- 生成cloudfile dal和model代码
```bash
go run ./cmd/gen/cloufile-gorm
```