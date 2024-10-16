-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ? LIMIT 1;

-- name: ListAllTasks :many
SELECT * FROM tasks
ORDER BY id;

-- name: ListUncompletedTasks :many
SELECT * FROM tasks
WHERE completed_timestamp IS NULL
ORDER BY id;

-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = ?
RETURNING id, description;

-- name: CreateTask :one
INSERT INTO tasks (
  description, create_timestamp
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateComplete :one
UPDATE tasks
set completed_timestamp = ?
WHERE id = ?
RETURNING id, description;