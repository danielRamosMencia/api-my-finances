package movementhistory

// Dates are compared with substr(created_at, 1, 19): the "YYYY-MM-DDTHH:MM:SS" prefix of the stored UTC string.
// SQLite's date functions round nanoseconds, e.g. date('2026-10-31T23:59:59.999999999Z') = '2026-11-01'.

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
INNER JOIN tb_user tu ON tm.user_id = tu.id AND tu.pid = ?
WHERE
	substr(tmh.created_at, 1, 19) BETWEEN ? AND ?
ORDER BY
	tmh.created_at DESC;
`

const SELECT_USER_METRICS = `
SELECT
	COUNT(tm.id) AS 'total_movements',
	ROUND(COALESCE(SUM(CASE WHEN tmt.code = 'INCOME' THEN tm.amount ELSE 0 END), 0), 2) AS 'total_income',
	ROUND(COALESCE(SUM(CASE WHEN tmt.code <> 'INCOME' THEN tm.amount ELSE 0 END), 0), 2) AS 'total_expense',
	ROUND(COALESCE(SUM(CASE WHEN tmt.code = 'INCOME' THEN tm.amount ELSE -tm.amount END), 0), 2) AS 'net'
FROM
	tb_movement tm
INNER JOIN tb_movement_type tmt ON tm.movement_type_id = tmt.id
INNER JOIN tb_user tu ON tm.user_id = tu.id AND tu.pid = ?
WHERE
	tm.deleted_at IS NULL
	AND substr(tm.created_at, 1, 19) BETWEEN ? AND ?;
`

// LEFT JOINs so accounts without movements in the month are returned with zeros
const SELECT_USER_ACCOUNT_METRICS = `
SELECT
	ta.pid AS 'account_pid',
	ta.alias AS 'account',
	COUNT(tm.id) AS 'total_movements',
	ROUND(COALESCE(SUM(CASE WHEN tmt.code = 'INCOME' THEN tm.amount ELSE 0 END), 0), 2) AS 'total_income',
	ROUND(COALESCE(SUM(CASE WHEN tmt.code <> 'INCOME' THEN tm.amount ELSE 0 END), 0), 2) AS 'total_expense',
	ROUND(COALESCE(SUM(CASE WHEN tmt.code = 'INCOME' THEN tm.amount WHEN tmt.code IS NOT NULL THEN -tm.amount ELSE 0 END), 0), 2) AS 'net'
FROM
	tb_account ta
INNER JOIN tb_user tu ON ta.user_id = tu.id AND tu.pid = ?
LEFT JOIN tb_movement tm ON tm.account_id = ta.id
	AND tm.deleted_at IS NULL
	AND substr(tm.created_at, 1, 19) BETWEEN ? AND ?
LEFT JOIN tb_movement_type tmt ON tm.movement_type_id = tmt.id
WHERE
	ta.deleted_at IS NULL
GROUP BY
	ta.id, ta.pid, ta.alias
ORDER BY
	ta.alias;
`
