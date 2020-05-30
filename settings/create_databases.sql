create database twitter_clone;

create user editer@localhost identified by 'Tr_kp_514021';

grant create on twitter_clone.* to editer@localhost;

use twitter_clone;

create table account (
  id int not null auto_increment primary key,
  account_id varchar(255) unique,
  account_name varchar(255),
  password varchar(255),
  account_image varchar(255),
  created_at datetime,
  deleated_at datetime
);

create table followers (
  following_id int,
  followed_id int
);

create table fabolits (
  id int not null auto_increment primary key,
  user_id int,
  tweet_id int
);

create table tweets (
  id int not null auto_increment primary key,
  user_id int,
  body varchar(255),
  created_at datetime,
  deleated_at datetime
);

create table comments (
  id int not null auto_increment primary key,
  user_id int,
  tweet_id int,
  body varchar(255),
  created_at datetime,
  deleated_at datetime
);
