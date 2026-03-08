package strategies

import (
	"github.com/gin-gonic/gin"

	strategiesUC "trader-backend_monorepo/internal/usecases/strategies"
)

type Handlers interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type handlers struct {
	svc strategiesUC.Service
}

func NewHandlers(svc strategiesUC.Service) Handlers {
	result := &handlers{
		svc: svc,
	}

	return result
}

func (h *handlers) Create(ctx *gin.Context) {

}

func (h *handlers) GetByID(ctx *gin.Context) {

}
