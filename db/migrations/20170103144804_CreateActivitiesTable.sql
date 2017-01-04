
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE activities (
  id serial,
  type varchar(255),
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE activities;

