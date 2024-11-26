-- +goose Up
-- +goose StatementBegin
create table users (
  id integer primary key autoincrement,
  created_at text default current_timestamp not null,
  updated_at text default current_timestamp not null,
  firebase_uid text not null,
  email text not null,
  name text not null,
  should_notify int not null default 0,
  is_banned int not null default 0,
  is_admin int not null default 0,
  constraint email_unique unique (email),
  constraint firebase_uid_unique unique (firebase_uid)
);

create table comments (
  id integer primary key autoincrement,
  lookup_id text default (lower(hex (randomblob (16)))) not null,
  user_id integer not null,
  created_at text default current_timestamp not null,
  updated_at text default current_timestamp not null,
  article_slug text not null,
  published_at text not null,
  approved_at text null,
  approval_code text not null,
  should_notify int not null default 0,
  content text not null,
  foreign key (user_id) references users (id),
  constraint approval_code_unique unique (approval_code),
  constraint lookup_id_unique unique (lookup_id)
);

create index comments_article_slug_published_at_idx on comments (article_slug, published_at);

create index users_firebase_uid_idx on users (firebase_uid);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table comments;

drop table users;

-- +goose StatementEnd
