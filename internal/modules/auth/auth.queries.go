package auth

const SELECT_USER_BY_USERNAME = `
SELECT
	tu.name,
	tu.username,
	tu.pid,
	tu.email,
	tu.password
FROM 
	tb_user tu
WHERE
	tu.username = ?
	AND tu.active = TRUE
	AND tu.deleted_at IS NULL;
`
