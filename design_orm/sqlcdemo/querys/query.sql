-- name: GetStudentById :one
SELECT * FROM students WHERE id=? LIMIT 1;

-- name: GetStudentLikeNo :one
SELECT * FROM students WHERE student_no like ? ;


-- name: GetUsers :one
SELECT * FROM users WHERE id=? LIMIT 1;


