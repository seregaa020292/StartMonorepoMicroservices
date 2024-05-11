### TODO: CI/CD
Например, в рабочих процессах GitHub есть опция под названием «путь». При использовании фильтров «ветви» и «пути» рабочий процесс будет выполняться только в том случае, если оба фильтра удовлетворены. Это означает, что если папка будет изменена, рабочий процесс запустится.
``` yml
on:
  pull_request:
    branches:
      - main
    paths:
      - './controlling'
```
Затем вы можете настроить рабочие процессы для каждой папки проекта.
