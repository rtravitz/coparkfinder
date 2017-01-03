
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE Table parks_activities (
  id serial,
  park_id integer,
  activity_id integer,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE parks_activities;

