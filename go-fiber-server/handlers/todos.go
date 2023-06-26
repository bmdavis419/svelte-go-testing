package handlers

import (
	"strconv"

	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/auth"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/storage"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Storage        *storage.TodoStorage
	SessionManager *auth.SessionManager
}

func NewTodoHandler(storage *storage.TodoStorage, session *auth.SessionManager) *TodoHandler {
	return &TodoHandler{Storage: storage, SessionManager: session}
}

type createTodoRequest struct {
	Title       string `json:"title" validate:"require,email"`
	Description string `json:"description" validate:"require,description"`
}

type createTodoResponse struct {
	Id int `json:"id"`
}

// CreateTodo godoc
// @Summary Create a new Todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body createTodoRequest true "The todo's info"
// @Security ApiKeyAuth
// @Success 200 {object} createTodoResponse
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// ensure the user is logged in
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]
	session, err := h.SessionManager.GetSession(sessionId)
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"message": "You are not logged in..."})
	}

	// get the request body
	var body createTodoRequest
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// create the todo
	id, err := h.Storage.CreateNewTodo(storage.NewTodoInput{
		UserId:      session.Id,
		Title:       body.Title,
		Description: body.Description,
	})
	if err != nil {
		return err
	}

	// send the id
	resp := createTodoResponse{Id: id}

	return c.JSON(resp)
}

type fetchOneTodoResponse struct {
	Todo storage.Todo_DB `json:"todo"`
}

// FetchTodo godoc
// @Summary Fetch one of a user's todos
// @Tags todos
// @Accept json
// @Produce json
// @Param			id	path		int	true	"Todo ID"
// @Security ApiKeyAuth
// @Success 200 {object} fetchOneTodoResponse
// @Router /todos/:id [get]
func (h *TodoHandler) FetchTodo(c *fiber.Ctx) error {
	// ensure the user is logged in
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]
	_, err := h.SessionManager.GetSession(sessionId)
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"message": "You are not logged in..."})
	}

	// get the id
	id := c.Params("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	// get the todo
	todo, err := h.Storage.GetOneTodo(aid)
	if err != nil {
		return err
	}

	// send
	resp := fetchOneTodoResponse{
		Todo: todo,
	}

	return c.JSON(resp)
}

type fetchTodosResponse struct {
	Todos []storage.Todo_DB `json:"todos"`
}

// FetchTodos godoc
// @Summary Fetch all of a user's todos
// @Tags todos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} fetchTodosResponse
// @Router /todos [get]
func (h *TodoHandler) FetchTodos(c *fiber.Ctx) error {
	// ensure the user is logged in
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]
	session, err := h.SessionManager.GetSession(sessionId)
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"message": "You are not logged in..."})
	}

	// get the todos
	todos, err := h.Storage.GetAllTodos(session.Id)
	if err != nil {
		return err
	}

	// send
	resp := fetchTodosResponse{
		Todos: todos,
	}
	return c.JSON(resp)
}

type basicResponse struct {
	Success bool `json:"success"`
}

// CompleteTodo godoc
// @Summary Mark a todo as completed
// @Tags todos
// @Accept json
// @Produce json
// @Param			id	path		int	true	"Todo ID"
// @Security ApiKeyAuth
// @Success 200 {object} basicResponse
// @Router /todos/:id/complete [put]
func (h *TodoHandler) CompleteTodo(c *fiber.Ctx) error {
	// ensure the user is logged in
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]
	_, err := h.SessionManager.GetSession(sessionId)
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"message": "You are not logged in..."})
	}

	// get the id
	id := c.Params("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	// complete todo
	err = h.Storage.CompleteTodo(aid)
	if err != nil {
		return err
	}

	// send
	resp := basicResponse{
		Success: true,
	}
	return c.JSON(resp)
}

// DeleteTodo godoc
// @Summary Delete a Todo
// @Tags todos
// @Accept json
// @Produce json
// @Param			id	path		int	true	"Todo ID"
// @Security ApiKeyAuth
// @Success 200 {object} basicResponse
// @Router /todos/:id [delete]
func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	// ensure the user is logged in
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]
	_, err := h.SessionManager.GetSession(sessionId)
	if err != nil {
		c.SendStatus(401)
		return c.JSON(fiber.Map{"message": "You are not logged in..."})
	}

	// get the id
	id := c.Params("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	// complete todo
	err = h.Storage.DeleteTodo(aid)
	if err != nil {
		return err
	}

	// send
	resp := basicResponse{
		Success: true,
	}
	return c.JSON(resp)
}
