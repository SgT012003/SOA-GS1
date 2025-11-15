package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"upskilling-api/db"
	"upskilling-api/model"
)

// UsuarioDAO é a interface para as operações de acesso a dados de Usuário.
type UsuarioDAO interface {
	Create(usuario *model.Usuario) error
	FindByID(id int64) (*model.Usuario, error)
	FindAll() ([]model.Usuario, error)
	Update(usuario *model.Usuario) error
	Delete(id int64) error
	FindByEmail(email string) (*model.Usuario, error)
}

// usuarioDAOImpl implementa a interface UsuarioDAO.
type usuarioDAOImpl struct{}

// NewUsuarioDAO cria uma nova instância de UsuarioDAO.
func NewUsuarioDAO() UsuarioDAO {
	return &usuarioDAOImpl{}
}

// Create insere um novo usuário no banco de dados.
func (d *usuarioDAOImpl) Create(usuario *model.Usuario) error {
	query := `
		INSERT INTO usuarios (nome, email, area_atuacao, nivel_carreira, data_cadastro)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, data_cadastro
	`
	err := db.GetDB().QueryRow(
		query,
		usuario.Nome,
		usuario.Email,
		usuario.AreaAtuacao,
		usuario.NivelCarreira,
		time.Now(),
	).Scan(&usuario.ID, &usuario.DataCadastro)

	if err != nil {
		// Tratar erro de email duplicado (Postgres)
		if err.Error() == "pq: duplicate key value violates unique constraint \"usuarios_email_key\"" {
			return &model.ConflictError{Msg: "Email já cadastrado."}
		}
		log.Printf("Erro ao criar usuário: %v", err)
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}
	return nil
}

// FindByID busca um usuário pelo ID.
func (d *usuarioDAOImpl) FindByID(id int64) (*model.Usuario, error) {
	usuario := &model.Usuario{}
	query := `
		SELECT id, nome, email, area_atuacao, nivel_carreira, data_cadastro
		FROM usuarios
		WHERE id = $1
	`
	err := db.GetDB().QueryRow(query, id).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Email,
		&usuario.AreaAtuacao,
		&usuario.NivelCarreira,
		&usuario.DataCadastro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.ResourceNotFoundError{Resource: "Usuário", ID: id}
		}
		log.Printf("Erro ao buscar usuário por ID: %v", err)
		return nil, fmt.Errorf("erro ao buscar usuário por ID: %w", err)
	}
	return usuario, nil
}

// FindByEmail busca um usuário pelo email.
func (d *usuarioDAOImpl) FindByEmail(email string) (*model.Usuario, error) {
	usuario := &model.Usuario{}
	query := `
		SELECT id, nome, email, area_atuacao, nivel_carreira, data_cadastro
		FROM usuarios
		WHERE email = $1
	`
	err := db.GetDB().QueryRow(query, email).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Email,
		&usuario.AreaAtuacao,
		&usuario.NivelCarreira,
		&usuario.DataCadastro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Retorna nil, nil se não encontrar
		}
		log.Printf("Erro ao buscar usuário por email: %v", err)
		return nil, fmt.Errorf("erro ao buscar usuário por email: %w", err)
	}
	return usuario, nil
}

// FindAll busca todos os usuários.
func (d *usuarioDAOImpl) FindAll() ([]model.Usuario, error) {
	rows, err := db.GetDB().Query(`
		SELECT id, nome, email, area_atuacao, nivel_carreira, data_cadastro
		FROM usuarios
		ORDER BY id
	`)
	if err != nil {
		log.Printf("Erro ao buscar todos os usuários: %v", err)
		return nil, fmt.Errorf("erro ao buscar todos os usuários: %w", err)
	}
	defer rows.Close()

	usuarios := make([]model.Usuario, 0)
	for rows.Next() {
		usuario := model.Usuario{}
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.AreaAtuacao,
			&usuario.NivelCarreira,
			&usuario.DataCadastro,
		)
		if err != nil {
			log.Printf("Erro ao escanear linha de usuário: %v", err)
			return nil, fmt.Errorf("erro ao escanear linha de usuário: %w", err)
		}
		usuarios = append(usuarios, usuario)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração de linhas: %v", err)
		return nil, fmt.Errorf("erro após iteração de linhas: %w", err)
	}

	return usuarios, nil
}

// Update atualiza um usuário existente.
func (d *usuarioDAOImpl) Update(usuario *model.Usuario) error {
	query := `
		UPDATE usuarios
		SET nome = $2, area_atuacao = $3, nivel_carreira = $4
		WHERE id = $1
	`
	result, err := db.GetDB().Exec(
		query,
		usuario.ID,
		usuario.Nome,
		usuario.AreaAtuacao,
		usuario.NivelCarreira,
	)
	if err != nil {
		log.Printf("Erro ao atualizar usuário: %v", err)
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return &model.ResourceNotFoundError{Resource: "Usuário", ID: usuario.ID}
	}

	return nil
}

// Delete remove um usuário pelo ID.
func (d *usuarioDAOImpl) Delete(id int64) error {
	result, err := db.GetDB().Exec("DELETE FROM usuarios WHERE id = $1", id)
	if err != nil {
		log.Printf("Erro ao deletar usuário: %v", err)
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return &model.ResourceNotFoundError{Resource: "Usuário", ID: id}
	}

	return nil
}
