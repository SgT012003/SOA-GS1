package dao

import (
	"fmt"
	"log"
	"time"

	"upskilling-api/db"
	"upskilling-api/model"
)

// MatriculaDAO é a interface para as operações de acesso a dados de Matrícula.
type MatriculaDAO interface {
	Create(matricula *model.Matricula) error
	FindByUsuarioID(usuarioID int64) ([]model.Matricula, error)
	// Adicionar métodos para buscar por TrilhaID, etc., se necessário
}

// matriculaDAOImpl implementa a interface MatriculaDAO.
type matriculaDAOImpl struct{}

// NewMatriculaDAO cria uma nova instância de MatriculaDAO.
func NewMatriculaDAO() MatriculaDAO {
	return &matriculaDAOImpl{}
}

// Create insere uma nova matrícula no banco de dados.
func (d *matriculaDAOImpl) Create(matricula *model.Matricula) error {
	query := `
		INSERT INTO matriculas (usuario_id, trilha_id, data_inscricao, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, data_inscricao
	`
	err := db.GetDB().QueryRow(
		query,
		matricula.UsuarioID,
		matricula.TrilhaID,
		time.Now(),
		"ATIVA", // Status inicial
	).Scan(&matricula.ID, &matricula.DataInscricao)

	if err != nil {
		log.Printf("Erro ao criar matrícula: %v", err)
		return fmt.Errorf("erro ao criar matrícula: %w", err)
	}
	return nil
}

// FindByUsuarioID busca todas as matrículas de um usuário.
func (d *matriculaDAOImpl) FindByUsuarioID(usuarioID int64) ([]model.Matricula, error) {
	rows, err := db.GetDB().Query(`
		SELECT id, usuario_id, trilha_id, data_inscricao, status
		FROM matriculas
		WHERE usuario_id = $1
		ORDER BY data_inscricao DESC
	`, usuarioID)
	if err != nil {
		log.Printf("Erro ao buscar matrículas por usuário: %v", err)
		return nil, fmt.Errorf("erro ao buscar matrículas por usuário: %w", err)
	}
	defer rows.Close()

	matriculas := make([]model.Matricula, 0)
	for rows.Next() {
		matricula := model.Matricula{}
		err := rows.Scan(
			&matricula.ID,
			&matricula.UsuarioID,
			&matricula.TrilhaID,
			&matricula.DataInscricao,
			&matricula.Status,
		)
		if err != nil {
			log.Printf("Erro ao escanear linha de matrícula: %v", err)
			return nil, fmt.Errorf("erro ao escanear linha de matrícula: %w", err)
		}
		matriculas = append(matriculas, matricula)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração de linhas: %v", err)
		return nil, fmt.Errorf("erro após iteração de linhas: %w", err)
	}

	return matriculas, nil
}
