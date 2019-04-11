-- migrate:up
create table news (
  id integer,
  title varchar(255),
  content varchar(255),
  created_at time
);

-- migrate:down
drop table news;