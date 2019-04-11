-- migrate:up
create table user (
  id integer,
  username varchar(255),
  nickname varchar(255),
  email varchar(255),
  password varchar(255),
  submission integer(11),
  solve integer(11),
  token varchar(255),
  created_at time,
  last_login time
);
-- migrate:down
drop table user;

