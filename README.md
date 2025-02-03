## RUN API

### Запус Postgres

Скачиваем контейнер: 

`docker pull postgres`

Запуск контейрена:

`docker run --name=evocar-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres`

Убедимся что контейнер запущен: 

`docker ps`

Пожключения к контейнеру с бд:

`docker exec -it container_id /bin/bash`

---
## Миграция 

Для установки утелите: `sudo port install go-migrate`

Создание файлов для миграции: 

`migrate create -ext sql -dir ./schema -seq init`

Для миграции: 

`migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up`

--- 
## Запус Redis
Для установки: 

`docker pull redis`

Для запуска кеш базы данных:

`docker run --name redis-server -d -p 6379:6379 redis`

Подключение к контейнеру

`docker exec -it redis-server redis-cli`
