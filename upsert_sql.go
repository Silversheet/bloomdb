package bloomdb

var upsertSql = `
WITH upd AS
  (UPDATE {{.Table}} t
   SET 
      {{range $i, $e := .Columns}}{{$e}} = s.{{$e}}{{if len $.Columns | sub 1 | eq $i | not}},{{end}}
{{end}}
   FROM {{.Table}}_temp s
   WHERE s.id = t.id
   RETURNING s.id)

INSERT INTO {{.Table}}(
  {{range $i, $e := .Columns}}{{$e}}{{if len $.Columns | sub 1 | eq $i | not}},{{end}}
{{end}}
)
SELECT DISTINCT ON (s.id)
  {{range $i, $e := .Columns}}{{$e}}{{if len $.Columns | sub 1 | eq $i | not}},{{end}}
{{end}}
FROM {{.Table}}_temp s
LEFT JOIN upd t USING(id)
WHERE t.id IS NULL
RETURNING {{.Table}}.id
`