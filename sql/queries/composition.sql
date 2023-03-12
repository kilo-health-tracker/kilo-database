-- name: SubmitComposition :one
INSERT INTO tracker.composition (
  SUBMITTED_ON, WEIGHT, BODYFAT
) VALUES (
  $1, $2, $3
)
ON CONFLICT (SUBMITTED_ON) 
DO UPDATE SET 
  WEIGHT = $2,
  BODYFAT = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING *;

-- name: GetComposition :one
SELECT * FROM tracker.composition
WHERE SUBMITTED_ON = $1 LIMIT 1;

-- name: DeleteComposition :exec
DELETE FROM tracker.composition
WHERE SUBMITTED_ON = $1;

-- name: GetCompositionDates :many
SELECT SUBMITTED_ON FROM tracker.composition
LIMIT $1;
