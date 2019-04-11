-- migrate:up
create table content (
  id integer,
  title varchar(255),
  pidList varchar(255),
  problems varchar(255),
  start_time time,
  end_time time,
  type varchar(255)
);

-- migrate:down
drop table content;
