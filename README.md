# 重要

需要观察Go的完整日志输出： 

```export GRPC_GO_LOG_VERBOSITY_LEVEL=99```


```export GRPC_GO_LOG_SEVERITY_LEVEL=info```


使用方法：

首先在终端下暴露上面两个环境变量

进入 client 目录运行 go run main.go

进入 server 目录运行 go run main.go

停止server，或者kill -9 强制杀死，查看client表现

恢复server，查看client表现