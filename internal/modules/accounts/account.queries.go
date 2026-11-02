package accounts

const SELECT_BY_PID = `
SELECT
	ta.id,
	ta.pid,
	ta.active,
	ta.alias,
	ta.amount,
	ta.user_id,
	tu.pid AS user_pid,
	ta.created_at,
	ta.updated_at
FROM 
	tb_account AS ta
JOIN 
	tb_user AS tu ON ta.user_id = tu.id
WHERE 
	ta.pid = ?;
`

const SELECT_BY_USER_PID = `
SELECT
	ta.id,
	ta.pid,
	ta.active,
	ta.alias,
	ta.amount,
	ta.user_id,
	tu.pid AS user_pid,
	ta.created_at,
	ta.updated_at
FROM 
	tb_account AS ta
JOIN 
	tb_user AS tu ON ta.user_id = tu.id
WHERE 
	tu.pid = ?;
`

const INSERT = `
INSERT INTO tb_account (
	pid,
	alias,
	amount,
	user_id
) VALUES (
	?,
	?,
	?,
	(SELECT id FROM tb_user WHERE pid = ?)
);
`

const UPDATE = `
UPDATE tb_account
SET 
	alias = ?,
	amount = ?,
	updated_at = ?
WHERE 
	pid = ?;
`
