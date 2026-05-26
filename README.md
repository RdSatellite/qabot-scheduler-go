# Scheduler-go

Generate pb:

```bash
# Add the module name prefix to your terminal call to override the invalid path
protoc --go_out=. --go_opt=Mproto/scheduler.proto=proto/pb \
       --go-grpc_out=. --go-grpc_opt=Mproto/scheduler.proto=proto/pb \
       proto/scheduler.proto
```
