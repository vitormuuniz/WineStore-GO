package utils

import (
	"errors"

	"github.com/vitormuuniz/winestore-go/api/models"
	"github.com/vitormuuniz/winestore-go/api/repositories"
)

func ValidateFields(ws *models.WineStore, wineStoreRepository repositories.WineStoreRepository) error {
	if ws.FaixaInicio <= 0 || ws.FaixaFim <= 0 {
		return errors.New("faixaInicio and faixaFim must be non null and greather than 0")
	}

	if ws.CodigoLoja == "" {
		return errors.New("codigoLoja must be non null and not blank")
	}

	wineStores, err := wineStoreRepository.FindWineStoresFiltered(ws)
	if err != nil {
		return err
	}

	count, err := wineStoreRepository.Count()
	if err != nil {
		return err
	}

	if count > 0 && len(wineStores) > 0 {
		return errors.New("There is a zip range conflict, verify your data")
	}
	return nil
}

func ValidateFieldsRange(w *models.WineStore) error {
	if w.FaixaFim <= w.FaixaInicio {
		return errors.New("faixaFim must be greather than faixaInicio")
	}
	return nil
}
