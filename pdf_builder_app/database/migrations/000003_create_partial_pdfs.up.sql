CREATE TABLE IF NOT EXISTS partial_pdfs(
   id UUID PRIMARY KEY,
   unified_pdf_id UUID REFERENCES users_pdfs(id),
   source_html_url TEXT NOT NULL,
   number INT NOT NULL,
   s3_url TEXT NOT NULL,
   pdf_created_at TIMESTAMP
);
