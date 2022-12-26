# [RBC Parser](https://github.com/6WhoAmI6/rbk-parser)

---

## Table of Contents

- [Start service](#start-service)
- [Parse news from RBC](#parse-news-from-rbc)
  - [Endpoint for parse RBC](#endpoint-for-parse-rbc)
  - [Response from service](#response-from-service)
    - [Data](#data)
      - [News](#news)
- [Status](#status)
  - [Endpoint for get status of service](#endpoint-for-get-status-of-service)
- [Root](#root)
- [Scalar Value Types](#scalar-value-types)

### Start service

To start a service you need to download Go from [oficial site](https://go.dev/dl/).

After that just download the service from repository

```shell
    git clone https://github.com/6WhoAmI6/rbk-parser.git
```

In folder with project performe next steps:

For Windows

```shell
    go build
    .\rbc-parser.exe localhost:8080
```

For MacOS or Linux

```shell
    go build
    ./rbc-parser localhost:8080
```

Now you can use it with [http://localhost:8080/](http://localhost:8080/)

### Parse news from RBC

The method to parse news from rbc.ru.

#### Endpoint for parse RBC

```shell
    http://localhost:8080/rbc-parse
```

#### Response from service

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) | success/unsuccess | Success status when request performed without any error |
| error | [string](#string) |  | When request was successful error is empty string |
| data | [Data](#data) | repeated | Data with news. If request was unsuccessful the data is empty list|

#### Data

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| main_news | [News](#news) |  | Main news of the day |
| top_news | [News](#news) | repeated | Top news of the day |
| central_news | [News](#news) | repeated | All other news |

#### News

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | News URL |
| title | [string](#string) |  | News title |

## Status

The method to get status of the service.

### Endpoint for get status of service

```shell
    http://localhost:8080/status
```

## Root

```shell
    http://localhost:8080/
```

## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
