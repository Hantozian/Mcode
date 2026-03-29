# 📚 Mcode - Backend

**Простой и надежный сервис для сохранения, организации и управления кусочками кода**

---

## 📖 Что это такое?

Представьте, что вы разработчик и много раз пишете одинаковый код. Вместо того, чтобы каждый раз писать это заново или искать в интернете, вы можете:

1. **Сохранить** этот код в нашем сервисе
2. **Организовать** его по языкам программирования
3. **Получить** его когда нужно заново

Этот проект — это backend для такого сервиса.

---

## 🎯 Что может делать этот сервис?

✅ Управлять языками программирования  
✅ Сохранять сниппеты кода  
✅ Получать список всех сохраненных сниппетов  
✅ Просматривать содержимое сниппета  
✅ Удалять ненужные сниппеты  

**Пример:** Вы добавили Python, затем сохранили 3 функции на Python, затем просмотрели и удалили одну.

---

## 🛠 Технологии (Просто объясняем)

| Технология | Что это | Зачем нужно |
|-----------|--------|-----------|
| **Go** | Язык программирования | Это язык, на котором написан наш сервер. Он очень быстрый и надежный |
| **Gin** | Веб-фреймворк | Инструмент для создания API (интерфейса взаимодействия). Как конструктор для веб-приложений |
| **PostgreSQL** | База данных | Место, где хранятся все наши сниппеты и языки. Как большая книга с табличками |
| **GORM** | ORM библиотека | Инструмент для общения с базой данных через код Go. Переводит команды Go в понятный БД язык |

**Простой пример:**
- Пользователь → отправляет запрос → Gin получает → Обрабатывает → GORM говорит PostgreSQL что-то сохранить → БД сохраняет

---

## 📁 Структура проекта (где что находится)

```
mcode/
├── cmd/
│   └── server/
│       ├── main.go              # Главный файл сервера
│       └── main_test.go          # Тесты
├── internal/
│   ├── db/
│   │   └── db.go                # Подключение к базе данных
│   ├── models/
│   │   ├── language.go          # Модель языка программирования
│   │   ├── snippet.go           # Модель сниппета кода
│   │   └── errors.go            # Ошибки
│   └── handlers/
│       ├── language_handlers.go # Функции для работы с языками
│       └── snippet_handlers.go  # Функции для работы со сниппетами
├── config/
│   └── .env                     # Настройки подключения к БД
├── go.mod                       # Зависимости проекта
└── README.md                    # Этот файл
```

**Кратко:**
- `cmd/` - главные файлы для запуска
- `internal/` - внутренняя логика
- `config/` - настройки
- `go.mod` - список нужных библиотек

---

## 🚀 Как установить и запустить

### Шаг 1: Проверка предусловий

Убедитесь, что у вас установлены:
- **Go** (версия 1.21 или выше) - скачать: https://golang.org/dl/
- **PostgreSQL** - скачать: https://www.postgresql.org/download/

Проверка:
```bash
go version          # Должно показать Go версию
psql --version      # Должно показать PostgreSQL версию
```

### Шаг 2: Подготовка базы данных

Откройте терминал и создайте базу:

```bash
# Подключитесь к PostgreSQL через psql
psql -U postgres

# Выполните эту команду в psql:
CREATE DATABASE snippets;

# Выход из psql
\q
```

**Что произошло?** Мы создали пустую базу данных с названием "snippets", куда наш сервер будет сохранять данные.

### Шаг 3: Клонирование/подготовка проекта

```bash
cd ~/mcode
```

### Шаг 4: Установка зависимостей

```bash
go mod download
```

**Что произошло?** Go скачал все необходимые библиотеки (Gin, GORM и т.д.)

### Шаг 5: Проверка конфигурации

Откройте файл `config/.env` и убедитесь, что данные совпадают с вашей БД:

```env
DB_HOST=localhost          # Где находится БД (обычно локально)
DB_PORT=5432              # Порт по умолчанию для PostgreSQL
DB_USER=postgres           # Пользователь БД
DB_PASSWORD=postgres       # Пароль пользователя
DB_NAME=snippets           # Название нашей БД
DB_SSLMODE=disable         # SSL отключен (для разработки нормально)
PORT=8080                  # На каком порту запустится наш сервер
```

### Шаг 6: Запуск сервера

```bash
go run cmd/server/main.go
```

Если все хорошо, вы увидите:
```
✅ Database connection established
✅ Migrations completed
✅ Added language: Go
✅ Added language: Python
...
🚀 Server starting on http://localhost:8080
📚 API available at http://localhost:8080/
```

**Поздравления! Сервер работает!** ✨

---

## 🌐 API (как общаться с сервером)

### Проверка здоровья сервера (Ping)

```bash
curl http://localhost:8080/health
```

Ответ:
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

---

### 1️⃣ Получить список языков

```bash
curl http://localhost:8080/languages
```

Ответ:
```json
[
  {
    "id": 1,
    "name": "Go",
    "slug": "go",
    "created_at": 1234567890000
  },
  {
    "id": 2,
    "name": "Python",
    "slug": "python",
    "created_at": 1234567891000
  }
]
```

**Объяснение:**
- `id` - уникальный номер языка
- `name` - красивое название (для отображения)
- `slug` - текстовый идентификатор (для URL)
- `created_at` - когда был добавлен язык

---

### 2️⃣ Создать сниппет

```bash
curl -X POST http://localhost:8080/languages/go/snippets \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Hello World Server",
    "filename": "server.go",
    "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello, World!\")\n}"
  }'
```

Ответ (успешное создание):
```json
{
  "id": 1,
  "title": "Hello World Server",
  "filename": "server.go",
  "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello, World!\")\n}",
  "created_at": 1234567890000,
  "updated_at": 1234567890000
}
```

**Объяснение:**
- `title` - название сниппета
- `filename` - имя файла (с расширением)
- `content` - сам код

---

### 3️⃣ Получить список сниппетов для языка

```bash
curl http://localhost:8080/languages/go/snippets
```

Ответ (краткий список, без полного кода):
```json
[
  {
    "id": 1,
    "title": "Hello World Server",
    "filename": "server.go",
    "created_at": 1234567890000
  },
  {
    "id": 2,
    "title": "HTTP Server",
    "filename": "http_server.go",
    "created_at": 1234567891000
  }
]
```

**Зачем краткий список?** Чтобы быстро показать пользователю что у него есть, без загрузки всего кода.

---

### 4️⃣ Получить полное содержимое сниппета

```bash
curl http://localhost:8080/languages/go/snippets/1/content
```

Ответ:
```json
{
  "id": 1,
  "title": "Hello World Server",
  "filename": "server.go",
  "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello, World!\")\n}",
  "created_at": 1234567890000,
  "updated_at": 1234567890000
}
```

---

### 5️⃣ Удалить сниппет

```bash
curl -X DELETE http://localhost:8080/languages/go/snippets/1
```

Ответ:
```json
{
  "message": "Snippet deleted successfully"
}
```

---

## 🗄 База данных (как хранятся данные)

### Таблица Languages (Языки)

```
id  | name       | slug       | created_at
----|------------|------------|------------------
1   | Go         | go         | 1234567890000
2   | Python     | python     | 1234567891000
3   | JavaScript | javascript | 1234567892000
```

### Таблица Snippets (Сниппеты)

```
id | language_id | title              | filename      | content        | created_at     | updated_at
---|-------------|-------------------|---------------|-----------------|---------------|------------------
1  | 1           | Hello World Server | server.go     | package main... | 1234567890000 | 1234567890000
2  | 1           | HTTP Server        | http.go       | package main... | 1234567891000 | 1234567891000
3  | 2           | Print Hello        | hello.py      | print("Hi")     | 1234567892000 | 1234567892000
```

**Связь:** Каждый сниппет указывает на свой язык через `language_id`.

---

## 🧪 Тестирование

### Запуск тестов

```bash
go test ./cmd/server -v
```

**Что тестируется:**
- ✅ Получение списка языков
- ✅ Создание сниппета
- ✅ Получение списка сниппетов
- ✅ Валидация неправильных данных
- ✅ Обработка ошибок

---

## ⚠️ Обработка ошибок

Если что-то пойдет не так, вы получите ответ с ошибкой:

```json
{
  "error": "language not found"
}
```

**Возможные ошибки:**
- `language not found` - Язык с таким slug не существует
- `snippet not found` - Сниппет не найден
- `Invalid request body` - Неправильные данные в запросе
- `Invalid snippet ID` - ID сниппета не число

---

## 🔄 Полный пример использования

### 1. Проверяем, какие языки доступны
```bash
curl http://localhost:8080/languages | jq
```

### 2. Создаем Go сниппет
```bash
curl -X POST http://localhost:8080/languages/go/snippets \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Create HTTP Server",
    "filename": "main.go",
    "content": "package main\n\nimport \"net/http\"\n\nfunc main() {\n  http.HandleFunc(\"/\", func(w http.ResponseWriter, r *http.Request) {\n    w.Write([]byte(\"Hello!\"))\n  })\n  http.ListenAndServe(\":8000\", nil)\n}"
  }'
```

### 3. Смотрим список сниппетов для Go
```bash
curl http://localhost:8080/languages/go/snippets | jq
```

### 4. Получаем полный код сниппета
```bash
curl http://localhost:8080/languages/go/snippets/1/content | jq
```

### 5. Удаляем сниппет
```bash
curl -X DELETE http://localhost:8080/languages/go/snippets/1
```

---

## 📝 Дополнительно

### Использование Postman

Вместо curl можно использовать Postman:

1. Скачайте Postman: https://www.postman.com/downloads/
2. Создайте новый запрос
3. Введите URL: `http://localhost:8080/languages`
4. Нажмите Send

---

## 🚧 Известные ограничения (пока что так :d )

- ❌ Нет авторизации (может добавить любой)
- ❌ Нет редактирования сниппетов (можно удалить и создать заново)
- ❌ Нет поиска по коду
- ❌ Нет загрузки файлов (только текст)

**Это запланировано для будущих версий!**

---


## 📞 Инструкции для Ubuntu

### Установка Go на Ubuntu

```bash
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

### Установка PostgreSQL на Ubuntu

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Проверка
psql --version
```

### Создание БД на Ubuntu

```bash
sudo -u postgres psql
CREATE DATABASE snippets;
\q
```

### Запуск проекта на Ubuntu

```bash
cd ~/mcode
go mod download
go run cmd/server/main.go
```

---

## 📚 Доп.ресурсы

- Go документация: https://golang.org/doc/
- Gin руководство: https://gin-gonic.com/docs/
- GORM документация: https://gorm.io/docs/
- PostgreSQL документация: https://www.postgresql.org/docs/

---

**Успехов в разработке! 🚀**
