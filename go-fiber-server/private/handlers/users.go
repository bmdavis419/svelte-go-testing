package handlers

import (
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/private/auth"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/private/storage"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Storage        *storage.UserStorage
	SessionManager *auth.SessionManager
}

func NewUserHandler(storage *storage.UserStorage, sessionManager *auth.SessionManager) *UserHandler {
	return &UserHandler{Storage: storage, SessionManager: sessionManager}
}

func (u *UserHandler) GenerateNewUser(c *fiber.Ctx) error {
	id, err := u.Storage.GenerateNewUser()
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"id": id})
}

func (u *UserHandler) SignOutUser(c *fiber.Ctx) error {
	// get the user id from the session
	sessionId := c.Cookies("session_id")
	if sessionId == "" {
		return c.JSON(fiber.Map{"error": "no session id found"})
	}

	// delete the session
	err := u.SessionManager.SignOut(sessionId)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	// delete the cookie
	c.Cookie(&fiber.Cookie{
		Name:   "session_id",
		MaxAge: -1,
	})

	return c.JSON(fiber.Map{"success": true})
}

func (u *UserHandler) GetUserInfo(c *fiber.Ctx) error {
	// get the user id from the session
	sessionId := c.Cookies("session_id")
	if sessionId == "" {
		return c.JSON(fiber.Map{"error": "no session id found"})
	}

	// get the user data from the session
	user, err := u.SessionManager.GetSession(sessionId)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

type signInRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *UserHandler) SignInUser(c *fiber.Ctx) error {
	var user signInRequestBody

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	// validate the user struct
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return err
	}

	// sign the user in
	sessionId, err := u.SessionManager.SignIn(user.Email, user.Password)
	if err != nil {
		return err
	}

	// return the session id
	c.Cookie(&fiber.Cookie{
		Name:  "session_id",
		Value: sessionId,
		// HTTPOnly: true,
		// Secure:   true,
	})

	return c.JSON(fiber.Map{"success": true})
}

type userRequestBody struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

func (u *UserHandler) SignUpUser(c *fiber.Ctx) error {
	// get the info from the request body
	var user userRequestBody

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	// validate the user struct
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return err
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// create the user
	id, err := u.Storage.CreateNewUser(storage.NewUser{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  string(hashedPassword),
	})

	if err != nil {
		return err
	}

	// generate the user's session
	sessionId, err := u.SessionManager.GenerateSession(auth.UserSession{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		return err
	}

	// return the session id
	c.Cookie(&fiber.Cookie{
		Name:  "session_id",
		Value: sessionId,
		// HTTPOnly: true,
		// Secure:   true,
	})

	return c.JSON(fiber.Map{"id": id})
}
