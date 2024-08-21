CREATE TABLE IF NOT EXISTS json_data (
    key TEXT PRIMARY KEY,
    data TEXT NOT NULL
);

INSERT INTO json_data (key, data)
VALUES ('profile', '{"name": "Goos van den Bekerom", "email": "goos.bekerom@gmail.com", "date_of_birth": "1995-06-07T00:00:00Z"}')