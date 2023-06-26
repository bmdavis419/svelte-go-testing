package handlers

import (
	"fmt"

	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/auth"
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/storage"
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

type SuccessResponse struct {
	Success bool `json:"success"`
}

// SignOutUser godoc
//
//	@Summary	Sign out a user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	SuccessResponse
//	@Router		/users/sign-out [post]
func (u *UserHandler) SignOutUser(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]

	// delete the session
	err := u.SessionManager.SignOut(sessionId)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

// GetUserInfo godoc
// @Summary Get the user's info
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} auth.UserSession
// @Router /users/me [get]
func (u *UserHandler) GetUserInfo(c *fiber.Ctx) error {
	// get the session from the authorization header
	sessionHeader := c.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]

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

// SignInUser godoc
// @Summary Sign in a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body signInRequestBody true "The user's email and password"
// @Success 200 {object} SuccessResponse
// @Header 200 {string} Authorization "contains the session id in bearer format"
// @Router /users/sign-in [post]
func (u *UserHandler) SignInUser(c *fiber.Ctx) error {
	var user signInRequestBody

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	fmt.Println(user)

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

	// set the session id as a header
	c.Response().Header.Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))

	return c.JSON(fiber.Map{"success": true})
}

type userRequestBody struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type signUpSuccessResponse struct {
	Id int `json:"id"`
}

// SignUpUser godoc
// @Summary Sign up a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body userRequestBody true "The user's first name, last name, email, and password"
// @Success 200 {object} signUpSuccessResponse
// @Header 200 {string} Authorization "contains the session id in bearer format"
// @Router /users/sign-up [post]
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

	// set the session id as a header
	c.Response().Header.Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))

	resp := signUpSuccessResponse{
		Id: id,
	}
	return c.JSON(resp)
}
