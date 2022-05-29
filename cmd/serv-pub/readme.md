To start serv-pub use 

```bash
dapr run --app-id checkout --app-protocol http --dapr-http-port 3500 --components-path ../../components -- go run serv-pub.go
```