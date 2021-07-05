CREATE TABLE users
(
    id serial               not null unique,
    firstName  varchar(255) not null,
    secondName varchar(255) not null,
    username   varchar(255) not null unique,
    password   varchar(255) not null
);

CREATE TABLE adverts
(
    id serial                not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    price       int          not null
);

CREATE TABLE adverts_list
(
    id serial                                              not null unique,
    advert_id int references adverts(id) on delete cascade not null,
    user_id   int references users(id)   on delete cascade not null
);