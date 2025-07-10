Прописать команды 
cd tasks
cd cmd
cd app
go run .\main.go

3 варианта запроса в postman
POST http://localhost:8080/tasks с пустным скобками в json raw формате {}
GET  http://localhost:8080/tasks/уникальныйuuid который получите после post запроса
DELETE http://localhost:8080/tasks/уникальныйuuid который получите после post запроса
