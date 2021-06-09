drop table if exists wines;
drop table if exists feedbacks;
drop table if exists querylogs;

create table wines (
  id           serial primary key,
  priority     int,
  key_id       varchar(255) not null unique,
  store        varchar(255),
  wine_name    varchar(255),
  locations    varchar(255)[],
  price        int,
  price_type   varchar(64),
  wine_type    varchar(64),
  country      varchar(255),
  grapes       varchar(255)[],
  acidity      int,
  sweetness    int,
  sparkling    int,
  food_matches varchar(255)[],
  image_url    text,
  created_at   timestamp not null  
);

create table feedbacks (
  id         serial primary key,
  body       text not null,
  created_at timestamp not null       
);

create table querylogs (
  id serial primary key,
  store        varchar(255),
  price        varchar(64),
  wine_type    varchar(64),
  food_match   varchar(255),
  created_at   timestamp not null
);