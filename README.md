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

Использовать стандар`зможно, написать скрипт либо отдельную программу которая заполняет матрицу небольшим количеством пробных данных.
</details>

## Init
<details>
  <summary>Init DB</summary>

  ```bash
    git clone 
    cd init
    docker build -t matrix .
    # connect to localhost:5432
    docker run -d --name matrix -p 5432:5432 matrix
    # проверка запуска образа
    docker ps
    # проверка базы
    docker exec -it matrix psql -d matrix_attack -U muser
  ```

</details>

## Check API

## Clear system
<details>
  <summary>After check DB & app - delete docker image</summary>

  ```bash
    docker stop matrix
    docker rm matrix
    docker rm $(docker ps -aq)
    docker rmi matrix
  ```

</details>

## License
This project is licensed under GPL license. Please read the [LICENSE](https:/github.com/dreddsa5dies/httprestapient/tree/master/LICENSE.md) file.
