package dao

import (
	"database/sql"
	"fmt"
	"log"

	"upskilling-api/db"
	"upskilling-api/model"
)

// TrilhaDAO é a interface para as operações de acesso a dados de Trilha.
type TrilhaDAO interface {
	Create(trilha *model.Trilha) error
	FindByID(id int64) (*model.Trilha, error)
	FindAll() ([]model.Trilha, error)
	Update(trilha *model.Trilha) error
	Delete(id int64) error
}

// trilhaDAOImpl implementa a interface TrilhaDAO.
type trilhaDAOImpl struct{}

// NewTrilhaDAO cria uma nova instância de TrilhaDAO.
func NewTrilhaDAO() TrilhaDAO {
	return &trilhaDAOImpl{}
}

// Create insere uma nova trilha no banco de dados.
func (d *trilhaDAOImpl) Create(trilha *model.Trilha) error {
	query := `
		INSERT INTO trilhas (nome, descricao, nivel, carga_horaria, foco_principal)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := db.GetDB().QueryRow(
		query,
		trilha.Nome,
		trilha.Descricao,
		trilha.Nivel,
		trilha.CargaHoraria,
		trilha.FocoPrincipal,
	).Scan(&trilha.ID)

	if err != nil {
		log.Printf("Erro ao criar trilha: %v", err)
		return fmt.Errorf("erro ao criar trilha: %w", err)
	}
	return nil
}

// FindByID busca uma trilha pelo ID.
func (d *trilhaDAOImpl) FindByID(id int64) (*model.Trilha, error) {
	trilha := &model.Trilha{}
	query := `
		SELECT id, nome, descricao, nivel, carga_horaria, foco_principal
		FROM trilhas
		WHERE id = $1
	`
	err := db.GetDB().QueryRow(query, id).Scan(
		&trilha.ID,
		&trilha.Nome,
		&trilha.Descricao,
		&trilha.Nivel,
		&trilha.CargaHoraria,
		&trilha.FocoPrincipal,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.ResourceNotFoundError{Resource: "Trilha", ID: id}
		}
		log.Printf("Erro ao buscar trilha por ID: %v", err)
		return nil, fmt.Errorf("erro ao buscar trilha por ID: %w", err)
	}
	return trilha, nil
}

// FindAll busca todas as trilhas.
func (d *trilhaDAOImpl) FindAll() ([]model.Trilha, error) {
	rows, err := db.GetDB().Query(`
		SELECT id, nome, descricao, nivel, carga_horaria, foco_principal
		FROM trilhas
		ORDER BY id
	`)
	if err != nil {
		log.Printf("Erro ao buscar todas as trilhas: %v", err)
		return nil, fmt.Errorf("erro ao buscar todas as trilhas: %w", err)
	}
	defer rows.Close()

	trilhas := make([]model.Trilha, 0)
	for rows.Next() {
		trilha := model.Trilha{}
		err := rows.Scan(
			&trilha.ID,
			&trilha.Nome,
			&trilha.Descricao,
			&trilha.Nivel,
			&trilha.CargaHoraria,
			&trilha.FocoPrincipal,
		)
		if err != nil {
			log.Printf("Erro ao escanear linha de trilha: %v", err)
			return nil, fmt.Errorf("erro ao escanear linha de trilha: %w", err)
		}
		trilhas = append(trilhas, trilha)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração de linhas: %v", err)
		return nil, fmt.Errorf("erro após iteração de linhas: %w", err)
	}

	return trilhas, nil
}

// Update atualiza uma trilha existente.
func (d *trilhaDAOImpl) Update(trilha *model.Trilha) error {
	query := `
		UPDATE trilhas
		SET nome = $2, descricao = $3, nivel = $4, carga_horaria = $5, foco_principal = $6
		WHERE id = $1
	`
	result, err := db.GetDB().Exec(
		query,
		trilha.ID,
		trilha.Nome,
		trilha.Descricao,
		trilha.Nivel,
		trilha.CargaHoraria,
		trilha.FocoPrincipal,
	)
	if err != nil {
		log.Printf("Erro ao atualizar trilha: %v", err)
		return fmt.Errorf("erro ao atualizar trilha: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return &model.ResourceNotFoundError{Resource: "Trilha", ID: trilha.ID}
	}

	return nil
}

// Delete remove uma trilha pelo ID.
func (d *trilhaDAOImpl) Delete(id int64) error {
	result, err := db.GetDB().Exec("DELETE FROM trilhas WHERE id = $1", id)
	if err != nil {
		log.Printf("Erro ao deletar trilha: %v", err)
		return fmt.Errorf("erro ao deletar trilha: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return &model.ResourceNotFoundError{Resource: "Trilha", ID: id}
	}

	return nil
}
