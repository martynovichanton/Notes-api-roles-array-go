-- User Queries
-- name: CreateUser :exec
INSERT INTO users (username, password, roles, active) VALUES ($1, $2, $3, $4);

-- name: GetUserByUsername :one
SELECT id, username, password, roles, active FROM users WHERE username = $1;

-- name: GetUsers :many
SELECT id, username, roles, active FROM users;

-- name: UpdateUser :exec
UPDATE users SET username = $1, password = $2, roles = $3, active = $4 WHERE id = $5;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;




-- Note Queries
-- name: CreateNote :exec
INSERT INTO notes (user_id, content) VALUES ($1, $2);

-- name: GetNotesByUserID :many
SELECT id, user_id, content, created_at, updated_at FROM notes WHERE user_id = $1;

-- name: GetNotesByUserIDWithUserNames :many
SELECT notes.id, notes.user_id, notes.content, users.username
FROM notes
INNER JOIN users ON users.id=notes.user_id WHERE user_id = $1;

-- name: GetNotes :many
SELECT id, user_id, content, created_at, updated_at FROM notes;

-- name: GetNotesWithUserNames :many
SELECT notes.id, notes.user_id, notes.content, users.username
FROM notes
INNER JOIN users ON users.id=notes.user_id;

-- name: UpdateNote :exec
UPDATE notes SET content = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2 AND user_id = $3;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1 AND user_id = $2;
