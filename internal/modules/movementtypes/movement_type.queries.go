package movementtypes

const SELECT_MOVEMENT_TYPES = `
SELECT
  tmt.pid,
  tmt.active,
  tmt.name,
  tmt.code,
  tmt.description
FROM
  tb_movement_type tmt
WHERE
  tmt.active = true
  AND tmt.deleted_at IS NULL;
`
