Quick Start
1. Клонируйте репозиторий и перейдите в папку проекта
bash
Копировать
Редактировать
git clone https://github.com/your_username/people-info-api.git
cd people-info-api
2. Поднимите инфраструктуру
bash
Копировать
Редактировать
make up
3. Примените миграции
bash
Копировать
Редактировать
make migrate
4. Наполните базу тестовыми данными
bash
Копировать
Редактировать
make seed
5. Откройте Swagger
http://localhost:8081/swagger/index.html

Файл .env уже в репозитории — ничего настраивать не нужно.
Всё запускается за ~10 секунд.