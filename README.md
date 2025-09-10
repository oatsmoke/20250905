## REST-сервис для агрегации данных об онлайн-подписках пользователей

### Endpoints:

`POST /subscriptions` - Создать подписку

`GET /subscriptions` - Получить список подписок

`GET /subscriptions/{id}` - Получить подписку по ID

`PUT /subscriptions/{id}` - Обновить подписку по ID

`DELETE /subscriptions/{id}` - Удалить подписку по ID

`GET /subscriptions/total` - Получить сумму подписок за период

`GET /swagger/` - Swagger UI

### Environments:

`HTTP_PORT` - http порт на котором слушает сервер. По умолчанию: `8080`

`POSTGRES_DSN` - Строка подключения к базе данных. По умолчанию: 
`postgres://root:password@localhost:5432/postgres?sslmode=disable`

### Makefile:

`migrate` - Создать миграцию. Через параметр `NAME=` можно указать имя миграции.

`swag` - Сгенерировать документацию.

### Запуск:

```bash
  docker compose up -d
```
