To launch serv-sub, use the following command:

```bash
dapr run --app-port 6001 --app-id serv-sub --app-protocol http --dapr-http-port 3501 --components-path ../../components -- go run serv-sub.go
```

To check status of Dapr applications run the Dapr dashboard:

```bash
dapr dashboard
```

To publish a message use application serv-pub or send a http POST request to localhost:app-port (app-port is specified when launching the Dapr application):

```
{
    "id": "2222",
    "type": "serv-pub.v1",
    "source": "golang-dapr",
    "specversion": "1.0",
    "datacontenttype": "application/json",
    "data": {
        "type": "orders",
        "id":"123",
        "description": "Hello, a new order!",
        "price": 23.23
    }
}
```
