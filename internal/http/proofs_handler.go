package http

import (
	"encoding/hex"
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

func (h *ProofsHandler) GetProof(c *gin.Context) {
	hashHex := c.Param("hash")

	// validate hex length: sha256 hex = 64
	if len(hashHex) != 64 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hash must be 64 hex chars (sha256)"})
		return
	}

	b, err := hex.DecodeString(hashHex)
	if err != nil || len(b) != 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hex hash"})
		return
	}

	var hash32 [32]byte
	copy(hash32[:], b)

	rec, err := h.Svc.GetRecord(c.Request.Context(), hash32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// timestamp=0이면 아직 기록 안 된 상태로 볼 수 있음
	if rec.Timestamp == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hash":      hashHex,
		"recorder":  rec.Recorder,
		"timestamp": rec.Timestamp,
	})
}
