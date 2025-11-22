package analysis

import (
	"net/http"

	"github.com/avraam311/analysis-utility/internal/api/http/responses"
	"github.com/avraam311/analysis-utility/internal/infra/logger"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAnalysis(c *gin.Context) {
	an, err := h.service.GetAnalysis(c.Request.Context())
	if err != nil {
		logger.Logger.Error().Err(err).Msg("failed to get analysis")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, an)
}
