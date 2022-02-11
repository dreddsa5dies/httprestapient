[![Go Report Card](https://goreportcard.com/badge/github.com/dreddsa5dies/httprestapient)](https://goreportcard.com/report/github.com/dreddsa5dies/httprestapient) ![License](https://img.shields.io/badge/License-GPL-blue.svg)  

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

Таблицу сделать плоскую. При желании, можно сделать несколько таблиц со
связями, например вынести Вендора в отдельную таблицу, в которой вендор
будет иметь характеристики "Страна".


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
<details>
  <summary>Init DB</summary>

  ```bash
    git clone https://github.com/dreddsa5dies/httprestapient
    cd ./httprestapient/init
    docker build -t matrix .
  ```

</details>

## Check API
<details>
  <summary>Start database docker image & application</summary>

```bash
    docker run -d --name matrix -p 5432:5432 matrix
    # проверка запуска образа
    docker ps
    cd ../cmd
    go run main.go
  ```

</details>

<details>
  <summary>Check application</summary>
  You need start new terminal and run commands in turn:

```bash
    # Получить список матриц с пагинацией результатов
    curl http://127.0.0.1:8080/api/
    # Получить одну матрицу по ключу
    curl http://127.0.0.1:8080/api/1
    # Создать матрицу
    curl -X 'POST' -H 'Content-Type: application/json' -d '{"id":3, "IdMatrix":"AAAAAAA","VendorName":"BBBBBBB","NameMatrix":"","VersionMatrix":"0.0.1","CreateDate":"2021-11-02T00:00:00Z","ModifyDate":"2022-01-15T00:00:00Z"}' http://127.0.0.1:8080/api/
    # Проверка создания
    curl http://127.0.0.1:8080/api/3
    # Изменить матрицу
    curl -X 'PUT' -H 'Content-Type: application/json' -d '{"id":3, "IdMatrix":"11111111","VendorName":"22222","NameMatrix":"33333","VersionMatrix":"0.0.2","CreateDate":"2021-11-02T00:00:00Z","ModifyDate":"2022-02-11T00:00:00Z"}' http://127.0.0.1:8080/api/3
    # Проверка изменения
    curl http://127.0.0.1:8080/api/3
    # Удалить матрицу
    curl -X 'DELETE' http://127.0.0.1:8080/api/3
    # Проверка удаления
    curl http://127.0.0.1:8080/api/
  ```

  Ctl+C для остановки main.go

</details>

## Clear system after
<details>
  <summary>Delete docker image</summary>

  ```bash
    docker stop matrix
    docker rm matrix
    docker rmi matrix
  ```

</details>

## License
This project is licensed under GPL license. Please read the [LICENSE](https:/github.com/dreddsa5dies/httprestapient/tree/master/LICENSE.md) file.
