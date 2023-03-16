package routes

import (
	"net/http"

	"github.com/vitormuuniz/winestore-go/api/controllers"
)

type WineStoreRoutes interface {
	Routes() []*Route
}

type wineStoreRoutesImpl struct {
	wsController controllers.WineStoreController
}

func NewWineStoreRoutes(wsController controllers.WineStoreController) *wineStoreRoutesImpl {
	return &wineStoreRoutesImpl{wsController}
}

func (r *wineStoreRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/winestores",
			Method:  http.MethodPost,
			Handler: r.wsController.Create,
		},
		{
			Path:    "/winestores/{id}",
			Method:  http.MethodGet,
			Handler: r.wsController.FindById,
		},
		{
			Path:    "/winestores",
			Method:  http.MethodGet,
			Handler: r.wsController.FindAll,
		},
		{
			Path:    "/winestores/{id}",
			Method:  http.MethodPut,
			Handler: r.wsController.Update,
		},
		{
			Path:    "/winestores/{id}",
			Method:  http.MethodDelete,
			Handler: r.wsController.Delete,
		},
	}
}
