package models

import (
	"time"
)

type WineStore struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CodigoLoja  string    `gorm:"size:512;not null;unique" json:"codigo_loja"`
	FaixaInicio uint64    `gorm:"default:0;unsigned" json:"faixa_inicio"`
	FaixaFim    uint64    `gorm:"default:0;unsigned" json:"faixa_fim"`
	CreatedAt   time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`
}

func (ws *WineStore) FillFieldsBeforeUpdate(wsDB *WineStore) {
	if ws.CodigoLoja == "" {
		ws.CodigoLoja = wsDB.CodigoLoja
	}
	if ws.FaixaInicio == 0 {
		ws.FaixaInicio = wsDB.FaixaInicio
	}
	if ws.FaixaFim == 0 {
		ws.FaixaFim = wsDB.FaixaFim
	}
}
