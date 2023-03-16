CREATE TABLE IF NOT EXISTS wine_stores (
  id BIGINT PRIMARY KEY auto_increment,
  codigo_loja VARCHAR(255) NOT NULL unique,
  faixa_inicio INT(10) unsigned default 0,
  faixa_fim INT(10) unsigned default 0,
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP default CURRENT_TIMESTAMP
)
engine = InnoDB
default charset = utf8;