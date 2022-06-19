create table schema_version (
  version text not null
);

create table articles (
  id bigserial not null,
  user_id int not null,
  category_id int null,
  title text not null,
  text text,
  html text,
  url text,
  image text,
  hash text not null,
  status article_status default 'unread',
  published_at timestamp with time zone null,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone null,
    
  primary key (id),
  unique (user_id, hash),
  foreign key (user_id) references users(id) on delete cascade,
  foreign key (category_id) references categories(id) on delete set null
);
