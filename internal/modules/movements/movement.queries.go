package movements

const SELECT_USER_MOVEMENTS = `
SELECT
	tm.id,
	tm.pid,
	tm.code,
	tm.description,
	tm.active,
	tm.amount,
	tm.account_id,
	ta.pid as account_pid,
	ta.alias as account_alias,
	tm.movement_type_id,
	tmt.pid as movement_type_pid,
	tmt.name as movement_type_name,
	tm.user_id,
	tu.pid as user_pid,
	tu.username,
	tm.created_at,
	tm.updated_at,
	tm.deleted_at
FROM
	tb_movement tm
JOIN tb_user tu ON tm.user_id = tu.id
JOIN tb_account ta ON tm.account_id = ta.id
JOIN tb_movement_type tmt ON tm.movement_type_id = tmt.id
WHERE
	tu.pid = ?
	AND substr(tm.created_at, 1, 19) BETWEEN ? AND ?
ORDER BY
	tm.created_at DESC;
`

const SELECT_BY_PID = `
SELECT
	tm.id,
	tm.pid,
	tm.code,
	tm.description,
	tm.active,
	tm.amount,
	tm.account_id,
	ta.pid as account_pid,
	ta.alias as account_alias,
	tm.movement_type_id,
	tmt.pid as movement_type_pid,
	tmt.name as movement_type_name,
	tm.user_id,
	tu.pid as user_pid,
	tu.username,
	tm.created_at,
	tm.updated_at,
	tm.deleted_at
FROM
	tb_movement tm
JOIN tb_user tu ON tm.user_id = tu.id
JOIN tb_account ta ON tm.account_id = ta.id
JOIN tb_movement_type tmt ON tm.movement_type_id = tmt.id
WHERE
	tm.pid = ?;
`

const SELECT_ACCOUNT_ID = `
SELECT 
	ta.id 
FROM
	tb_account ta WHERE ta.pid = ?
	AND ta.user_id = (SELECT id FROM tb_user WHERE pid = ?)
;
`

const SELECT_MOVEMENT_TYPE_ID = `
SELECT
	tmt.id
FROM
	tb_movement_type tmt WHERE tmt.pid = ?;
`

const SELECT_USER_ID = `
SELECT
	tu.id
FROM
	tb_user tu WHERE tu.pid = ?;
`

const INSERT = `
INSERT INTO tb_movement (
	pid,
	active,
	code,
	description,
	amount,
	account_id,
	movement_type_id,
	user_id,
	created_at,
	updated_at
) VALUES (
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?
);
`

const UPDATE_ACCOUNT_AMOUNT = `
UPDATE
    tb_account
SET
    amount = ROUND(amount - ?, 2),
    updated_at = ?
WHERE
    id = ?;
`

const INSERT_MOVEMENT_HISTORY = `
INSERT INTO tb_movement_history (
	pid,
	amount,
	movement_id,
	created_at
) VALUES (
	?, 
	?, 
	?,
	?
)
`

const SELECT_MOVEMENT_TYPE_CODE = `
SELECT
	tmt.code
FROM
	tb_movement_type tmt WHERE tmt.pid = ?;
`
