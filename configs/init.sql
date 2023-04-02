CREATE TABLE urls(
    id SERIAL NOT NULL UNIQUE,
    origin VARCHAR(255) NOT NULL UNIQUE,
    short VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY(origin)
);

INSERT INTO urls (origin, short) VALUES ('https://google.com', 'Iv1a');