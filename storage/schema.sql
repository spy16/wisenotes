-- SQLite3 Schema

CREATE TABLE IF NOT EXISTS profile (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    emoji TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS config (
    key TEXT PRIMARY KEY,
    value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS article (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    profile_id INTEGER NOT NULL,
    title TEXT UNIQUE NOT NULL,
    kind TEXT NOT NULL,
    spec TEXT NOT NULL,
    content TEXT NOT NULL,
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (profile_id) REFERENCES profile(id)
);

CREATE INDEX IF NOT EXISTS idx_article_kind ON article (kind);
CREATE INDEX IF NOT EXISTS idx_article_created_at ON article (created_at);

CREATE TABLE IF NOT EXISTS article_tag (
    article_id INTEGER,
    tag TEXT,
    FOREIGN KEY (article_id) REFERENCES article(id)
);

CREATE INDEX IF NOT EXISTS idx_article_tag_article_id ON article_tag (article_id);

CREATE TABLE IF NOT EXISTS article_chunk (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    article_id INTEGER,
    offset INTEGER,
    length INTEGER,
    FOREIGN KEY (article_id) REFERENCES article(id)
);

