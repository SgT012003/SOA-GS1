package service

import (
	"fmt"
	"time"

	"upskilling-api/dao"
	"upskilling-api/model"
)

// UsuarioService é a interface para as operações de negócio de Usuário.
type UsuarioService interface {
	Create(req *model.CreateUsuarioRequest) (*model.UsuarioResponse, error)
	FindByID(id int64) (*model.UsuarioResponse, error)
	FindAll() ([]model.UsuarioResponse, error)
	Update(id int64, req *model.UpdateUsuarioRequest) (*model.UsuarioResponse, error)
	Delete(id int64) error
}

// usuarioServiceImpl implementa a interface UsuarioService.
type usuarioServiceImpl struct {
	dao dao.UsuarioDAO
}

// NewUsuarioService cria uma nova instância de UsuarioService.
func NewUsuarioService() UsuarioService {
	return &usuarioServiceImpl{
		dao: dao.NewUsuarioDAO(),
	}
}

// Create cria um novo usuário após validações de negócio.
func (s *usuarioServiceImpl) Create(req *model.CreateUsuarioRequest) (*model.UsuarioResponse, error) {
	// 1. Validação de Negócio: Verificar se o email já existe
	existingUser, err := s.dao.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, &model.ConflictError{Msg: fmt.Sprintf("O email '%s' já está cadastrado.", req.Email)}
	}

	// 2. Mapeamento DTO para Entidade
	usuario := &model.Usuario{
		Nome:          req.Nome,
		Email:         req.Email,
		AreaAtuacao:   req.AreaAtuacao,
		NivelCarreira: req.NivelCarreira,
		DataCadastro:  time.Now(), // Será sobrescrito pelo valor do DB, mas é bom ter um default
	}

	// 3. Persistência
	if err := s.dao.Create(usuario); err != nil {
		return nil, err
	}

	// 4. Mapeamento Entidade para Response DTO
	return &model.UsuarioResponse{
		ID:            usuario.ID,
		Nome:          usuario.Nome,
		Email:         usuario.Email,
		AreaAtuacao:   usuario.AreaAtuacao,
		NivelCarreira: usuario.NivelCarreira,
		DataCadastro:  usuario.DataCadastro,
	}, nil
}

// FindByID busca um usuário pelo ID.
func (s *usuarioServiceImpl) FindByID(id int64) (*model.UsuarioResponse, error) {
	usuario, err := s.dao.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &model.UsuarioResponse{
		ID:            usuario.ID,
		Nome:          usuario.Nome,
		Email:         usuario.Email,
		AreaAtuacao:   usuario.AreaAtuacao,
		NivelCarreira: usuario.NivelCarreira,
		DataCadastro:  usuario.DataCadastro,
	}, nil
}

// FindAll busca todos os usuários.
func (s *usuarioServiceImpl) FindAll() ([]model.UsuarioResponse, error) {
	usuarios, err := s.dao.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]model.UsuarioResponse, len(usuarios))
	for i, u := range usuarios {
		responses[i] = model.UsuarioResponse{
			ID:            u.ID,
			Nome:          u.Nome,
			Email:         u.Email,
			AreaAtuacao:   u.AreaAtuacao,
			NivelCarreira: u.NivelCarreira,
			DataCadastro:  u.DataCadastro,
		}
	}
	return responses, nil
}

// Update atualiza um usuário existente.
func (s *usuarioServiceImpl) Update(id int64, req *model.UpdateUsuarioRequest) (*model.UsuarioResponse, error) {
	// 1. Buscar o usuário existente
	usuario, err := s.dao.FindByID(id)
	if err != nil {
		return nil, err // Trata ResourceNotFoundError
	}

	// 2. Aplicar as atualizações (apenas campos fornecidos)
	if req.Nome != "" {
		usuario.Nome = req.Nome
	}
	if req.AreaAtuacao != "" {
		usuario.AreaAtuacao = req.AreaAtuacao
	}
	if req.NivelCarreira != "" {
		usuario.NivelCarreira = req.NivelCarreira
	}

	// 3. Persistência
	if err := s.dao.Update(usuario); err != nil {
		return nil, err
	}

	// 4. Mapeamento Entidade para Response DTO
	return &model.UsuarioResponse{
		ID:            usuario.ID,
		Nome:          usuario.Nome,
		Email:         usuario.Email,
		AreaAtuacao:   usuario.AreaAtuacao,
		NivelCarreira: usuario.NivelCarreira,
		DataCadastro:  usuario.DataCadastro,
	}, nil
}

// Delete remove um usuário pelo ID.
func (s *usuarioServiceImpl) Delete(id int64) error {
	return s.dao.Delete(id)
}
