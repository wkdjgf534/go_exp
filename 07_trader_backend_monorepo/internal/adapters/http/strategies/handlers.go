package strategies

import (
	"net/http"
	"strings"

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
		Name:        request.Name,
		Description: request.Description,
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
	request := &strategiesUC.GetStrategyByIDRequest{
		StrategyID: strings.TrimSpace(ctx.Param("strategy_id")),
	}

	svcResponse, err := h.svc.GetByID(ctx, request)
	if err != nil {
		apiErr := apierrors.FromError(err)
		ctx.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	response := &GetByIDResponse{
		Strategy: *fromStrategyCoreToHTTP(svcResponse.Strategy),
	}

	ctx.JSON(http.StatusOK, response)
}
