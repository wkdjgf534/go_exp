package strategies

import (
	"net/http"

	"github.com/gin-gonic/gin"

	strategiesUC "trader-backend_monorepo/internal/usecases/strategies"
	"trader-backend_monorepo/pkg/apierrors"
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
	var request CreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		apiErr := apierrors.NewBadRequestError("invalid json body")
		ctx.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	svcReq := &strategiesUC.CreateStrategyRequest{
		Name: request.Name,
	}

	svcResponse, err := h.svc.Create(ctx, svcReq)
	if err != nil {
		apiErr := apierrors.FromError(err)
		ctx.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	response := &CreateResponse{
		StrategyID: svcResponse.StrategyID,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *handlers) GetByID(ctx *gin.Context) {

}
