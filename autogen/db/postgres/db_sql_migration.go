// Code generated. DO NOT EDIT!

package postgres

// DatabaseSQLMigration is generated form a fileset
var DatabaseSQLMigration = map[string]string{
	"db_migration_1": `create table schema_version (
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
`,
}

// DatabaseSQLMigrationChecksums is generated from a fileset and contains files checksums
var DatabaseSQLMigrationChecksums = map[string]string{
	"db_migration_1": "23084cf1364e035673bbe966aaf0efd64c7a6ed1f5a25cd1f8413d32b953f6de",
}
