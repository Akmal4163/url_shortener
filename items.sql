PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE items ( id INTEGER PRIMARY KEY AUTOINCREMENT, long_url TEXT NOT NULL, short_url TEXT NOT NULL, timestamp TEXT NOT NULL);
DELETE FROM sqlite_sequence;
COMMIT;
