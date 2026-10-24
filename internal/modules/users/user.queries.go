package users

const SELECT_ALL = `
SELECT
	pid,
	active,
	name,
	username,
	email,
	type,
	created_at,
	updated_at
FROM
	tb_user;
`

const SELECT_BY_PID = `
SELECT
	pid,
	active,
	name,
	username,
	email,
	type,
	created_at,
	updated_at
FROM
	tb_user AS tu
WHERE
	tu.pid = ?;
`

const INSERT = `
INSERT INTO tb_user (
	pid,
	name,
	username,
	email,
	token,
	password,
	type
) VALUES (
	?,
	?,
	?,
	?,
	?,
	?,
	?
);
`
