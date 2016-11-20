CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY, name TEXT);
CREATE TABLE IF NOT EXISTS postings (id INTEGER PRIMARY KEY, categoryId INTEGER, cookie TEXT, caption TEXT, content TEXT, phonenumber TEXT, created INTEGER);
CREATE TABLE IF NOT EXISTS favorites (id INTEGER PRIMARY KEY, cookie TEXT, postingid INTEGER);