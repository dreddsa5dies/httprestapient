[![Go Report Card](https://goreportcard.com/badge/github.com/dreddsa5dies/automateGo)](https://goreportcard.com/report/github.com/dreddsa5dies/httprestapient) ![License](https://img.shields.io/badge/License-GPL-blue.svg)  

## Тестовое задание
<details>
  <summary>Текст</summary>
Предлагается реализовать HTTP сервер с использованием ORM https://entgo.io/

Сервер должен реализовать CRUD работы с сущностью "Матрица атаки".

Сущность "Матрица атаки" имеет характеристики:
  - Вендор
  - Наименование матрицы
  - Версия матрицы
  - Дата создания
  - Дата обновления

Таблицу сделать плоскую. При желании, можно сделать несколько таблиц со связями, например вынести Вендора в отдельную таблицу, в которой вендор будет иметь характеристики "Страна".

Требования к CRUD:
  - Создать матрицу
  - Получить список матриц с пагинацией результатов
  - Получить одну матрицу по ключу
  - Изменить матрицу
  - Удалить матрицу

Использовать стандарт HTTP REST JSON API.

Результат оформить как репозиторий Github. Выслать ссылку на репозиторий.

В README к репозиторию описать как можно запустить проект, привести пример JSON для вставки пробных записей. Возможно, написать скрипт либо отдельную программу которая заполняет матрицу небольшим количеством пробных данных.
</details>

## Init
Init DB
```bash
  git clone 
  cd init
  docker build -t matrix .
  # connect to localhost:5555
  docker run -d --name matrix -p 5555:5432 matrix
  # проверка запуска образа
  docker ps
  # проверка базы
  docker exec -it matrix psql -d matrixdb -U muser
  ```

## Clear system
After check DB & app - delete docker image
```bash
docker stop matrix
docker rm matrix
docker rm $(docker ps -aq)
docker rmi matrix
```

## License
This project is licensed under GPL license. Please read the [LICENSE](https:/github.com/dreddsa5dies/httprestapient/tree/master/LICENSE.md) file.
