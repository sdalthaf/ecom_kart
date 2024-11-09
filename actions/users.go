package actions

import (
	"ecom_kart/models"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// UserRegisterGet displays a register form
func UsersRegisterGet(c buffalo.Context) error {
	// Make user available inside the html template
	return c.Render(http.StatusOK, r.HTML("users/register.html"))
}

func UsersRegisterPost(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := &models.User{}
	err := c.Bind(user)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	// Validate passwords match before hashing
	if user.Password != user.PasswordConfirm {
		c.Flash().Add("danger", "Passwords do not match.")
		return c.Render(422, r.HTML("users/register.html"))
	}

	// Hash the user's password
	if err := user.HashedPassword(); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndCreate(user)

	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if verrs.HasAny() {
		c.Set("errors", verrs.Errors)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("users/register.html"))
	}

	// Make user available inside the html template
	return c.Redirect(http.StatusSeeOther, ("/login"))
}

func LoginPage(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("/users/login.html"))
}

func LoginHandler(c buffalo.Context) error {

	user := &models.User{}
	err := c.Bind(user)

	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	// Query the user from the database by email
	tx := c.Value("tx").(*pop.Connection)

	if tx == nil {
		log.Println("No transaction found in context.")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	existingUser := &models.User{}
	log.Printf("Searching for user with email: %s", user.Email)

	lerr := tx.Where("email = ?", user.Email).First(existingUser)
	if lerr != nil {
		log.Printf("Error fetching user: %v", err) // Log the actual error
		c.Flash().Add("danger", "Invalid email or password")
		return c.Redirect(302, "/login")
	}

	if err := user.Athenticate(existingUser.PasswordHash); err != nil {
		c.Flash().Add("danger", "Invalid password")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	c.Session().Set("user_id", existingUser.ID)
	c.Session().Set("username", existingUser.Username)

	log.Printf("Session user_id: %v", c.Session().Get("user_id"))

	log.Printf("Does session have user_id? %v", c.Session().Session.Values)

	return c.Redirect(http.StatusSeeOther, "/")
}

func LogoutHandler(c buffalo.Context) error {
	// Clear the session
	c.Session().Clear()

	// Save the session
	err := c.Session().Save()
	if err != nil {
		return c.Error(500, err)
	}

	// Redirect to home after logout
	return c.Redirect(302, "/")
}
