# Pull Request Management API

API для управления pull requests на Go + PostgreSQL.

## Старт
- Переменные окружения
```sh
cp .env.sample .env
```
- Запуск БД
```sh
docker-compose up --build db
```
- Установка зависимостей
```sh
go mod download
```
- Запуск (и перезапуск) приложения вместе с БД
```sh
make docker-restart
# Сервер на http://localhost:8080
```

## Реализованные функции

### 1. `/team/add` - Работает исправно

### 2. `/team/get` - Работает исправно

### 3. `/pullRequest/create` - Работает исправно 
(PR ID создаётся автоматически)

##  Нереализованные функции

### 1. `setIsActive` - Не работает
Требуется отладка логики обновления статуса пользователя в БД.

### 2. `pullRequest/merge` - Не реализовано
Изменение статуса PR на "merged". Недостаток времени.

### 3. `pullRequest/reassign` - Не реализовано
Изменение списка рецензентов. Недостаток времени.

### 4. `users/getReview` - Не реализовано
Получение всех PR, где пользователь рецензент. Недостаток времени.

### 5. Тесты всех функций. Недостаток времени

---
## Примеры запросов

### 1. `/team/add`

```sh
curl -X POST http://localhost:8080/api/v1/team/add \
 -H "Content-Type: application/json" \
  -d '{"team_name":"payments","members":[{"user_id":"u1","username":"Alice","is_active":true},{"user_id":"u2","username":"Bob","is_active":true}]}'
```

### 2. `team/get`

```sh
curl "http://localhost:8080/api/v1/team/get?team_name=payments"
```

### 3. `/pullRequest/create`
```sh
curl -X POST http://localhost:8080/api/v1/pullRequest/create \
  -H "Content-Type: application/json" \
  -d '{"pr_name": "Fix bug", "author_id": 1}'
```


