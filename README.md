# IdeaForge

Каркас веб-платформы на Go и Gin. Страницы-заглушки используют встроенные HTML-шаблоны с заголовками и переходами между разделами.

## Запуск

Запустите из корня репозитория с Go 1.27.2 или новее.

```sh
go run .
```

Откройте [главную страницу](http://localhost:8080/) или [каталог](http://localhost:8080/catalog). Сервер слушает порт 8080. Переменная `HTTP_ADDR` задаёт другой адрес, например `HTTP_ADDR=127.0.0.1:8090 go run .`.

Команды для Docker

```sh
docker build -t ideaforge .
docker run --rm -p 127.0.0.1:8080:8080 ideaforge
```

Команды для проверки сборки и кода

```sh
go test ./...
go vet ./...
```

## Структура

| Путь | Ответственность |
| --- | --- |
| [main.go](main.go) | Запуск HTTP-сервера |
| [internal/app](internal/app) | Сборка Gin-роутера, подключение модулей и `GET /health` |
| [internal/web](internal/web) | Маршруты и шаблоны страниц-заглушек |
| [internal/accounts](internal/accounts) | Регистрация, вход, сессии, роли, профиль и настройки |
| [internal/projects](internal/projects) | Главная, поиск, каталог и фильтры, проекты, черновики, избранное и прогресс |
| [internal/solutions](internal/solutions) | Решения, версии, отметки полезности и оценки |
| [internal/discussions](internal/discussions) | Комментарии, ответы и реакции |
| [internal/teams](internal/teams) | Команды, открытые роли, заявки и участники |
| [internal/notifications](internal/notifications) | Уведомления и предпочтения пользователя |
| [internal/collections](internal/collections) | Подборки проектов в сообществе и учебные подборки наставников |
| [internal/moderation](internal/moderation) | Жалобы, проверка публикаций, категории, теги, аудит и метрики |
| [internal/sources](internal/sources) | Внешние источники, происхождение кейсов, черновики и импорт |
| [internal/intelligence](internal/intelligence) | Рекомендации, помощник автора, поиск похожих проектов, оценки сложности и актуальности |
| [internal/platform/scaffold](internal/platform/scaffold) | Общие типы пагинации, ошибок и маркер заглушки |

В каждом модуле `model.go` описывает сущности и состояния. Типы запросов и ответов находятся рядом с сущностями либо в `contracts.go`. Операции из `service.go` возвращают `scaffold.ErrNotImplemented`. Интерфейсы из `repository.go` задают точки подключения хранилищ и внешних провайдеров. В `http.go` регистрируются маршруты.

## Заглушки интерфейса и API

Кнопка «Поиск» открывает `/search`, а ссылка «Создать проект» ведёт на `/projects/new`. Эти и остальные страницы показывают заголовок и переходы к связанным разделам. Все ручки под `/api/v1` возвращают `501 Not Implemented` и заголовок `X-IdeaForge-Mode: mock`.

Примеры запросов

```sh
curl -i http://localhost:8080/api/v1/projects
curl -i -X POST http://localhost:8080/api/v1/projects
```

| Область | Пути под `/api/v1` |
| --- | --- |
| Поиск и каталог | `/home`, `/search/suggestions`, `/catalog/options`, `/projects` |
| Проекты | `/projects/:projectId`, подмаршруты `preview`, `publish`, `favorite`, `work` и `work/stages/:stageId` |
| Аккаунты | `/auth/register`, `/auth/login`, `/auth/refresh`, `/auth/logout`, `/me`, `/me/settings`, `/users/:userId`, `/me/projects` |
| Решения и обсуждения | `/projects/:projectId/solutions`, `/solutions/:solutionId`, `/projects/:projectId/comments`, `/solutions/:solutionId/comments`, `/comments/:commentId` |
| Команды и сообщество | `/projects/:projectId/teams`, `/teams/:teamId/applications`, `/me/teams`, `/collections` |
| Уведомления | `/me/notifications`, `/me/notifications/:notificationId/read`, `/me/notifications/settings` |
| Модерация и источники | `/reports`, `/admin/moderation`, `/admin/catalog`, `/admin/sources`, `/admin/imports`, `/admin/analytics/usage` |
| Рекомендации и помощник | `/me/recommendations`, `/intelligence/author-assistance`, `/intelligence/similar-projects`, `/intelligence/difficulty`, `/intelligence/relevance` |
