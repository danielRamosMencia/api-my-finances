package movements

const SELECT_USER_MOVEMENTS = `
SELECT
	tm.id,
	tm.pid,
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
	tu.pid = ?;
`
