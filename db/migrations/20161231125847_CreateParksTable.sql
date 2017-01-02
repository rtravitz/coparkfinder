
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE parks (
  id serial,
  name varchar(255),
  street varchar(255),
  city varchar(255),
  zip varchar(255),
  email varchar(255),
  description text,
  url text,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE parks;

