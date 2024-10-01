package handler

import (
	"github.com/gin-gonic/gin"
	request2 "go-microservices/cmd/to_do/internal/api/handler/request"
	"go-microservices/cmd/to_do/internal/service"
	"go-microservices/cmd/to_do/pkg/model"
	"go-microservices/internal/midleware"
	"net/http"
	"strconv"
)

type TodoItemHandler struct {
	TodoItemService service.TodoItemServiceInterface
}

func NewTodoItemHandler(todoItemService service.TodoItemServiceInterface) *TodoItemHandler {
	return &TodoItemHandler{TodoItemService: todoItemService}
}

func (h *TodoItemHandler) GetTodoItemById(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	value, oke := c.Get(midleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		c.Abort()
		return
	}
	userPrincipal := value.(*midleware.UserPrincipal)

	item, err := h.TodoItemService.FindTodoItemById(uint(id), userPrincipal.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})

}

func (h *TodoItemHandler) GetListTodoItem(c *gin.Context) {

	value, oke := c.Get(midleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		return
	}

	userPrincipal := value.(*midleware.UserPrincipal)

	items, err := h.TodoItemService.FindAllTodoItem(userPrincipal.Id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})

}

func (h *TodoItemHandler) CreateTodoItem(c *gin.Context) {

	var request request2.TodoItemCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, oke := c.Get(midleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		return
	}
	userPrincipal := value.(*midleware.UserPrincipal)

	err := h.TodoItemService.CreateTodoItem(&model.TodoItem{
		Name:   request.Name,
		State:  request.State,
		UserId: userPrincipal.Id,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})

}

func (h *TodoItemHandler) DeleteTodoItem(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	value, oke := c.Get(midleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		c.Abort()
		return
	}
	userPrincipal := value.(*midleware.UserPrincipal)

	err = h.TodoItemService.DeleteTodoItem(uint(id), userPrincipal.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})

}
func (h *TodoItemHandler) UpdateTodoItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var request request2.TodoItemUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value, oke := c.Get(midleware.USER_PRICIPAL_CONTEXT_KEY)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "internal server error",
		})
		return
	}
	userPrincipal := value.(*midleware.UserPrincipal)

	item := &model.TodoItem{
		Name:  request.Name,
		State: request.State,
	}
	err = h.TodoItemService.UpdateTodoItem(uint(id), userPrincipal.Id, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})

}
