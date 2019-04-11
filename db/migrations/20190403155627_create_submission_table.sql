-- migrate:up

create table submission (
  id integer,
  uid varchar(255),
  username varchar(255),
  pid varchar(255),
  judge varchar(255),
  code varchar(255),
  time integer(11),
  memory int(11),
  token varchar(255),
  created_at time
);

-- migrate:down
drop table submission;

