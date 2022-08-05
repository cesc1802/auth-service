
-- +migrate Up
alter table users
    modify deleted_at bigint null;
alter table permissions
    modify deleted_at bigint null;
alter table role_permissions
    modify deleted_at bigint null;
alter table user_roles
    modify deleted_at bigint null;
alter table roles
    modify deleted_at bigint null;
-- +migrate Down
