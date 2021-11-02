-- +goose Up
CREATE TABLE review (
  id BIGSERIAL PRIMARY KEY,
  foo BIGINT NOT NULL
);

-- +goose Down
DROP TABLE review;
