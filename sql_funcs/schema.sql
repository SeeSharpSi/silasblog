CREATE TABLE articles (
    article_id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    clicks INTEGER DEFAULT 0
);
