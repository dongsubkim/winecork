drop table admins cascade;
drop table sessions cascade;
drop table wines cascade;
drop table feedbacks cascade;

create table admins (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  admin_id   integer references admins(id),
  created_at timestamp not null   
);

create table wines (
  id          serial primary key,
  key         varchar(255) not null unique,
  store       varchar(255),
  wine_name   varchar(255),
  locations   varchar(255)[],
  price       int,
  wine_sort   varchar(64),
  wine_origin varchar(255),
  graphs      varchar(255)[],
  acidity     int,
  sweetness   int,
  sparkling   int
  foodmatchs  varchar(255)[],
  created_at  timestamp not null  
);

create table feedbacks (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text not null,
  created_at timestamp not null       
);
