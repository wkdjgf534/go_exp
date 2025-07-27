CREATE TABLE IF NOT EXISTS user_invitations (
  token bytea PRIMARY KEY,
  user_id bigint NOT NULL

  -- if  you want a composite key for one to one relation on DB layer
  -- PRIMARY KEY (token, user_id)
)
