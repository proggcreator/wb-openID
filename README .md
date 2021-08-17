# OpenID авторизация

Цель: научиться работать с OpenID авторизацией и разграничением прав при работе с методами сервиса.

## Задание 1.

Изучить:
* Протоколы авторизации Basic, OAuth 2.0, 
* OpenId connect; 
* RFC стандартами JWT, JWE, JWS.
* https://openid.net/specs/openid-connect-core-1_0.html (официальная документация по OpenID connect)
* https://auth0.com/docs/tokens/access-tokens
* https://tools.ietf.org/html/rfc7519
* https://aaronparecki.com/oauth-2-simplified

## Задание 2.

Написать сервис, у которого будет путь – /api/v1/all, который будет отдавать произвольный JSON (например структуру из задания с сотрудниками).
Данные будут отдаваться только если у клиента будет достаточно прав.
Для проверки авторизации необходимо написать миддлвар, в нем сделать проверку на время жизни токена, и проверку на наличие права в токене, в данном случае на наличие права «world»

Дополнительная информация:

Тестовый URL авторизационного сервиса:
https://payments.wildberries.ru/authtest/.well-known/openid-configuration
авторизационные данные клиента:

```
ClientId: "finance-test-client"
ClientSecret:"testsecret"
```

Для сервиса: !!! устарело, т.к. верификация токена происходит на стороне сервиса без использования introspection endpoint авторизационного сервиса

```
ApiName (scope): "test-finance-microservices-api"
ApiSecret: "test-finance-microservices-api_secret"
```
