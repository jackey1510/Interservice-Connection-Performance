
# Interservice connection Performance

Measure the performance of various protocols for interservice connection

<img src="https://storage.googleapis.com/poc-interservice-connection/poc.drawio.png"/>


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`ELASTIC_APM_SERVER_URL` Elastic APM Credentials

`ELASTIC_APM_SECRET_TOKEN` Elastic APM Credentials

`ELASTIC_APM_SERVICE_NAME` Elastic APM Credentials

`CERT_PATH` CA certificate path

`KEY_PATH` CA key path

Put your certs in /certs/ca
## Run Locally

Docker Compose

```bash
docker compose up
```

Go Build & Go Run

```bash
go build
go run server
```
## API Reference

### Hello World

```http
  GET /
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
|  N/A| N/A | return hello world from `${APP_NAME}` |

### Proxy Request

```http
  GET /get
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `url`      | `string` | **Required**. URL to send request to |
| `protocol` | `string` | **Required**. `HTTP/1.1` \| `HTTP/1.1TLS` \| `HTTP/2` \| `HTTP/3`|
| `requestCount` | `string` | **Required**. No of requests|
#### Example

Get Hello World from http1

`GET 'https://localhost:8080/get?url=http://http1:8080&protocol=HTTP/1.1&requestCount=1'`

Get Hello World from http2

`GET 'https://localhost:8080/get?url=https://http2:8080&protocol=HTTP/2&requestCount=1'`

Get Hello World from http3

`GET 'https://localhost:8080/get?url=https://http3:8080&protocol=HTTP/3&requestCount=1'`

Get Hello World from http3 via http2

`GET 'https://localhost:8080/get?url=https://http2:8080/get?url=https://http3:8080%26protocol=HTTP/3%26requestCount=1&protocol=HTTP/2&requestCount=1'`
