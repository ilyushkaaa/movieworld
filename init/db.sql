CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL
);

CREATE TABLE IF NOT EXISTS films
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    duration INTEGER NOT NULL,
    min_age INTEGER NOT NULL,
    country varchar(255) NOT NULL,
    producer_name varchar(255) NOT NULL,
    date_of_release DATE NOT NULL,
    sum_mark      INTEGER NOT NULL,
    num_of_marks    INTEGER NOT NULL,
    rating        DECIMAL(3, 1) NOT NULL,
);

CREATE TABLE IF NOT EXISTS actors
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    nationality VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL
    );

CREATE TABLE IF NOT EXISTS actor_films
(
    id   SERIAL PRIMARY KEY,
    film_id INTEGER NOT NULL REFERENCES films (id),
    actor_id INTEGER NOT NULL REFERENCES actors (id)
);


CREATE TABLE IF NOT EXISTS genres
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS film_genres
(
    id   SERIAL PRIMARY KEY,
    film_id INTEGER NOT NULL REFERENCES films (id),
    genre_id INTEGER NOT NULL REFERENCES genres (id)
);


CREATE TABLE IF NOT EXISTS reviews
(
    id   SERIAL PRIMARY KEY,
    film_id INTEGER NOT NULL REFERENCES films (id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    mark INTEGER NOT NULL,
    comment TEXT
);


CREATE TABLE IF NOT EXISTS favourite_films
(
    id   SERIAL PRIMARY KEY,
    film_id INTEGER NOT NULL REFERENCES films (id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    );
