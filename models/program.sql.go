// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: program.sql

package models

import (
	"context"
	"database/sql"
)

const deleteProgram = `-- name: DeleteProgram :exec
DELETE FROM tracker.program
WHERE NAME = $1
`

func (q *Queries) DeleteProgram(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteProgram, name)
	return err
}

const getProgram = `-- name: GetProgram :many
SELECT a.name as program_name, b.name as workout_name, c.group_id, c.exercise_name, c.weight, c.sets, c.reps FROM tracker.program a
JOIN tracker.workout b
ON a.name = b.program_name
JOIN tracker.workout_details c
ON b.name = c.workout_name
WHERE a.name = $1
`

type GetProgramRow struct {
	ProgramName  string        `json:"programName"`
	WorkoutName  string        `json:"workoutName"`
	GroupID      int16         `json:"groupID"`
	ExerciseName string        `json:"exerciseName"`
	Weight       sql.NullInt16 `json:"weight"`
	Sets         int16         `json:"sets"`
	Reps         int16         `json:"reps"`
}

func (q *Queries) GetProgram(ctx context.Context, name string) ([]GetProgramRow, error) {
	rows, err := q.db.QueryContext(ctx, getProgram, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProgramRow
	for rows.Next() {
		var i GetProgramRow
		if err := rows.Scan(
			&i.ProgramName,
			&i.WorkoutName,
			&i.GroupID,
			&i.ExerciseName,
			&i.Weight,
			&i.Sets,
			&i.Reps,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProgramNames = `-- name: GetProgramNames :many
SELECT name FROM tracker.program
LIMIT $1
`

func (q *Queries) GetProgramNames(ctx context.Context, limit int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getProgramNames, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const submitProgram = `-- name: SubmitProgram :one
INSERT INTO tracker.program (
  NAME
) VALUES (
  $1
)
RETURNING name, cret_ts, updt_ts
`

func (q *Queries) SubmitProgram(ctx context.Context, name string) (TrackerProgram, error) {
	row := q.db.QueryRowContext(ctx, submitProgram, name)
	var i TrackerProgram
	err := row.Scan(&i.Name, &i.CretTs, &i.UpdtTs)
	return i, err
}

const submitProgramDetails = `-- name: SubmitProgramDetails :one
INSERT INTO tracker.program_details (
  PROGRAM_NAME, WORKOUT_NAME
) VALUES (
  $1, $2
)
RETURNING program_name, workout_name, cret_ts, updt_ts
`

type SubmitProgramDetailsParams struct {
	ProgramName string `json:"programName"`
	WorkoutName string `json:"workoutName"`
}

func (q *Queries) SubmitProgramDetails(ctx context.Context, arg SubmitProgramDetailsParams) (TrackerProgramDetail, error) {
	row := q.db.QueryRowContext(ctx, submitProgramDetails, arg.ProgramName, arg.WorkoutName)
	var i TrackerProgramDetail
	err := row.Scan(
		&i.ProgramName,
		&i.WorkoutName,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
