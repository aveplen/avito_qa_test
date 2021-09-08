# Тестовое задание для QA-cтажёра

Необходимо написать программу для отправки сообщений в Slack. Программа должна принимать на вход json-файл (формат представлен в example.json) и отправлять текстовые сообщения в каналы из этого json-файла.

## Инструкция по запуску

Если на ПК не стоит Go, но есть Docker (что было бы странно), то можно запустить так:

```
(точка в конце первой команды - пусть к Dockerfile)

docker build -t aveplen/avito_qa_test .
docker run -p 443:443 aveplen/avito_qa_test
```

Если Go всё таки стоит:
```
go mod download
go run cmd/main.go [-f message.json]


Или так:
go mod download
go build -o slack_spammer.exe cmd\main.go   (Windows)
.\slack_spammer.exe -f message.json

go mod download
go build -o slack_spammer cmd/main.go       (Linux)
./slack_spammer -f message.json
```