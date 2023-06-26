package handlers

import (
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
