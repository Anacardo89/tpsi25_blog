

/*  Create DB */
DROP DATABASE tpsi25_blog;
CREATE DATABASE tpsi25_blog;
USE tpsi25_blog;

CREATE TABLE users (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	username VARCHAR(32) NOT NULL DEFAULT '',
	email VARCHAR(128) NOT NULL DEFAULT '',
	hashpass VARCHAR(128) NOT NULL DEFAULT '',
	profile_pic VARCHAR(64) NOT NULL DEFAULT '',
	profile_pic_ext VARCHAR(10) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	active TINYINT NOT NULL DEFAULT 0,
	PRIMARY KEY(id),
	UNIQUE KEY username (username),
	UNIQUE KEY email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE follows (
	follower_id INT NOT NULL REFERENCES users(id),
	followed_id INT NOT NULL REFERENCES users(id),
	UNIQUE KEY follow_relation (follower_id, followed_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE sessions (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	session_id VARCHAR(256) NOT NULL DEFAULT '',
	user_id INT NOT NULL REFERENCES users(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	active TINYINT NOT NULL DEFAULT 0,
	PRIMARY KEY (id),
	UNIQUE KEY session_id (session_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE posts (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	post_guid VARCHAR(256) NOT NULL DEFAULT '',
	author_id INT NOT NULL REFERENCES users(user_id),
	title VARCHAR(256) DEFAULT NULL,
	content MEDIUMTEXT,
	post_image VARCHAR(64) NOT NULL DEFAULT '',
	image_ext VARCHAR(10) NOT NULL DEFAULT '',
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	is_public TINYINT NOT NULL DEFAULT 0,
	rating INT NOT NULL DEFAULT 0,
	active TINYINT NOT NULL DEFAULT 0,
	PRIMARY KEY (id),
	UNIQUE KEY post_guid (post_guid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE post_ratings (
	post_id INT NOT NULL REFERENCES posts(id),
	user_id INT NOT NULL REFERENCES users(id),
	rating_value INT NOT NULL DEFAULT 0,
	UNIQUE KEY post_rating (post_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE comments (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	post_guid VARCHAR(256) NOT NULL REFERENCES posts(post_guid),
	author_id INT NOT NULL REFERENCES users(user_id),
	content MEDIUMTEXT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	rating INT NOT NULL DEFAULT 0,
	active TINYINT NOT NULL DEFAULT 0,
	PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE comment_ratings (
	comment_id INT NOT NULL REFERENCES comments(id),
	user_id INT NOT NULL REFERENCES users(id),
	rating_value INT NOT NULL DEFAULT 0,
	UNIQUE KEY comment_rating (comment_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DELIMITER $$

CREATE TRIGGER trg_after_insert_comment_ratings
AFTER INSERT ON comment_ratings
FOR EACH ROW
BEGIN
    UPDATE comments
    SET rating = rating + NEW.rating_value
    WHERE id = NEW.comment_id;
END$$

DELIMITER ;


DELIMITER $$

CREATE TRIGGER trg_after_update_comment_ratings
AFTER UPDATE ON comment_ratings
FOR EACH ROW
BEGIN
    UPDATE comments
    SET rating = rating - OLD.rating_value + NEW.rating_value
    WHERE id = NEW.comment_id;
END$$

DELIMITER ;


DELIMITER $$

CREATE TRIGGER trg_after_insert_post_ratings
AFTER INSERT ON post_ratings
FOR EACH ROW
BEGIN
    UPDATE posts
    SET rating = rating + NEW.rating_value
    WHERE id = NEW.post_id;
END$$

DELIMITER ;


DELIMITER $$

CREATE TRIGGER trg_after_update_post_ratings
AFTER UPDATE ON post_ratings
FOR EACH ROW
BEGIN
    UPDATE posts
    SET rating = rating - OLD.rating_value + NEW.rating_value
    WHERE id = NEW.post_id;
END$$

DELIMITER ;