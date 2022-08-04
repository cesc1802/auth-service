-- +migrate Up
CREATE TABLE users
(
    id               int          not null auto_increment,
    full_name        varchar(100),
    last_name        varchar(50)  not null,
    first_name       varchar(50)  not null,
    login_id         varchar(50)  not null,
    password         varchar(100) not null,
    salt             varchar(100) not null,
    refresh_token_id varchar(100) not null,
    status           int          not null default 1,
    created_at       timestamp null default current_timestamp,
    updated_at       timestamp null default current_timestamp on update current_timestamp,
    deleted_at       timestamp null,
    primary key (id)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
-- +migrate Down
drop table users;