package actions

import (
	"ecom_kart/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func AuthorizeAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		userId := c.Session().Get("user_id")
		if userId == nil {
			c.Flash().Add("danger", "Please log in first")
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		tx := c.Value("tx").(*pop.Connection)
		user := &models.User{}

		if err := tx.Find(user, userId); err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}

		if !user.Admin {
			c.Flash().Add("danger", "Only admin is authorized to access this page")
			return c.Redirect(http.StatusSeeOther, "/")
		}

		return next(c)
	}
}

func AuthorizeUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		userId := c.Session().Get("user_id")
		if userId == nil {
			c.Flash().Add("danger", "Please login in!")
			c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}

func LoadCartAndWishlist(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		categories := &models.Categories{}

		if err := models.DB.Q().All(categories); err == nil {
			c.Set("categories", categories)
		}

		userId := c.Session().Get("user_id")
		if userId != nil {
			var wishlistItems []models.WishlistItem
			var cartItems []struct {
				models.CartItem
				Product models.Product `belongs_to:"product"`
			}

			// Fetch cart items with associated product details :=
			if err := models.DB.Q().
				InnerJoin("carts", "carts.id = cart_items.cart_id").
				InnerJoin("products", "products.id = cart_items.product_id").
				Where("carts.user_id = ?", userId).
				All(&cartItems); err == nil {
				c.Set("cartItems", cartItems)
			} else {
				fmt.Print(err)
			}
			// Fetch wishlist items
			if err := models.DB.Q().
				InnerJoin("wishlists", "wishlists.id = wishlist_items.wishlist_id").
				Where("wishlists.user_id = ?", userId).
				All(&wishlistItems); err == nil {
				c.Set("wishlistItems", wishlistItems)
			}
		}
		return next(c)
	}
}
