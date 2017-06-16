CREATE SEQUENCE equipes_cod_seq;
    
CREATE SEQUENCE partidas_cod_seq;
	
CREATE SEQUENCE jogadores_cod_seq;

CREATE TABLE IF NOT exists public.Jogadores (
                Cod_Jogador BIGINT NOT NULL DEFAULT NEXTVAL('jogadores_cod_seq'),
                Camisa TEXT NOT NULL,
                CONSTRAINT jogador_pk PRIMARY KEY (Cod_Jogador)
);


CREATE TABLE IF NOT exists public.Equipes (
                Cod_Equipe BIGINT NOT NULL DEFAULT NEXTVAL('equipes_cod_seq'),
                Nome TEXT NOT NULL,
                Estado VARCHAR(2),
                Pontuacao BIGINT NOT NULL DEFAULT 0,
                Gols_Pro BIGINT NOT NULL DEFAULT 0,
                Gols_Contra BIGINT NOT NULL DEFAULT 0,
		Partidas BIGINT NOT NULL DEFAULT 0,
                Vitorias BIGINT NOT NULL DEFAULT 0,
                Derrotas BIGINT NOT NULL DEFAULT 0,
                Empates BIGINT NOT NULL DEFAULT 0,
                Tecnico TEXT,
                CONSTRAINT equipes_pk PRIMARY KEY (Cod_Equipe)
);


CREATE TABLE IF NOT exists public.Partidas (
                Cod_Partida BIGINT NOT NULL DEFAULT NEXTVAL('partidas_cod_seq'),
                Local TEXT,
                Data TIMESTAMP,
                Cod_Equipe1 BIGINT NOT NULL,
                Gols_Equipe1 BIGINT NOT NULL DEFAULT 0,
                Cod_Equipe2 BIGINT NOT NULL,
                Gols_Equipe2 BIGINT NOT NULL DEFAULT 0,
                CONSTRAINT partidas_pk PRIMARY KEY (Cod_Partida)
);


CREATE TABLE  IF NOT exists public.Partidas_Jogadores (
                Cod_Partida BIGINT NOT NULL,
                Cod_Jogador BIGINT NOT NULL,
                Posicao TEXT NOT NULL,
                Gols BIGINT,
                Faltas BIGINT NOT NULL DEFAULT 0,
                Defesas BIGINT NOT NULL DEFAULT 0,
                CONSTRAINT partidas_jogadores_pk PRIMARY KEY (Cod_Partida, Cod_Jogador)
);


ALTER TABLE public.Partidas_Jogadores ADD CONSTRAINT jogador_partidas_jogadores_fk
FOREIGN KEY (Cod_Jogador)
REFERENCES public.Jogadores (Cod_Jogador)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE public.Partidas ADD CONSTRAINT equipes_partidas_fk
FOREIGN KEY (Cod_Equipe1)
REFERENCES public.Equipes (Cod_Equipe)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE public.Partidas ADD CONSTRAINT equipes_partidas_fk1
FOREIGN KEY (Cod_Equipe2)
REFERENCES public.Equipes (Cod_Equipe)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;

ALTER TABLE public.Partidas_Jogadores ADD CONSTRAINT partidas_partidas_jogadores_fk
FOREIGN KEY (Cod_Partida)
REFERENCES public.Partidas (Cod_Partida)
ON DELETE NO ACTION
ON UPDATE NO ACTION
NOT DEFERRABLE;