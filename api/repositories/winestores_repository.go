package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vitormuuniz/winestore-go/api/models"
)

type WineStoreRepository interface {
	Save(*models.WineStore) (*models.WineStore, error)
	FindById(uint64) (*models.WineStore, error)
	FindAll() ([]*models.WineStore, error)
	FindWineStoresFiltered(*models.WineStore) ([]*models.WineStore, error)
	Update(*models.WineStore) error
	Delete(uint64) error
	Count() (int64, error)
}

type wineStoreRepositoryImpl struct {
	db *gorm.DB
}

func NewWineStoreRepository(db *gorm.DB) *wineStoreRepositoryImpl {
	return &wineStoreRepositoryImpl{db}
}

func (r *wineStoreRepositoryImpl) Save(wineStore *models.WineStore) (*models.WineStore, error) {
	tx := r.db.Begin()
	err := tx.Model(&models.WineStore{}).Create(wineStore).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return wineStore, tx.Commit().Error
}

func (r *wineStoreRepositoryImpl) FindById(id uint64) (*models.WineStore, error) {
	wineStore := &models.WineStore{}
	err := r.db.Debug().Model(&models.WineStore{}).Where("id = ?", id).Find(&wineStore).Error
	return wineStore, err
}

func (r *wineStoreRepositoryImpl) FindAll() ([]*models.WineStore, error) {
	wineStores := []*models.WineStore{}
	err := r.db.Debug().Model(&models.WineStore{}).Find(&wineStores).Error
	return wineStores, err
}

func (r *wineStoreRepositoryImpl) FindWineStoresFiltered(ws *models.WineStore) ([]*models.WineStore, error) {
	var wineStores []*models.WineStore
	err := r.db.Raw("SELECT * FROM wine_stores WHERE ID != ? AND ((? BETWEEN faixa_inicio AND faixa_fim) OR (? BETWEEN faixa_inicio AND faixa_fim) OR (? >= faixa_inicio AND ? <= faixa_fim) OR (? <= faixa_inicio AND ? >= faixa_fim))", ws.ID, ws.FaixaInicio, ws.FaixaFim, ws.FaixaInicio, ws.FaixaFim, ws.FaixaInicio, ws.FaixaFim).Scan(&wineStores).Error
	if err != nil {
		return nil, err
	}
	return wineStores, nil
}

func (r *wineStoreRepositoryImpl) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.WineStore{}).Count(&count).Error
	return count, err
}

func (r *wineStoreRepositoryImpl) Update(wineStore *models.WineStore) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"id":           wineStore.ID,
		"codigo_loja":  wineStore.CodigoLoja,
		"faixa_inicio": wineStore.FaixaInicio,
		"faixa_fim":    wineStore.FaixaFim,
		"updated_at":   time.Now(),
	}

	err := tx.Debug().Model(&models.WineStore{}).Where("id = ?", wineStore.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *wineStoreRepositoryImpl) Delete(id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.WineStore{}).Where("id = ?", id).Delete(&models.WineStore{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
