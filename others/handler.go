package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage *MemoryStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	h.storage.Insert(&employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	/* _ -> id */ _, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
	}
}
