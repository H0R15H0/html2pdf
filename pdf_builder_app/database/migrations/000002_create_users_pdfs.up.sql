CREATE TABLE IF NOT EXISTS users_pdfs(
   id UUID PRIMARY KEY,
   user_id UUID NOT NULL REFERENCES users(id),
   s3_url TEXT NOT NULL,
   unified_at TIMESTAMP
);
