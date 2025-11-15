package service

import (
	"upskilling-api/dao"
	"upskilling-api/model"
)

// TrilhaService é a interface para as operações de negócio de Trilha.
type TrilhaService interface {
	Create(req *model.CreateTrilhaRequest) (*model.TrilhaResponse, error)
	FindByID(id int64) (*model.TrilhaResponse, error)
	FindAll() ([]model.TrilhaResponse, error)
	Update(id int64, req *model.UpdateTrilhaRequest) (*model.TrilhaResponse, error)
	Delete(id int64) error
}

// trilhaServiceImpl implementa a interface TrilhaService.
type trilhaServiceImpl struct {
	dao dao.TrilhaDAO
}

// NewTrilhaService cria uma nova instância de TrilhaService.
func NewTrilhaService() TrilhaService {
	return &trilhaServiceImpl{
		dao: dao.NewTrilhaDAO(),
	}
}

// Create cria uma nova trilha.
func (s *trilhaServiceImpl) Create(req *model.CreateTrilhaRequest) (*model.TrilhaResponse, error) {
	// 1. Mapeamento DTO para Entidade
	trilha := &model.Trilha{
		Nome:          req.Nome,
		Descricao:     req.Descricao,
		Nivel:         req.Nivel,
		CargaHoraria:  req.CargaHoraria,
		FocoPrincipal: req.FocoPrincipal,
	}

	// 2. Persistência
	if err := s.dao.Create(trilha); err != nil {
		return nil, err
	}

	// 3. Mapeamento Entidade para Response DTO
	return &model.TrilhaResponse{
		ID:            trilha.ID,
		Nome:          trilha.Nome,
		Descricao:     trilha.Descricao,
		Nivel:         trilha.Nivel,
		CargaHoraria:  trilha.CargaHoraria,
		FocoPrincipal: trilha.FocoPrincipal,
	}, nil
}

// FindByID busca uma trilha pelo ID.
func (s *trilhaServiceImpl) FindByID(id int64) (*model.TrilhaResponse, error) {
	trilha, err := s.dao.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &model.TrilhaResponse{
		ID:            trilha.ID,
		Nome:          trilha.Nome,
		Descricao:     trilha.Descricao,
		Nivel:         trilha.Nivel,
		CargaHoraria:  trilha.CargaHoraria,
		FocoPrincipal: trilha.FocoPrincipal,
	}, nil
}

// FindAll busca todas as trilhas.
func (s *trilhaServiceImpl) FindAll() ([]model.TrilhaResponse, error) {
	trilhas, err := s.dao.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]model.TrilhaResponse, len(trilhas))
	for i, t := range trilhas {
		responses[i] = model.TrilhaResponse{
			ID:            t.ID,
			Nome:          t.Nome,
			Descricao:     t.Descricao,
			Nivel:         t.Nivel,
			CargaHoraria:  t.CargaHoraria,
			FocoPrincipal: t.FocoPrincipal,
		}
	}
	return responses, nil
}

// Update atualiza uma trilha existente.
func (s *trilhaServiceImpl) Update(id int64, req *model.UpdateTrilhaRequest) (*model.TrilhaResponse, error) {
	// 1. Buscar a trilha existente
	trilha, err := s.dao.FindByID(id)
	if err != nil {
		return nil, err // Trata ResourceNotFoundError
	}

	// 2. Aplicar as atualizações (apenas campos fornecidos)
	if req.Nome != "" {
		trilha.Nome = req.Nome
	}
	if req.Descricao != "" {
		trilha.Descricao = req.Descricao
	}
	if req.Nivel != "" {
		trilha.Nivel = req.Nivel
	}
	if req.CargaHoraria != 0 {
		trilha.CargaHoraria = req.CargaHoraria
	}
	if req.FocoPrincipal != "" {
		trilha.FocoPrincipal = req.FocoPrincipal
	}

	// 3. Persistência
	if err := s.dao.Update(trilha); err != nil {
		return nil, err
	}

	// 4. Mapeamento Entidade para Response DTO
	return &model.TrilhaResponse{
		ID:            trilha.ID,
		Nome:          trilha.Nome,
		Descricao:     trilha.Descricao,
		Nivel:         trilha.Nivel,
		CargaHoraria:  trilha.CargaHoraria,
		FocoPrincipal: trilha.FocoPrincipal,
	}, nil
}

// Delete remove uma trilha pelo ID.
func (s *trilhaServiceImpl) Delete(id int64) error {
	return s.dao.Delete(id)
}
