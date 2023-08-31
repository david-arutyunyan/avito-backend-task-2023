CREATE TABLE users
(
    id            varchar(255) PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE segments
(

    id   varchar(255) PRIMARY KEY,
    name varchar(255) not null unique
);

CREATE TABLE users_segments
(
    id         varchar(255) PRIMARY KEY,
    user_id    varchar(255) references users (id) on delete cascade    not null,
    segment_id varchar(255) references segments (id) on delete cascade not null,
    expiration_time timestamp, -- Now + date
    CONSTRAINT unique_user_segment_pair UNIQUE (user_id, segment_id)
);

