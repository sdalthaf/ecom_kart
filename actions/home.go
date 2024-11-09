package actions

import (
	"ecom_kart/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.html"))
}

func ShopHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/shop.html"))
}

func DetailHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/detail.html"))
}

func ContactHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/contact.html"))
}

func CheckoutHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/checkout.html"))
}

func CartHandler(c buffalo.Context) error {

	userId := c.Session().Get("user_id")

	if userId != nil {
		var currentCartItems []struct {
			models.CartItem
			Product models.Product `belongs_to:"product"`
		}

		// Fetch cart items with associated product details
		if err := models.DB.Q().
			InnerJoin("carts", "carts.id = cart_items.cart_id").
			InnerJoin("products", "products.id = cart_items.product_id").
			Where("carts.user_id = ?", userId).
			All(&currentCartItems); err != nil {
			return c.Error(500, err)
		}

		// Set cart items in context to pass them to the view
		c.Set("cartItems", currentCartItems)
	}
	return c.Render(http.StatusOK, r.HTML("home/cart.html"))
}

func AccountHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("/users/account.html"))
}
