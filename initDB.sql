DROP TABLE users;
DROP TABLE diary;

CREATE TABLE users (
  id         VARCHAR(26) NOT NULL UNIQUE,
  name       VARCHAR(255),
  email      VARCHAR(255) NOT NULL UNIQUE,
  password   VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL   
);

CREATE TABLE diary (
  id           VARCHAR(26) NOT NULL UNIQUE,
  fishingDate  DATE,
  place        VARCHAR(100),       
  caughtFish   VARCHAR(80),       
  diaryComment VARCHAR(500),       
  rod          VARCHAR(100),
  method       VARCHAR(100),   
  lure         VARCHAR(100),       
  weather      VARCHAR(30),       
  wind         TINYINT UNSIGNED,       
  tide         VARCHAR(2)
);