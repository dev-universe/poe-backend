package http

import "github.com/gin-gonic/gin"

func NewRouter(h *ProofsHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/v1")
	{
		v1.POST("/proofs", h.CreateProof)
	}

	return r
}
