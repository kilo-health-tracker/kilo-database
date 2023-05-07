-- name: SubmitBodyPart :one
INSERT INTO tracker.body_parts (
  NAME, REGION, UPPER_OR_LOWER
) VALUES (
  $1, $2, $3
)
ON CONFLICT (NAME) 
DO UPDATE SET 
  REGION = $2,
  UPPER_OR_LOWER = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING *;

-- name: GetBodyPart :one
SELECT * FROM tracker.body_parts
WHERE NAME = $1 LIMIT 1;

-- name: DeleteBodyPart :exec
DELETE FROM tracker.body_parts
WHERE NAME = $1;