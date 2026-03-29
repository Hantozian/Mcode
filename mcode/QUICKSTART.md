# 🚀 QUICK START Guide

## 1️⃣ Убедитесь что установлены:

```bash
go version      # Go 1.21+
psql --version  # PostgreSQL 12+
```

## 2️⃣ Создайте базу данных

```bash
psql -U postgres
CREATE DATABASE snippets;
\q
```

## 3️⃣ Установите зависимости

```bash
cd ~/mcode
go mod tidy
```

Если возникают проблемы с зависимостями:

```bash
go clean -modcache
go mod download
```

## 4️⃣ Запустите сервер

```bash
# Способ 1: С помощью go run
go run cmd/server/main.go

# Способ 2: Скомпилировать первым
go build -o bin/server cmd/server/main.go
./bin/server

# Способ 3: С помощью Makefile
make run
```

Ожидаемый вывод:
```
✅ Database connection established
✅ Migrations completed
✅ Added language: Go
✅ Added language: Python
...
🚀 Server starting on http://localhost:8080
📚 API available at http://localhost:8080/
```

## 5️⃣ Тестирование

```bash
# Проверка здоровья
curl http://localhost:8080/health

# Получить список языков
curl http://localhost:8080/languages

# Создать сниппет
curl -X POST http://localhost:8080/languages/go/snippets \
  -H "Content-Type: application/json" \
  -d '{"title":"Hello","filename":"h.go","content":"package main"}'
```

## 📚 Структура проекта

✅ **CMD** - запускаемые файлы  
✅ **INTERNAL** - внутренняя логика  
✅ **CONFIG** - конфигурация  
✅ **MODELS** - структуры данных  
✅ **HANDLERS** - HTTP обработчики  
✅ **DB** - работа с БД  

## 📖 Документация

- **README.md** - полное описание проекта
- **CODE_EXPLANATION.md** - пояснение каждой части кода
- **.env.example** - пример переменных окружения

## 🐛 Решение проблем

### Ошибка: "failed to connect to database"
Проверьте что:
- PostgreSQL запущен
- БД `snippets` создана
- Данные в `.env` правильные

### Ошибка: "missing go.sum entry"
Выполните:
```bash
go mod tidy
go mod download
```

### Порт уже занят
Измените PORT в `.env`:
```env
PORT=8081
```

## ✨ Готово!

Сервер работает и готов к использованию! 🎉
