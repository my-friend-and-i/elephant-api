-- migrate:up
create table solution (
  id integer,
  sid integer(11),
  uid integer(11),
  judge integer(11),
  code varchar(255),
  memory integer(11),
  language integer(11),
  status integer(11)
);

-- migrate:down
drop table solution;

