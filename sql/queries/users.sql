-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, DEFAULT, DEFAULT, $2)
RETURNING id, created_at, updated_at, name;