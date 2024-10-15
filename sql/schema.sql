CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  description TEXT NOT NULL,
  create_timestamp DATETIME NOT NULL,
  completed_timestamp DATETIME,
  due_date_timestamp DATETIME
);