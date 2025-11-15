package service

import (
	"fmt"

	"upskilling-api/dao"
	"upskilling-api/model"
)

// MatriculaService é a interface para as operações de negócio de Matrícula.
type MatriculaService interface {
	Matricular(usuarioID, trilhaID int64) (*model.Matricula, error)
	GetMatriculasByUsuario(usuarioID int64) ([]model.Matricula, error)
}

// matriculaServiceImpl implementa a interface MatriculaService.
type matriculaServiceImpl struct {
	matriculaDAO dao.MatriculaDAO
	usuarioDAO   dao.UsuarioDAO
	trilhaDAO    dao.TrilhaDAO
}

// NewMatriculaService cria uma nova instância de MatriculaService.
func NewMatriculaService() MatriculaService {
	return &matriculaServiceImpl{
		matriculaDAO: dao.NewMatriculaDAO(),
		usuarioDAO:   dao.NewUsuarioDAO(),
		trilhaDAO:    dao.NewTrilhaDAO(),
	}
}

// Matricular realiza a inscrição de um usuário em uma trilha.
func (s *matriculaServiceImpl) Matricular(usuarioID, trilhaID int64) (*model.Matricula, error) {
	// 1. Validação de Existência: Usuário
	_, err := s.usuarioDAO.FindByID(usuarioID)
	if err != nil {
		// Se for ResourceNotFoundError, retorna o erro
		if _, ok := err.(*model.ResourceNotFoundError); ok {
			return nil, &model.BusinessRuleError{Msg: fmt.Sprintf("Usuário com ID %d não encontrado.", usuarioID)}
		}
		return nil, err
	}

	// 2. Validação de Existência: Trilha
	_, err = s.trilhaDAO.FindByID(trilhaID)
	if err != nil {
		// Se for ResourceNotFoundError, retorna o erro
		if _, ok := err.(*model.ResourceNotFoundError); ok {
			return nil, &model.BusinessRuleError{Msg: fmt.Sprintf("Trilha com ID %d não encontrada.", trilhaID)}
		}
		return nil, err
	}

	// 3. Validação de Negócio: Verificar se o usuário já está matriculado (simplificado)
	// Em um cenário real, seria necessário verificar se já existe uma matrícula ATIVA.
	// Por simplicidade, vamos apenas criar a matrícula.

	// 4. Criação da Matrícula
	matricula := &model.Matricula{
		UsuarioID: usuarioID,
		TrilhaID:  trilhaID,
		Status:    "ATIVA",
	}

	if err := s.matriculaDAO.Create(matricula); err != nil {
		return nil, err
	}

	return matricula, nil
}

// GetMatriculasByUsuario busca todas as matrículas de um usuário.
func (s *matriculaServiceImpl) GetMatriculasByUsuario(usuarioID int64) ([]model.Matricula, error) {
	// 1. Validação de Existência: Usuário
	_, err := s.usuarioDAO.FindByID(usuarioID)
	if err != nil {
		// Se for ResourceNotFoundError, retorna o erro
		if _, ok := err.(*model.ResourceNotFoundError); ok {
			return nil, &model.BusinessRuleError{Msg: fmt.Sprintf("Usuário com ID %d não encontrado.", usuarioID)}
		}
		return nil, err
	}

	// 2. Busca no DAO
	return s.matriculaDAO.FindByUsuarioID(usuarioID)
}
