# Pull Request Management API

API для управления pull requests в Go + PostgreSQL.

## 🚀 Быстрый старт

```bash
go mod download
make docker-restart
# Сервер на http://localhost:8080
```

## 📚 API Endpoints

| Метод | Endpoint | Статус |
|-------|----------|--------|
| POST | `/api/v1/pullRequest/create` | ✅ Работает |
| GET | `/api/v1/pullRequest` | ✅ Работает |
| GET | `/api/v1/pullRequest/{id}` | ✅ Работает |
| PATCH | `/api/v1/users/{id}/setIsActive` | ⚠️ Не работает |
| POST | `/api/v1/pullRequest/{id}/merge` | ❌ Не реализовано |
| PATCH | `/api/v1/pullRequest/{id}/reassign` | ❌ Не реализовано |
| GET | `/api/v1/users/{id}/getReview` | ❌ Не реализовано |

## ❌ Нереализованные функции

### 1. `setIsActive` - Не работает
Требуется отладка логики обновления статуса пользователя в БД.

### 2. `pullRequest/merge` - Не реализовано
Изменение статуса PR на "merged". Недостаток времени.

### 3. `pullRequest/reassign` - Не реализовано
Изменение списка рецензентов. Недостаток времени.

### 4. `users/getReview` - Не реализовано
Получение всех PR, где пользователь рецензент. Недостаток времени.

## 🛠️ Стек

- Go 1.20+
- PostgreSQL 12+
- Docker Compose

## 📝 Пример запроса

```bash
curl -X POST http://localhost:8080/api/v1/pullRequest/create \
  -H "Content-Type: application/json" \
  -d '{"pr_name": "Fix bug", "author_id": 1, "reviewers_id": [2, 3]}'
```

---

**Avito Tech Internship - осень 2025**
