package http

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dev-universe/poe-backend/internal/service"
)

type ProofsHandler struct {
	Svc       *service.PoEService
	MaxUpload int64
}

func NewProofsHandler(svc *service.PoEService, maxUpload int64) *ProofsHandler {
	return &ProofsHandler{Svc: svc, MaxUpload: maxUpload}
}

func (h *ProofsHandler) CreateProof(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.MaxUpload)

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file (form field: file)"})
		return
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read uploaded file"})
		return
	}

	res, err := h.Svc.RecordBytes(c.Request.Context(), b)
	if err != nil {
		if errors.Is(err, service.ErrAlreadyRecorded) {
			c.JSON(http.StatusConflict, res)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
