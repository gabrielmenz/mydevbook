CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nameuser VARCHAR(100) NOT NULL,
    nick VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    pw VARCHAR(255) NOT NULL,  -- Longer length for password hash storage
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE = InnoDB;

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
) ENGINE=INNODB;
CREATE TABLE posts(
    id int auto_increment primary key, 
    title varchar(100) not null,
    content varchar(300) not null,

    author_id int not null,
    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdAt timestamp default current_timestamp
) ENGINE=INNODB;