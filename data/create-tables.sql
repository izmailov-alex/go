DROP TABLE IF EXISTS segments;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_segments;
CREATE TABLE segments (
  segment_id         INT AUTO_INCREMENT NOT NULL,
  segment_name      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`segment_id`),
  UNIQUE (`segment_name`)
);

CREATE TABLE users (
    user_id INT AUTO_INCREMENT NOT NULL,
    PRIMARY KEY (`user_id`)
);

CREATE TABLE user_segments (
    user_id INT NOT NULL,
    segment_id INT NOT NULL
)
INSERT INTO users

VALUES
    (),
    (),
    (),
    ();

INSERT INTO segments
  (segment_name)
VALUES
  ('AVITO_VOICE_MESSAGES'),
  ('AVITO_PERFORMANCE_VAS'),
  ('AVITO_DISCOUNT_30'),
  ('AVITO_DISCOUNT_50');