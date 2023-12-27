# Сервис подбора фильмов MOVIEWORLD


api

1. GET /actors/ - список всех актеров +
2. GET /films - список всех фильмов, может принимать query параметры 
genre country  producer
3. GET /films/{ACTOR_ID} список фильмов в которых снимался актер с таким айди
4. GET /film/{FILM_ID} информация о конкретном фильме 
5. GET /actor/{ACTOR_ID} информация о конкретном актере
6. GET /films/soon/ список предстоящих релизов
7. POST /register - регистрация
8. POST /login - вход по логину и паролю
9. POST /review/{FILM_ID} - оставить отзыв
10. DELETE /review/{FILM_ID}/{REVIEW_ID} - удалить отзыв 
11. PUT /review/{FILM_ID}/{REVIEW_ID} - изменить отзыв
12. GET /films/favourite/{USER_ID} - избранные фильмы пользователя
13. POST /films/favourite/{FILM_ID} - добавить фильм в избранное
14. POST /logout - выйти из аккаунта
15. GET /review/{FILM_ID} - получение отзывов о фильме