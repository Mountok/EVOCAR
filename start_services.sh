#!/bin/bash

# Запуск контейнера с PostgreSQL
docker run --name=evocar-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

# Проверяем, успешно ли стартовал контейнер
if [ $? -ne 0 ]; then
    echo "Ошибка при запуске PostgreSQL"
    exit 1
fi

# Ожидание старта БД (чтобы миграция не запустилась слишком рано)
echo "Ожидание запуска PostgreSQL..."
sleep 5

# Выполняем миграции
migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable" up

# Проверяем, успешно ли прошла миграция
if [ $? -ne 0 ]; then
    echo "Ошибка при выполнении миграций"
    exit 1
fi

# Запуск Redis
docker run --name redis-server -d -p 6379:6379 redis

# Проверяем успешность запуска Redis
if [ $? -ne 0 ]; then
    echo "Ошибка при запуске Redis"
    exit 1
fi

echo "Все сервисы успешно запущены!"
