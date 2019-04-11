-- migrate:up
create table problem (
  id integer,
  title varchar(255),
  description varchar(255),
  input varchar(255),
  output varchar(255),
  sample_input varchar(255),
  sample_output varchar(255),
  solve integer(11),
  submission integer(11)
);

-- migrate:down
drop table problem;