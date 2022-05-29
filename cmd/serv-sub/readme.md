To launch serv-sub, use the following command:

```bash
dapr run --app-port 6001 --app-id order-processor --app-protocol http --dapr-http-port 3501 --components-path ../../components -- go run serv-sub.go
```