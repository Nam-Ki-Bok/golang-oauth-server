# NIMS - OAuth server

NIMS OAuth server 입니다.

Public API 인증, 인가를 담당하고 있습니다.

**자세한 정보는 [wiki](https://gitlab.gabia.com/infradev/infradev-practice/-/wikis/NIMS-Core-Process-(OAuth))를 참조해 주세요.**

## Stack

- ##### [go](https://github.com/golang/go) (v1.17.6)

- ##### [gin](https://github.com/gin-gonic/gin) (v1.7.7)

- ##### [go-oauth2](https://github.com/go-oauth2/oauth2) (v4.4.3)

## Database

- ##### Maria (v10.6.5)

- ##### Redis (v3.2.1)

- ##### Mongo (v5.0.6)

**자세한 Database 용도는 [wiki](https://gitlab.gabia.com/infradev/infradev-practice/-/wikis/NIMS-Core-Process-(OAuth))를 참조해 주세요.**

## Access Control List

ACL에 사용자가 정의한 client를 저장합니다.

| client_id  | client_secret  | client_ip | grant_type        | scope                          |
| ---------- | -------------- | --------- | ----------------- | ------------------------------ |
| foo_client | SHA256(secret) | 127.0.0.1 | client_credential | GET_resource POST_resource ... |

**자세한 client 검증 과정은 [wiki](https://gitlab.gabia.com/infradev/infradev-practice/-/wikis/NIMS-Core-Process-(OAuth))를 참조해 주세요.**

## API

#### Token

##### 👉 GET /oauth/token

- Request (Basic Auth)

  ACL에 저장되어 있는 **client id, client secret**을 입력해야 합니다.

```json
{
    "client_id": "foo_client",
    "client_secret": "bar"
}
```

- Response

  유효한 client 인증 요청으로 판단 된다면, **access token**이 포함된 발급 기록을 반환합니다.

```json
{
    "client_id": "foo_client",
    "access_token": "your_access_token",
    "scope": [
        "Method_resource",
        "Method_resource",
        "Method_resource",
    ],
    "created_in": "yyyy-mm-dd hh:mm:ss",
    "expires_in": "created_in + 2 hours"
}
```

**자세한 access token 검증 과정은 [wiki](https://gitlab.gabia.com/infradev/infradev-practice/-/wikis/NIMS-Core-Process-(OAuth))를 참조해 주세요.**

## Scope

Client가 접근 할 수 있는 자원은 Scope를 통해 제한합니다.

Scope는 "Method_resource" 형태로 이루어집니다.

> ex) GET_stock, POST_stock, PUT_stock, DELETE_stock

**자세한 Scope 검증 과정은 [wiki](https://gitlab.gabia.com/infradev/infradev-practice/-/wikis/NIMS-Core-Process-(OAuth))를 참조해 주세요.**

## Enviroment

민감한 정보는 .env 파일에 저장합니다.

```go
MARIA_CONNECTION_INFO="user:password@tcp(ip:port)/database"

REDIS_CONNECTION_INFO="ip:port"

TOKEN_END_POINT="http://localhost:1054/token"

MONGO_URI="mongodb://ip:port"
MONGO_USER="user"
MONGO_PW="password"
```

## Todo

- Badge 추가
- Testcode 추가
- Dockerfile 추가
- MIT License 추가
