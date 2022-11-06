package database

import (
	"github.com/cleysonph/hyperprof/internal/model"
)

const findAllProfessoresQuery = `
SELECT
	id,
	nome,
	email,
	idade,
	descricao,
	valor_hora,
	foto_perfil,
	created_at,
	updated_at
FROM
	professores
WHERE
	LOWER(descricao) LIKE CONCAT('%', LOWER(?), '%')
ORDER BY
	created_at ASC
`

func FindAllProfessores(q string) ([]*model.Professor, error) {
	rows, err := db.Query(findAllProfessoresQuery, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	professores := []*model.Professor{}
	for rows.Next() {
		professor := &model.Professor{}
		err := rows.Scan(
			&professor.ID,
			&professor.Nome,
			&professor.Email,
			&professor.Idade,
			&professor.Descricao,
			&professor.ValorHora,
			&professor.FotoPerfil,
			&professor.CreatedAt,
			&professor.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		professores = append(professores, professor)
	}

	return professores, nil
}

const findProfessorByIDQuery = `
SELECT
	id,
	nome,
	email,
	idade,
	descricao,
	valor_hora,
	foto_perfil,
	created_at,
	updated_at
FROM
	professores
WHERE
	id = ?
LIMIT 1
`

func FindProfessorByID(id int64) (*model.Professor, error) {
	row := db.QueryRow(findProfessorByIDQuery, id)
	professor := &model.Professor{}
	err := row.Scan(
		&professor.ID,
		&professor.Nome,
		&professor.Email,
		&professor.Idade,
		&professor.Descricao,
		&professor.ValorHora,
		&professor.FotoPerfil,
		&professor.CreatedAt,
		&professor.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return professor, nil
}

const findProfessorByEmailQuery = `
SELECT
	id,
	nome,
	email,
	idade,
	descricao,
	valor_hora,
	foto_perfil,
	password,
	created_at,
	updated_at
FROM
	professores
WHERE
	email = ?
LIMIT 1
`

func FindProfessorByEmail(email string) (*model.Professor, error) {
	row := db.QueryRow(findProfessorByEmailQuery, email)
	professor := &model.Professor{}
	err := row.Scan(
		&professor.ID,
		&professor.Nome,
		&professor.Email,
		&professor.Idade,
		&professor.Descricao,
		&professor.ValorHora,
		&professor.FotoPerfil,
		&professor.Password,
		&professor.CreatedAt,
		&professor.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return professor, nil
}

const existsProfessorByEmailQuery = `
SELECT
	email
FROM
	professores
WHERE
	email = ?
LIMIT 1
`

func ExistsProfessorByEmail(email string) bool {
	row := db.QueryRow(existsProfessorByEmailQuery, email)
	var professorEmail string
	err := row.Scan(&professorEmail)
	return err == nil && professorEmail == email
}

const existsProfessorByEmailAndNotIDQuery = `
SELECT
	email
FROM
	professores
WHERE
	email = ?
AND
	id != ?
LIMIT 1
`

func ExistsProfessorByEmailAndNotID(email string, id int64) bool {
	row := db.QueryRow(existsProfessorByEmailAndNotIDQuery, email, id)
	var professorEmail string
	err := row.Scan(&professorEmail)
	return err == nil && professorEmail == email
}

const existsProfessorByIDQuery = `
SELECT
	id
FROM
	professores
WHERE
	id = ?
LIMIT 1
`

func ExistsProfessorByID(id int64) bool {
	row := db.QueryRow(existsProfessorByIDQuery, id)
	var professorID int64
	err := row.Scan(&professorID)
	return err == nil && professorID == id
}

const createProfessorQuery = `
INSERT INTO
	professores (nome, email, idade, descricao, valor_hora, password)
VALUES
	(?, ?, ?, ?, ?, ?)
`

func CreateProfessor(professor *model.Professor) (*model.Professor, error) {
	result, err := db.Exec(
		createProfessorQuery,
		professor.Nome,
		professor.Email,
		professor.Idade,
		professor.Descricao,
		professor.ValorHora,
		professor.Password,
	)
	if err != nil {
		return nil, err
	}
	professorID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return FindProfessorByID(professorID)
}

const updateProfessorQuery = `
UPDATE
	professores
SET
	nome = ?,
	email = ?,
	idade = ?,
	descricao = ?,
	valor_hora = ?,
	password = ?
WHERE
	id = ?
`

func UpdateProfessor(professor *model.Professor) (*model.Professor, error) {
	_, err := db.Exec(
		updateProfessorQuery,
		professor.Nome,
		professor.Email,
		professor.Idade,
		professor.Descricao,
		professor.ValorHora,
		professor.Password,
		professor.ID,
	)
	if err != nil {
		return nil, err
	}
	return FindProfessorByID(professor.ID)
}

const updateProfessorFotoPerfilByEmailQuery = `
UPDATE
	professores
SET
	foto_perfil = ?
WHERE
	email = ?
`

func UpdateProfessorFotoPerfilByEmail(email string, fotoPerfil string) error {
	_, err := db.Exec(
		updateProfessorFotoPerfilByEmailQuery,
		fotoPerfil,
		email,
	)
	return err
}

const deleteProfessorByEmailQuery = `
DELETE FROM
	professores
WHERE
	email = ?
`

func DeleteProfessorByEmail(email string) error {
	_, err := db.Exec(deleteProfessorByEmailQuery, email)
	return err
}

const findAlunoByIDQuery = `
SELECT
	id,
	nome,
	email,
	data_aula,
	professor_id,
	created_at,
	updated_at
FROM
	alunos
WHERE
	id = ?
LIMIT 1
`

func FindAlunoByID(id int64) (*model.Aluno, error) {
	row := db.QueryRow(findAlunoByIDQuery, id)
	aluno := &model.Aluno{}
	err := row.Scan(
		&aluno.ID,
		&aluno.Nome,
		&aluno.Email,
		&aluno.DataAula,
		&aluno.ProfessorID,
		&aluno.CreatedAt,
		&aluno.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return aluno, nil
}

const createAlunoQuery = `
INSERT INTO
	alunos (nome, email, data_aula, professor_id)
VALUES
	(?, ?, ?, ?)
`

func CreateAluno(aluno *model.Aluno) (*model.Aluno, error) {
	result, err := db.Exec(createAlunoQuery, aluno.Nome, aluno.Email, aluno.DataAula, aluno.ProfessorID)
	if err != nil {
		return nil, err
	}
	alunoID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return FindAlunoByID(alunoID)
}

const findAlunosByProfessorEmailQuery = `
SELECT
	a.id,
	a.nome,
	a.email,
	a.data_aula,
	a.professor_id,
	a.created_at,
	a.updated_at
FROM
	alunos as a
LEFT JOIN
	professores as p
ON
	a.professor_id = p.id
WHERE
	p.email = ?
`

func FindAlunosByProfessorEmail(email string) ([]*model.Aluno, error) {
	rows, err := db.Query(findAlunosByProfessorEmailQuery, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	alunos := make([]*model.Aluno, 0)
	for rows.Next() {
		aluno := &model.Aluno{}
		err := rows.Scan(
			&aluno.ID,
			&aluno.Nome,
			&aluno.Email,
			&aluno.DataAula,
			&aluno.ProfessorID,
			&aluno.CreatedAt,
			&aluno.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		alunos = append(alunos, aluno)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return alunos, nil
}

const deleteAlunosByProfessorEmailQuery = `
DELETE FROM
	alunos
WHERE
	professor_id = (
		SELECT
			id
		FROM
			professores
		WHERE
			email = ?
	)
`

func DeleteAlunosByProfessorEmail(email string) error {
	_, err := db.Exec(deleteAlunosByProfessorEmailQuery, email)
	return err
}

const CreateInvalidatedTokenQuery = `
INSERT INTO
	invalidated_tokens (token)
VALUES
	(?)
`

func CreateInvalidatedToken(token string) error {
	_, err := db.Exec(CreateInvalidatedTokenQuery, token)
	return err
}

const existsInvalidatedTokenByTokenQuery = `
SELECT
	token
FROM
	invalidated_tokens
WHERE
	token = ?
LIMIT 1
`

func ExistsInvalidatedTokenByToken(token string) bool {
	row := db.QueryRow(existsInvalidatedTokenByTokenQuery, token)
	var invalidatedToken string
	err := row.Scan(&invalidatedToken)
	return err == nil && invalidatedToken == token
}
