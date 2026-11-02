package movementhistory

// TODO: add date filters
const SELECT_USER_HISTORY = `
SELECT
	tmh.pid,
	tmh.amount AS 'history_amount',
	tm.amount AS 'amount',
	tm.code AS 'movement_code',
	ta.alias AS 'account',
	tmt.name AS 'type',
	tmt.code AS 'type_code',
	tu.username,
	tu.email,
	tmh.created_at AS 'date'
FROM
	tb_movement_history tmh
INNER JOIN tb_movement tm ON tmh.movement_id = tm.id
INNER JOIN tb_account ta ON tm.account_id = ta.id
INNER JOIN tb_movement_type tmt ON tm.movement_type_id = tmt.id
INNER JOIN tb_user tu ON tm.user_id = tu.id AND tu.pid = ?;
`
