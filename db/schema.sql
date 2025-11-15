-- Tabela de usuários da plataforma
CREATE TABLE IF NOT EXISTS usuarios (
    id BIGSERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    area_atuacao VARCHAR(100),
    nivel_carreira VARCHAR(50), -- exemplo: "Junior", "Pleno", "Senior", "Em transição"
    data_cadastro DATE NOT NULL DEFAULT CURRENT_DATE
);

-- Tabela de trilhas de aprendizagem
CREATE TABLE IF NOT EXISTS trilhas (
    id BIGSERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    descricao TEXT,
    nivel VARCHAR(50) NOT NULL, -- exemplo: "INICIANTE", "INTERMEDIARIO", "AVANCADO"
    carga_horaria INT NOT NULL, -- em horas
    foco_principal VARCHAR(100) -- ex: "IA", "Dados", "Soft Skills", "Green Tech"
);

-- Tabela de competências (skills)
CREATE TABLE IF NOT EXISTS competencias (
    id BIGSERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    categoria VARCHAR(100), -- ex: "Tecnologia", "Humana", "Gestão"
    descricao TEXT
);

-- Relação N:N entre trilhas e competências
CREATE TABLE IF NOT EXISTS trilha_competencia (
    trilha_id BIGINT NOT NULL,
    competencia_id BIGINT NOT NULL,
    PRIMARY KEY (trilha_id, competencia_id),
    CONSTRAINT fk_trilha_competencia_trilha
        FOREIGN KEY (trilha_id) REFERENCES trilhas (id) ON DELETE CASCADE,
    CONSTRAINT fk_trilha_competencia_competencia
        FOREIGN KEY (competencia_id) REFERENCES competencias (id) ON DELETE CASCADE
);

-- Matrícula de usuários em trilhas
CREATE TABLE IF NOT EXISTS matriculas (
    id BIGSERIAL PRIMARY KEY,
    usuario_id BIGINT NOT NULL,
    trilha_id BIGINT NOT NULL,
    data_inscricao DATE NOT NULL DEFAULT CURRENT_DATE,
    status VARCHAR(50) NOT NULL, -- ex: "ATIVA", "CONCLUIDA", "CANCELADA"
    CONSTRAINT fk_matricula_usuario
        FOREIGN KEY (usuario_id) REFERENCES usuarios (id) ON DELETE CASCADE,
    CONSTRAINT fk_matricula_trilha
        FOREIGN KEY (trilha_id) REFERENCES trilhas (id) ON DELETE CASCADE
);

-- Adicionando índice para busca rápida
CREATE INDEX IF NOT EXISTS idx_usuarios_email ON usuarios (email);
CREATE INDEX IF NOT EXISTS idx_trilhas_nivel ON trilhas (nivel);
CREATE INDEX IF NOT EXISTS idx_matriculas_usuario ON matriculas (usuario_id);
CREATE INDEX IF NOT EXISTS idx_matriculas_trilha ON matriculas (trilha_id);
