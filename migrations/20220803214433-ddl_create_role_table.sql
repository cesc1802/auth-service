-- +migrate Up
CREATE TABLE roles
(
    id          int          not null auto_increment,
    name        varchar(100) not null,
    description varchar(255),
    status      int          not null default 1,
    created_at  timestamp null default current_timestamp,
    updated_at  timestamp null default current_timestamp on update current_timestamp,
    deleted_at  timestamp null,
    primary key (id)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
-- +migrate Down
drop table roles;