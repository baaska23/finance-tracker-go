package subcategories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubcategoryHandler struct {
	service *SubcategoryService
}

func NewSubcategoryHandler(service *SubcategoryService) *SubcategoryHandler {
	return &SubcategoryHandler{service: service}
}

func (handler *SubcategoryHandler) Create(c *gin.Context) {
	var req SubCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	SubCategory, err := handler.service.CreateSubcategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, SubCategory)
}

func (handler *SubcategoryHandler) ListExpenseTypes(c *gin.Context) {
	SubCategorys, err := handler.service.ListExpenseTypes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, SubCategorys)
}

func (handler *SubcategoryHandler) ListIncomeTypes(c *gin.Context) {
	SubCategorys, err := handler.service.ListIncomeTypes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, SubCategorys)
}

func (handler *SubcategoryHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	SubCategory, err := handler.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, SubCategory)
}

func (handler *SubcategoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubCategory"})
		return
	}

	var req SubCategoryPatch
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	req.SubCategoryId = &id

	SubCategory, err := handler.service.UpdateSubcategory(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, SubCategory)
}

func (handler *SubcategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SubCategory id"})
		return
	}

	SubCategory, err := handler.service.DeleteSubcategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, SubCategory)
}
