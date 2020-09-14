-- CREATE DATABASE viewer_app;
--
-- \c viewer_app

CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY ,
  name VARCHAR( 25 ) NOT NULL ,
  email VARCHAR( 35 ) NOT NULL ,
  password VARCHAR( 60 ) NOT NULL ,
  status int ,
  UNIQUE (email)
);

CREATE TABLE questions (
  id SERIAL NOT NULL PRIMARY KEY ,
  user_id int NOT NULL references users(id) ,
  sentence_id int NOT NULL references sentences(id) ,
  title VARCHAR( 100 ) NOT NULL ,
  content text,
  page_num int,
  row_num int,
  created_at timestamp with time zone NOT NULL
);

CREATE TABLE sentences (
  id SERIAL NOT NULL PRIMARY KEY ,
  page_id int NOT NULL references pages(id) ,
  content text NOT NULL
);

CREATE TABLE pages (
  id SERIAL NOT NULL PRIMARY KEY ,
  chapter_id int NOT NULL references chapters(id) ,
  number int NOT NULL
);

CREATE TABLE chapters (
  id SERIAL NOT NULL PRIMARY KEY ,
  book_id int NOT NULL references books(id) ,
  number int NOT NULL
);

CREATE TABLE books (
  id SERIAL NOT NULL PRIMARY KEY ,
  genre_id int NOT NULL references genres(id) ,
  title VARCHAR( 100 ) NOT NULL ,
  image VARCHAR( 100 ),
  price int NOT NULL , 
  author VARCHAR( 100 ),
  bio text
);

CREATE TABLE user_books (
  id SERIAL NOT NULL PRIMARY KEY ,
  user_id int NOT NULL references users(id) ,
  status int DEFAULT 0 NOT NULL ,
  created_at timestamp with time zone NOT NULL
);

CREATE TABLE genres (
  id SERIAL NOT NULL PRIMARY KEY ,
  name VARCHAR( 100 ) NOT NULL
);

CREATE TABLE replys (
  id SERIAL NOT NULL PRIMARY KEY ,
  user_id int NOT NULL references users(id) ,
  question_id int NOT NULL references questions(id) ,
  content text  NOT NULL,
  created_at timestamp with time zone NOT NULL
);

CREATE TABLE likes (
  id SERIAL NOT NULL PRIMARY KEY ,
  user_id int NOT NULL references users(id) ,
  reply_id int NOT NULL references replys(id) ,
);
