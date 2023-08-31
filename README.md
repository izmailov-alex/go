# Сервис динамического сегментирования пользователей
***
### Описание
Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

***
### Как запустить:
1. Создайте и заполните .env файл
2. docker-compose up --build
***
### Реализованные запросы
1. GET /segments - возвращает список сегментов
Пример запроса: ```http://0.0.0.0:8081/segments```
Пример ответа:
```
Status: 200 OK
[
  {
    "segment_id": "3",
    "segment_name": "AVITO_DISCOUNT_30"
  },
  {
    "segment_id": "4",
    "segment_name": "AVITO_PERFORMANCE_VAS"
  },
  {
    "segment_id": "6",
    "segment_name": "AVITO_VOICE_MESSAGES"
  }
]
```

2. GET /segments/:segment_id - возвращает сегмент с нужным id
Пример запроса: ```http://0.0.0.0:8081/segments/7```
Пример ответа:
```
Status: 200 OK
{
  "segment_id": "7",
  "segment_name": "AVITO_DISCOUNT_30"
}
```

3. GET /user_segments/:user_id - возвращает активные сегменты пользователя с id
Пример запроса: ```http://0.0.0.0:8081/user_segments/7```
Пример ответа:
```
Status: 200 OK
[
  {
    "segment_id": "2",
    "segment_name": "AVITO_PERFORMANCE_VAS"
  },
  {
    "segment_id": "6",
    "segment_name": "AVITO_VOICE_MESSAGES"
  }
]
```

4. POST /segments/:segment_name - создает сегмент с заданным именем
Пример запроса: ```http://0.0.0.0:8081/segments/AVITO_VOICE_MESSAGES```
Пример ответа:
```
Status: 201 Created
{
  "segment_id": "9",
  "segment_name": "AVITO_VOICE_MESSAGES"
}
```

5. POST /users/:user_id - создает пользователя с заданным id
Пример запроса: ```http://0.0.0.0:8081/users/1009```
Пример ответа:
```
Status: 201 Created
{
  "user_id": "1009"
}
```

6. PUT /user_segments/:user_id - обновляет активные сегменты пользователя с заданным id
Пример запроса:
```
curl http://0.0.0.0:8081/user_segments/1000 \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"segment_add": ["AVITO_VOICE_MESSAGES", "AVITO_DISCOUNT_30"], "segment_delete":["AVITO_PERFORMANCE_VAS","AVITO_DISCOUNT_50"]}'
```
Пример ответа:
```
Status: 200 OK
{
  "segment_add": [
    "AVITO_VOICE_MESSAGES",
    "AVITO_DISCOUNT_30"
  ],
  "segment_delete": [
    "AVITO_PERFORMANCE_VAS",
    "AVITO_DISCOUNT_50"
  ]
}
```
7. DELETE /segments/:segment_name - удаляет заданный сегмент из всех таблиц
Пример запроса:  ``` http://0.0.0.0:8081/segments/AVITO_DISCOUNT_30 ```
Пример ответа:
```
Status: 200 OK
{
  "segment_id": "10",
  "segment_name": "AVITO_DISCOUNT_30"
}
```