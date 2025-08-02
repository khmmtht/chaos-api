## Docker
```
docker pull khemmathat141/chaos-api:simple-latest
```


## Notes
In simple version will not use Api-Token and Project so project-will be any value

## Environment Variables

| Name | Value |
| --- | --- |
| `DRIVER` | `mongodb` or `cache` (default = `cache`) |
| `MONGODB_URI` | `mongodb://root:example@localhost:27017` (example) |

## Simulate Endpoints

| Endpoint                          | Description |
|-----------------------------------| --- |
| `GET /api/v1/simulate/delay/{ms}` | Simulate network delay |
| `GET /api/v1/simulate/error/{code}` | Return specific error code |

## Chaos:
Required request header `Project-Id` (example `Project-Id=774617ce-f0b2-4649-80ec-bd9c7dceabc7`)

| Endpoint | Description |
| --- | --- |
| `GET /api/v1/chaos/status/{service_name}` | Get current chaos configuration |
| `POST /api/v1/chaos/configure` | Upsert chaos settings |
| `POST /api/v1/chaos/trigger/{service_name}` | Default trigger by configure |
| `POST /api/v1/chaos/reset/{service_name}` | Reset settings |


Example `POST /api/v1/chaos/configure` body
```json
{
  "name": "notification",
  "mode": "response",
  "value": "400",
  "response": "{ \"error\": \"bad request from mock\" }"
}
```

| Mode | Description |
| --- | --- |
| `latency` | Artificial delay in responses (in ms) |
| `response` | Hang the response indefinitely or past timeout window |
