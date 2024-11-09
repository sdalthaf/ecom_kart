package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"ecom_kart/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Cart)
// DB Table: Plural (carts)
// Resource: Plural (Carts)
// Path: Plural (/carts)
// View Template Folder: Plural (/templates/carts/)

// CartsResource is the resource for the Cart model
type CartsResource struct {
	buffalo.Resource
}

// List gets all Carts. This function is mapped to the path
// GET /carts
func (v CartsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	carts := &models.Carts{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Carts from the DB
	if err := q.All(carts); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("carts", carts)
		return c.Render(http.StatusOK, r.HTML("carts/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(carts))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(carts))
	}).Respond(c)
}

// Show gets the data for one Cart. This function is mapped to
// the path GET /carts/{cart_id}
func (v CartsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cart
	cart := &models.Cart{}

	// To find the Cart the parameter cart_id is used.
	if err := tx.Find(cart, c.Param("cart_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("cart", cart)

		return c.Render(http.StatusOK, r.HTML("carts/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(cart))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(cart))
	}).Respond(c)
}

// New renders the form for creating a new Cart.
// This function is mapped to the path GET /carts/new
func (v CartsResource) New(c buffalo.Context) error {
	c.Set("cart", &models.Cart{})

	return c.Render(http.StatusOK, r.HTML("carts/new.plush.html"))
}

// Create adds a Cart to the DB. This function is mapped to the
// path POST /carts
func (v CartsResource) Create(c buffalo.Context) error {
	// Allocate an empty Cart
	cart := &models.Cart{}

	// Bind cart to the html form elements
	if err := c.Bind(cart); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(cart)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("cart", cart)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("carts/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "cart.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/carts/%v", cart.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(cart))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(cart))
	}).Respond(c)
}

// Edit renders a edit form for a Cart. This function is
// mapped to the path GET /carts/{cart_id}/edit
func (v CartsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cart
	cart := &models.Cart{}

	if err := tx.Find(cart, c.Param("cart_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("cart", cart)
	return c.Render(http.StatusOK, r.HTML("carts/edit.plush.html"))
}

// Update changes a Cart in the DB. This function is mapped to
// the path PUT /carts/{cart_id}
func (v CartsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cart
	cart := &models.Cart{}

	if err := tx.Find(cart, c.Param("cart_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Cart to the html form elements
	if err := c.Bind(cart); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(cart)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("cart", cart)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("carts/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "cart.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/carts/%v", cart.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(cart))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(cart))
	}).Respond(c)
}

// Destroy deletes a Cart from the DB. This function is mapped
// to the path DELETE /carts/{cart_id}
func (v CartsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cart
	cart := &models.Cart{}

	// To find the Cart the parameter cart_id is used.
	if err := tx.Find(cart, c.Param("cart_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(cart); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "cart.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/carts")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(cart))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(cart))
	}).Respond(c)
}