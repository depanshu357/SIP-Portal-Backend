package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"

	"github.com/gin-gonic/gin"
)

func GetProforma(c *gin.Context) {
	proformaId := c.Query("proformaId")
	var proforma models.JobDescription
	if err := database.DB.Where("id = ?", proformaId).First(&proforma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Proforma not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"proforma": proforma})
}
