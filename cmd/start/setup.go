package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"upskilling-api/db"
	"upskilling-api/model"
)

// Dados iniciais para o seeder
var (
	usuarios = []model.Usuario{
		{Nome: "Ana Silva", Email: "ana.silva@exemplo.com", AreaAtuacao: "TI", NivelCarreira: "Pleno"},
		{Nome: "Bruno Costa", Email: "bruno.costa@exemplo.com", AreaAtuacao: "Finanças", NivelCarreira: "Em transição"},
		{Nome: "Carla Souza", Email: "carla.souza@exemplo.com", AreaAtuacao: "Marketing", NivelCarreira: "Junior"},
		{Nome: "Daniel Pereira", Email: "daniel.pereira@exemplo.com", AreaAtuacao: "Recursos Humanos", NivelCarreira: "Senior"},
	}

	trilhas = []model.Trilha{
		{Nome: "Inteligência Artificial para Negócios", Descricao: "Trilha focada em aplicação de IA em processos empresariais.", Nivel: "AVANCADO", CargaHoraria: 80, FocoPrincipal: "IA"},
		{Nome: "Análise de Dados com Python", Descricao: "Fundamentos e práticas de Data Science.", Nivel: "INTERMEDIARIO", CargaHoraria: 60, FocoPrincipal: "Dados"},
		{Nome: "Comunicação e Liderança Remota", Descricao: "Desenvolvimento de soft skills essenciais para o trabalho híbrido.", Nivel: "INICIANTE", CargaHoraria: 40, FocoPrincipal: "Soft Skills"},
	}

	competencias = []model.Competencia{
		{Nome: "Machine Learning", Categoria: "Tecnologia", Descricao: "Capacidade de desenvolver modelos de aprendizado de máquina."},
		{Nome: "Pensamento Crítico", Categoria: "Humana", Descricao: "Habilidade de analisar informações de forma objetiva."},
		{Nome: "Gestão de Projetos Ágeis", Categoria: "Gestão", Descricao: "Conhecimento em metodologias ágeis como Scrum e Kanban."},
	}
)

func main() {
	fmt.Println("Starting setup...")
	db.InitDB()
	defer db.CloseDB()

	createTables()
	seedTables()

	fmt.Println("Setup completed.")
}

// createTables executa o script SQL para criar as tabelas.
func createTables() {
	log.Println("Criando tabelas...")
	schemaSQL, err := os.ReadFile("/db/schema.sql")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de esquema SQL: %v", err)
	}

	_, err = db.GetDB().Exec(string(schemaSQL))
	if err != nil {
		log.Fatalf("Erro ao executar o script de criação de tabelas: %v", err)
	}
	log.Println("Tabelas criadas com sucesso.")
}

// seedTables popula as tabelas com dados iniciais.
func seedTables() {
	seedUsuarios()
	seedTrilhas()
	seedCompetencias()
	seedTrilhaCompetencia()
	seedMatriculas()
}

func seedUsuarios() {
	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM usuarios").Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar tabela de usuários: %v", err)
		return
	}
	if count > 0 {
		log.Println("Tabela de usuários já possui dados, pulando seeder.")
		return
	}

	log.Println("Populando tabela de usuários...")
	query := `INSERT INTO usuarios (nome, email, area_atuacao, nivel_carreira, data_cadastro) VALUES ($1, $2, $3, $4, $5);`
	for _, u := range usuarios {
		_, err := db.GetDB().Exec(query, u.Nome, u.Email, u.AreaAtuacao, u.NivelCarreira, time.Now())
		if err != nil {
			log.Printf("Erro ao popular usuário %s: %v", u.Nome, err)
		}
	}
	log.Println("Usuários populados com sucesso.")
}

func seedTrilhas() {
	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM trilhas").Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar tabela de trilhas: %v", err)
		return
	}
	if count > 0 {
		log.Println("Tabela de trilhas já possui dados, pulando seeder.")
		return
	}

	log.Println("Populando tabela de trilhas...")
	query := `INSERT INTO trilhas (nome, descricao, nivel, carga_horaria, foco_principal) VALUES ($1, $2, $3, $4, $5);`
	for _, t := range trilhas {
		_, err := db.GetDB().Exec(query, t.Nome, t.Descricao, t.Nivel, t.CargaHoraria, t.FocoPrincipal)
		if err != nil {
			log.Printf("Erro ao popular trilha %s: %v", t.Nome, err)
		}
	}
	log.Println("Trilhas populadas com sucesso.")
}

func seedCompetencias() {
	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM competencias").Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar tabela de competências: %v", err)
		return
	}
	if count > 0 {
		log.Println("Tabela de competências já possui dados, pulando seeder.")
		return
	}

	log.Println("Populando tabela de competências...")
	query := `INSERT INTO competencias (nome, categoria, descricao) VALUES ($1, $2, $3);`
	for _, c := range competencias {
		_, err := db.GetDB().Exec(query, c.Nome, c.Categoria, c.Descricao)
		if err != nil {
			log.Printf("Erro ao popular competência %s: %v", c.Nome, err)
		}
	}
	log.Println("Competências populadas com sucesso.")
}

func seedTrilhaCompetencia() {
	// Associa a primeira trilha (IA) com a primeira competência (Machine Learning)
	// Associa a segunda trilha (Dados) com a segunda competência (Pensamento Crítico)
	// Associa a terceira trilha (Soft Skills) com a terceira competência (Gestão de Projetos Ágeis)

	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM trilha_competencia").Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar tabela de trilha_competencia: %v", err)
		return
	}
	if count > 0 {
		log.Println("Tabela de trilha_competencia já possui dados, pulando seeder.")
		return
	}

	log.Println("Populando tabela de trilha_competencia...")
	query := `INSERT INTO trilha_competencia (trilha_id, competencia_id) VALUES ($1, $2);`

	// Busca IDs
	var trilhaID1, trilhaID2, trilhaID3 int64
	db.GetDB().QueryRow("SELECT id FROM trilhas WHERE nome = $1", trilhas[0].Nome).Scan(&trilhaID1)
	db.GetDB().QueryRow("SELECT id FROM trilhas WHERE nome = $1", trilhas[1].Nome).Scan(&trilhaID2)
	db.GetDB().QueryRow("SELECT id FROM trilhas WHERE nome = $1", trilhas[2].Nome).Scan(&trilhaID3)

	var competenciaID1, competenciaID2, competenciaID3 int64
	db.GetDB().QueryRow("SELECT id FROM competencias WHERE nome = $1", competencias[0].Nome).Scan(&competenciaID1)
	db.GetDB().QueryRow("SELECT id FROM competencias WHERE nome = $1", competencias[1].Nome).Scan(&competenciaID2)
	db.GetDB().QueryRow("SELECT id FROM competencias WHERE nome = $1", competencias[2].Nome).Scan(&competenciaID3)

	// Inserções
	if trilhaID1 > 0 && competenciaID1 > 0 {
		db.GetDB().Exec(query, trilhaID1, competenciaID1)
	}
	if trilhaID2 > 0 && competenciaID2 > 0 {
		db.GetDB().Exec(query, trilhaID2, competenciaID2)
	}
	if trilhaID3 > 0 && competenciaID3 > 0 {
		db.GetDB().Exec(query, trilhaID3, competenciaID3)
	}

	log.Println("Associações trilha-competência populadas com sucesso.")
}

func seedMatriculas() {
	var count int
	err := db.GetDB().QueryRow("SELECT COUNT(*) FROM matriculas").Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar tabela de matriculas: %v", err)
		return
	}
	if count > 0 {
		log.Println("Tabela de matriculas já possui dados, pulando seeder.")
		return
	}

	log.Println("Populando tabela de matriculas...")
	query := `INSERT INTO matriculas (usuario_id, trilha_id, data_inscricao, status) VALUES ($1, $2, $3, $4);`

	// Busca IDs
	var usuarioID1, usuarioID2 int64
	db.GetDB().QueryRow("SELECT id FROM usuarios WHERE email = $1", usuarios[0].Email).Scan(&usuarioID1)
	db.GetDB().QueryRow("SELECT id FROM usuarios WHERE email = $1", usuarios[1].Email).Scan(&usuarioID2)

	var trilhaID1, trilhaID2 int64
	db.GetDB().QueryRow("SELECT id FROM trilhas WHERE nome = $1", trilhas[0].Nome).Scan(&trilhaID1)
	db.GetDB().QueryRow("SELECT id FROM trilhas WHERE nome = $1", trilhas[1].Nome).Scan(&trilhaID2)

	// Ana matriculada em IA
	if usuarioID1 > 0 && trilhaID1 > 0 {
		db.GetDB().Exec(query, usuarioID1, trilhaID1, time.Now().AddDate(0, -1, 0), "ATIVA")
	}
	// Bruno matriculado em Dados
	if usuarioID2 > 0 && trilhaID2 > 0 {
		db.GetDB().Exec(query, usuarioID2, trilhaID2, time.Now().AddDate(0, 0, -5), "ATIVA")
	}

	log.Println("Matrículas populadas com sucesso.")
}
