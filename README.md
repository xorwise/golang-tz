# Тестовое задание на должность Golang-Junior разработчика.

В данном тестовом задании реализован весь необходимый функционал, который включает в себя два REST-маршрута:
- `GET /api/login/{GUID}` - который возвращает пару access_token и refresh_token
- `POST /api/refresh` - который обновляет access_token по refresh_token, в теле запроса передается refresh_token

Из сторонних библиотек используются `github.com/golang-jwt/jwt` и `go.mongodb.org/mongo-driver`

Так как в задании было сказано использовать GUID (он же UUID) в качестве индентификатора пользователя, то для этого был реализован кастомный тип UUID.

Приложение написано с использованием **Clean Architecture**.

## Порядок запуска
- Перед запуском убедитесь, что у вас запущен MongoDB.
- Клонируем репозиторий: `git clone https://github.com/xorwise/golang-tz.git`
- Запускаем приложение в из папки bin. (Для Windows: `./bin/test.exe`, для Linux: `./bin/test`)

## Описание решения
В качестве подключения к MongoDB используется URI: `mongodb://127.0.0.1:27017`

При желании можно поменять в `./bootstrap/env` эту ссылку, и другие константы, используемые в приложении.

После чего можно запускать приложение на версии **Golang 1.22**:
- `go run ./cmd/`

P.S. Не стал писать юнит-тесты к этому приложению, т.к. этого не требуется в задании. Однако могу предоставить несколько интеграционных тестов:
- `curl -X GET http://localhost:8080/api/login/{Любой GUID}`
- `curl -X POST -H "Content-Type: application/json" -d '{"refresh_token": "'{Полученный refresh_token'"}' http://localhost:8080/api/refresh`

В задании не указан способ добавления новых пользователей с GUID в базу данных, поэтому они добавляются (изменяются при наличии) при каждом запросе с новым GUID.

Refresh токен хранится, как поле пользователя в формате HS256 хэша.

Алгоритм кодирования Access токена: **HS512**, алгоритм кодирования Refresh токена: **HS256**

Телеграм для связи: [https://t.me/xorwise](https://t.me/xorwise)

Надеюсь вам понравится мое решение.
