create table schema_version (
  version text not null
);

create table users (
  id bigserial not null,
  username varchar(40) not null unique,
  email varchar(255) not null unique,
  password varchar not null,
  "createdAt" timestamp with time zone not null default now(),
  "updatedAt" timestamp with time zone null,

  primary key (id)
);

create table categories (
  id bigserial not null,
  "userId" int not null,
  title varchar(255) not null,
  "createdAt" timestamp with time zone not null default now(),
  "updatedAt" timestamp with time zone null,

  primary key (id),
  unique ("userId", title),
  foreign key ("userId") references users(id) on delete cascade
);

create type article_status as enum('unread', 'read');

create table articles (
  id bigserial not null,
  "userId" int not null,
  "categoryId" int null,
  title text not null,
  text text,
  html text,
  url text,
  image text,
  hash text not null,
  status article_status default 'unread',
  "createdAt" timestamp with time zone not null default now(),
  "updatedAt" timestamp with time zone null,
    
  primary key (id),
  unique ("userId", hash),
  foreign key ("userId") references users(id) on delete cascade,
  foreign key ("categoryId") references categories(id) on delete set null
);

