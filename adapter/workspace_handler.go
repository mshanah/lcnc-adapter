package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mshanah/lcnc-domain/domain"
	"github.com/mshanah/lcnc-port/port"
)

type WorkspaceHandler struct {
	repo port.WorkspaceRepository
}

func NewWorkspaceHandler(repo port.WorkspaceRepository) *WorkspaceHandler {
	return &WorkspaceHandler{repo: repo}
}

func (h *WorkspaceHandler) CreateWorkspace(c *gin.Context) {
	var ws domain.Workspace
	if err := c.ShouldBindJSON(&ws); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.repo.Save(&ws)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ws)
}

func (h *WorkspaceHandler) GetWorkspaces(c *gin.Context) {
	workspaces, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workspaces)
}
