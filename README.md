# 📝 Quote API (Мини-сервис “Цитатник”)

REST API на Go для хранения и управления цитатами.

## 🚀 Запуск

1. Установи Go 1.18+ [https://golang.org/dl/](https://golang.org/dl/)
2. Клонируй репозиторий:

```bash
git clone https://github.com/mayten10000/Qoute_api.git
cd Qoute_api
# Qoute_api
go mod tidy
go run main.go


Метод	URL	Описание
POST	/quotes	Добавить цитату
GET	/quotes	Получить все цитаты
GET	/quotes/random	Получить случайную цитату
GET	/quotes?author=Автор	Получить цитаты по автору
DELETE	/quotes/{id}	Удалить цитату по ID
