INSERT INTO users (id, created_at)
  VALUES ('4fae2802-7ad3-4ad3-a605-78685081cda8', current_timestamp);

INSERT INTO users_pdfs (id, user_id, s3_url, unified_at)
  VALUES ('76110992-2296-49e6-8cd5-106b494a66f7', '4fae2802-7ad3-4ad3-a605-78685081cda8', 'url', current_timestamp);

INSERT INTO partial_pdfs (id, unified_pdf_id, source_html_url, number, s3_url, pdf_created_at) VALUES
  ('2dceb213-9c1c-425f-bc74-c264aacdd7d7', '76110992-2296-49e6-8cd5-106b494a66f7', 'https://hoge.com/1', 1, 's3_url', current_timestamp),
  ('d67d37e4-03b7-4fe8-9b82-e92af0076dde', '76110992-2296-49e6-8cd5-106b494a66f7', 'https://hoge.com/2', 2, 's3_url', current_timestamp),
  ('085e2296-e8f3-42b6-b392-0071f76bbc80', '76110992-2296-49e6-8cd5-106b494a66f7', 'https://hoge.com/3', 3, 's3_url', current_timestamp),
  ('99563c3f-e928-4bb6-90fa-56c3c60274d9', '76110992-2296-49e6-8cd5-106b494a66f7', 'https://hoge.com/4', 4, 's3_url', current_timestamp);
