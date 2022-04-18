
# Interservice connection Performance

Measure the performance of various protocols for interservice connection

<img src="https://00f74ba44b76e8b3ec88c72cb647025302f35f2ddf-apidata.googleusercontent.com/download/storage/v1/b/poc-interservice-connection/o/poc.drawio.png?jk=AFshE3Waqpg8t45L9T5BYX_Ds4BvzCNfi_Q6fLhHwvmRMgwAEbgT-sV1ZC-6uhFpIojKTzB6jqW4L5ZCjR59-UVmmvbufsnCCUbqEZ0mDuxOuEuPmaPPwpjMAbbo-WCpyOWrBtRYEFpmpoa5_FQ2bcKQ56ZKhYX5XsJ5Hg1-b-_uMoEv1acwpSlZ7CwCvgVgZTVO0HVBaRolDYEW9f7ErY0JjWc-VlIVhx1CqBC0gPQLcrgKFoA4YZb-ixcUdee5PI9EN4PFu8gPcBISek-lzWhrXNJLp8IBAg1Q-x2VU97X9cbi7S104yhIVraQ9S6CROajdZSTZ__P0cF9r_11BX2kNZNg6eyWSDUrhO1SBzpZuy8_k1W6_eVGEdk8QHIgAEW-x1E_SVzqL1qbxx4hhvGvPpAXxtkyW8teLl7VtIPt7GPIwW635or_rkq3stzMDeP43pGgYqmrh4dGwhr0IZzJ_edMQjxTb0ZSqTLuPbe4OYMQv68Nt_Da_Dhr8X-oMzMIOVbEQsv1PJAPbI72dT4qu0S_3CEjhQ8gWA9D83Of4GoS081eyQnbUO5VvDWOHuEk_6WfwnBYvRY1wnH8dFyN9GeFfzC-Udn7Q-S5YTRiQQdrthWQnQt9zPplpDz0D7aHANNV6nNyfo_NhqJ4uxqNfygy6fodpTlddWpLjqPf0lNaTYbVTYYsXG3xa3jXwCe6qHs8Aq_fGBTnQxdARWSG3i34g1Et-LT-2kQqVr9ZwqasfQL59XjMfHGGrdvnCZwXuhGDsNbwpGZAnGa5SYZ-aLWAT_XMMnUT2HxftQarTVsJ0hzyImNZ7-u_W6Trgr665Qc6YZ9qXV7o2zGo4bRCRfpbcR_PJWXhWtV4M2mfAw5kp6BJIX9XIuvPJamSCA_uizkuiBZPak6jrDgXCc_FQZW-KpBrfQg6illTFSYMk0M8jO4h-ZmVCQcVhYnjfrbE04vzgSl-OyfuS7woTptMbalBdq47nSyTnoXannuLL7Sz6VeYeTzoUz5AIw&isca=1"/>


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
