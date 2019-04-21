
-- ARTICLE
alter table articles add user_id uuid references users (id);