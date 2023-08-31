# Avito-backend-task-2023

## Подготовка

Перед запуском приложения можно настроить пароль к базе данных, а также порты, на которых будут работать сервис и база данных, название базы данных и так далее (хотя изначально уже все настроено, может появиться необходимость изменить конфиг). 

Пароль можно поменять в файле .env.

Все, что связано с БД можно поменять в файле ```configs/config.yml```, а также в ```docker-compose.yml``` в соответствующих местах.

Кроме того, необходимо запустить Docker

## Сборка

Чтобы сбилдить проект, нужно ввести в терминале команду ```docker-compose up --build avito-backend-task-2023```. После некоторого количества времени запустятся контейнеры с postgres и самим сервисом. Если все хорошо, значит сервис запустился и можно делать запросы.

### Возможные ошибки при сборке

Если вдруг возникла ошибка, связанная с файлом ```wait-for-postgres.sh``` (такого быть не должно, но может появиться), то необходимо немного изменить файлы ```Dockerfile``` и ```docker-compose.yml```. В ```Dockerfile``` нужно удалить строку 11, а в ```docker-compose.yml``` убрать ```"./wait-for-postgres.sh db"``` из строки 6.

## Задание

Выполнены все обязательные задания, добавлен Swagger для API, а также выполнено дополнительное задание 1.

### Запросы

Все end-points находятся в группе ```/users-segmentation```

#### End-point ```/user``` необходим для работы с пользователями (создание/удаление)

***POST /user - создание пользователя***

Пример тела запроса:
```
{
  "name": "David",
  "username": "d4viid",
  "password": "h7dsg8d"
}
```

Пример ответа:
```
{
  "id": "2574c49b-4f15-49d0-8e27-b5741ecb43af" // uuid нового пользователя
}
```

***DELETE /user/:id - удаление пользователя по его uuid***

Пример запроса: ```localhost:8000/users-segmentation/user/2574c49b-4f15-49d0-8e27-b5741ecb43af```

Пример ответа:
```
{
  "status": "ok" // успешное удаление пользователя
}
```


#### End-point ```/segment``` необходим для работы с сегментами (создание/удаление)

***POST /segment - создание пользователя***

Пример тела запроса:
```
{
  "name": "AVITO_TEST_SEGMENT_1"
}
```

Пример ответа:
```
{
  "id": "2574c49b-4f15-49d0-8e27-b5741ecb43af" // uuid нового сегмента
}
```

***DELETE /segment - удаление сегмента***

Пример тела запроса:
```
{
  "name": "AVITO_TEST_SEGMENT_1"
}
```

Пример ответа:
```
{
  "status": "ok" // успешное удаление сегмента
}
```

