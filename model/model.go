package model

import (
	"fmt"
	"time"
)

// DTOs de Erro
type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// --------------------------------------------------------------------------------
// Entidades Principais (Mapeamento DB)
// --------------------------------------------------------------------------------

// Usuario representa um profissional/aluno na plataforma.
type Usuario struct {
	ID            int64     `json:"id"`
	Nome          string    `json:"nome"`
	Email         string    `json:"email"`
	AreaAtuacao   string    `json:"area_atuacao,omitempty"`
	NivelCarreira string    `json:"nivel_carreira,omitempty"`
	DataCadastro  time.Time `json:"data_cadastro"`
}

// Trilha representa uma trilha de aprendizagem.
type Trilha struct {
	ID            int64  `json:"id"`
	Nome          string `json:"nome"`
	Descricao     string `json:"descricao,omitempty"`
	Nivel         string `json:"nivel"`                    // INICIANTE, INTERMEDIARIO, AVANCADO
	CargaHoraria  int    `json:"carga_horaria"`            // em horas
	FocoPrincipal string `json:"foco_principal,omitempty"` // IA, Dados, Soft Skills, Green Tech
}

// Competencia representa uma skill do futuro do trabalho.
type Competencia struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Categoria string `json:"categoria,omitempty"` // Tecnologia, Humana, Gestão
	Descricao string `json:"descricao,omitempty"`
}

// Matricula representa a inscrição de um usuário em uma trilha.
type Matricula struct {
	ID            int64     `json:"id"`
	UsuarioID     int64     `json:"usuario_id"`
	TrilhaID      int64     `json:"trilha_id"`
	DataInscricao time.Time `json:"data_inscricao"`
	Status        string    `json:"status"` // ATIVA, CONCLUIDA, CANCELADA
}

// --------------------------------------------------------------------------------
// DTOs de Requisição (Input)
// --------------------------------------------------------------------------------

// CreateUsuarioRequest é o DTO para criar um novo usuário.
type CreateUsuarioRequest struct {
	Nome          string `json:"nome" binding:"required,min=3,max=100"`
	Email         string `json:"email" binding:"required,email"`
	AreaAtuacao   string `json:"area_atuacao,omitempty" binding:"max=100"`
	NivelCarreira string `json:"nivel_carreira,omitempty" binding:"max=50"`
}

// UpdateUsuarioRequest é o DTO para atualizar um usuário existente.
type UpdateUsuarioRequest struct {
	Nome          string `json:"nome,omitempty" binding:"omitempty,min=3,max=100"`
	AreaAtuacao   string `json:"area_atuacao,omitempty" binding:"max=100"`
	NivelCarreira string `json:"nivel_carreira,omitempty" binding:"max=50"`
}

// CreateTrilhaRequest é o DTO para criar uma nova trilha.
type CreateTrilhaRequest struct {
	Nome          string `json:"nome" binding:"required,min=5,max=150"`
	Descricao     string `json:"descricao,omitempty"`
	Nivel         string `json:"nivel" binding:"required,oneof=INICIANTE INTERMEDIARIO AVANCADO"`
	CargaHoraria  int    `json:"carga_horaria" binding:"required,gt=0"`
	FocoPrincipal string `json:"foco_principal,omitempty" binding:"max=100"`
}

// UpdateTrilhaRequest é o DTO para atualizar uma trilha existente.
type UpdateTrilhaRequest struct {
	Nome          string `json:"nome,omitempty" binding:"omitempty,min=5,max=150"`
	Descricao     string `json:"descricao,omitempty"`
	Nivel         string `json:"nivel,omitempty" binding:"omitempty,oneof=INICIANTE INTERMEDIARIO AVANCADO"`
	CargaHoraria  int    `json:"carga_horaria,omitempty" binding:"omitempty,gt=0"`
	FocoPrincipal string `json:"foco_principal,omitempty" binding:"max=100"`
}

// --------------------------------------------------------------------------------
// DTOs de Resposta (Output)
// --------------------------------------------------------------------------------

// UsuarioResponse é o DTO de resposta para um usuário.
type UsuarioResponse struct {
	ID            int64     `json:"id"`
	Nome          string    `json:"nome"`
	Email         string    `json:"email"`
	AreaAtuacao   string    `json:"area_atuacao,omitempty"`
	NivelCarreira string    `json:"nivel_carreira,omitempty"`
	DataCadastro  time.Time `json:"data_cadastro"`
}

// TrilhaResponse é o DTO de resposta para uma trilha.
type TrilhaResponse struct {
	ID            int64  `json:"id"`
	Nome          string `json:"nome"`
	Descricao     string `json:"descricao,omitempty"`
	Nivel         string `json:"nivel"`
	CargaHoraria  int    `json:"carga_horaria"`
	FocoPrincipal string `json:"foco_principal,omitempty"`
}

// --------------------------------------------------------------------------------
// Exceções Customizadas (para tratamento de erros)
// --------------------------------------------------------------------------------

// CustomError é a interface para erros customizados.
type CustomError interface {
	Error() string
	StatusCode() int
	Message() string
}

// ResourceNotFoundError representa a exceção TrilhaNaoEncontradaException ou similar.
type ResourceNotFoundError struct {
	Resource string
	ID       int64
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%s não encontrado(a) com ID: %d", e.Resource, e.ID)
}

func (e *ResourceNotFoundError) StatusCode() int {
	return 404
}

func (e *ResourceNotFoundError) Message() string {
	return e.Resource + " não encontrado(a)."
}

// ConflictError representa um erro de conflito (ex: email já cadastrado).
type ConflictError struct {
	Msg string
}

func (e *ConflictError) Error() string {
	return e.Msg
}

func (e *ConflictError) StatusCode() int {
	return 409
}

func (e *ConflictError) Message() string {
	return e.Msg
}

// BusinessRuleError representa a exceção UsuarioNaoElegivelParaTrilhaException ou similar.
type BusinessRuleError struct {
	Msg string
}

func (e *BusinessRuleError) Error() string {
	return e.Msg
}

func (e *BusinessRuleError) StatusCode() int {
	return 422
}

func (e *BusinessRuleError) Message() string {
	return e.Msg
}
