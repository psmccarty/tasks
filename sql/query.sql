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

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;

-- name: CreateTask :one
INSERT INTO tasks (
  description, create_timestamp
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateComplete :exec
UPDATE tasks
set completed_timestamp = ?
WHERE id = ?;