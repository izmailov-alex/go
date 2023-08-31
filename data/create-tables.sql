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
    user_id INT NOT NULL,
    PRIMARY KEY (`user_id`),
    UNIQUE (`user_id`)
);

CREATE TABLE user_segments (
    user_id INT NOT NULL,
    segment_id INT NOT NULL,
    PRIMARY KEY (`user_id`,`segment_id`)
);