-- +migrate Up
create table user_roles
(
    role_id    int not null,
    user_id    int not null,
    status     int not null default 1,
    created_at timestamp null default current_timestamp,
    updated_at timestamp null default current_timestamp on update current_timestamp,
    deleted_at timestamp null,
    primary key (role_id, user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
-- +migrate Down
drop table user_roles;
