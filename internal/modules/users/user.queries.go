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

const SELECT_TOKEN_BY_PID = `
SELECT
	tu.token
FROM
	tb_user AS tu
WHERE
	tu.pid = ?
	AND tu.active = TRUE;
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

const UPDATE = `
UPDATE tb_user AS tu
SET
	name = ?,
	username = ?,
	email = ?
WHERE
	tu.pid = ?;
`

const UPDATE_ACTIVE = `
UPDATE tb_user AS tu
SET
	active = ?
WHERE
	tu.pid = ?;
`

const UPDATE_PASSWORD = `
UPDATE tb_user AS tu
SET
	password = ?
WHERE
	tu.pid = ?;
`

const UPDATE_TOKEN = `
UPDATE tb_user AS tu
SET
	token = ?
WHERE
	tu.pid = ?;
`
