CREATE TABLE users
(
    id           varchar(45) NOT NULL,
    first_name   varchar(45)  DEFAULT NULL,
    last_name    varchar(45)  DEFAULT NULL,
    email        varchar(45)  DEFAULT NULL,
    password     varchar(100) DEFAULT NULL,
    time_created varchar(45)  DEFAULT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE posts
(
    id           varchar(45) NOT NULL,
    title        varchar(50) DEFAULT NULL,
    boby         varchar(225) DEFAULT NULL,
    time_created varchar(45) DEFAULT NULL,
    user_id      varchar(45) DEFAULT NULL,
    access       int         DEFAULT NULL,
    PRIMARY KEY (id),
    CONSTRAINT id FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE comments
(
    id          varchar(45) NOT NULL,
    commentt    varchar(200) DEFAULT NULL,
    user_id     varchar(45)  DEFAULT NULL,
    post_id     varchar(45)  DEFAULT NULL,
    time_posted varchar(45)  DEFAULT NULL,
    PRIMARY KEY (id)
);