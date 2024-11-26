-- name: FindCommentsForArticle :many
select
  c.id,
  c.lookup_id,
  c.created_at,
  c.updated_at,
  c.article_slug,
  c.published_at,
  c.content,
  c.user_id,
  u.email,
  u.name,
  u.should_notify as user_should_notify
from
  comments c
  join users u on u.id = c.user_id
where
  c.article_slug = sqlc.arg ('slug')
  and c.approved_at is not null
order by
  c.published_at;

-- name: CreateComment :one
insert into comments (article_slug, user_id, published_at, approval_code, should_notify, content)
  values (sqlc.arg ('article_slug'), sqlc.arg ('user_id'), sqlc.arg ('published_at'), sqlc.arg ('approval_code'),
    sqlc.arg ('should_notify'), sqlc.arg ('content'))
returning
  *;

-- name: ApproveComment :one
update
  comments
set
  approved_at = current_timestamp
where
  approval_code = sqlc.arg ('approval_code')
returning
  *;
