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
// Model: Singular (Wishlist)
// DB Table: Plural (wishlists)
// Resource: Plural (Wishlists)
// Path: Plural (/wishlists)
// View Template Folder: Plural (/templates/wishlists/)

// WishlistsResource is the resource for the Wishlist model
type WishlistsResource struct {
	buffalo.Resource
}

// List gets all Wishlists. This function is mapped to the path
// GET /wishlists
func (v WishlistsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	wishlists := &models.Wishlists{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Wishlists from the DB
	if err := q.All(wishlists); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("wishlists", wishlists)
		return c.Render(http.StatusOK, r.HTML("wishlists/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(wishlists))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(wishlists))
	}).Respond(c)
}

// Show gets the data for one Wishlist. This function is mapped to
// the path GET /wishlists/{wishlist_id}
func (v WishlistsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Wishlist
	wishlist := &models.Wishlist{}

	// To find the Wishlist the parameter wishlist_id is used.
	if err := tx.Find(wishlist, c.Param("wishlist_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("wishlist", wishlist)

		return c.Render(http.StatusOK, r.HTML("wishlists/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(wishlist))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(wishlist))
	}).Respond(c)
}

// New renders the form for creating a new Wishlist.
// This function is mapped to the path GET /wishlists/new
func (v WishlistsResource) New(c buffalo.Context) error {
	c.Set("wishlist", &models.Wishlist{})

	return c.Render(http.StatusOK, r.HTML("wishlists/new.plush.html"))
}

// Create adds a Wishlist to the DB. This function is mapped to the
// path POST /wishlists
func (v WishlistsResource) Create(c buffalo.Context) error {
	// Allocate an empty Wishlist
	wishlist := &models.Wishlist{}

	// Bind wishlist to the html form elements
	if err := c.Bind(wishlist); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(wishlist)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("wishlist", wishlist)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("wishlists/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "wishlist.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/wishlists/%v", wishlist.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(wishlist))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(wishlist))
	}).Respond(c)
}

// Edit renders a edit form for a Wishlist. This function is
// mapped to the path GET /wishlists/{wishlist_id}/edit
func (v WishlistsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Wishlist
	wishlist := &models.Wishlist{}

	if err := tx.Find(wishlist, c.Param("wishlist_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("wishlist", wishlist)
	return c.Render(http.StatusOK, r.HTML("wishlists/edit.plush.html"))
}

// Update changes a Wishlist in the DB. This function is mapped to
// the path PUT /wishlists/{wishlist_id}
func (v WishlistsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Wishlist
	wishlist := &models.Wishlist{}

	if err := tx.Find(wishlist, c.Param("wishlist_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Wishlist to the html form elements
	if err := c.Bind(wishlist); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(wishlist)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("wishlist", wishlist)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("wishlists/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "wishlist.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/wishlists/%v", wishlist.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(wishlist))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(wishlist))
	}).Respond(c)
}

// Destroy deletes a Wishlist from the DB. This function is mapped
// to the path DELETE /wishlists/{wishlist_id}
func (v WishlistsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Wishlist
	wishlist := &models.Wishlist{}

	// To find the Wishlist the parameter wishlist_id is used.
	if err := tx.Find(wishlist, c.Param("wishlist_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(wishlist); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "wishlist.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/wishlists")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(wishlist))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(wishlist))
	}).Respond(c)
}
