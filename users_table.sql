-- usersテーブル
CREATE TABLE IF NOT EXISTS golang.users(
    id int auto_increment,
    name VARCHAR (64) UNIQUE NOT NULL,
    PRIMARY KEY (id)
);