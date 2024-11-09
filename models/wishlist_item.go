package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// WishlistItem is used by pop to map your wishlist_items database table to your go code.
type WishlistItem struct {
	ID         uuid.UUID `json:"id" db:"id"`
	WishlistID uuid.UUID `json:"wishlist_id" db:"wishlist_id"`
	ProductID  uuid.UUID `json:"product_id" db:"product_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (w WishlistItem) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// WishlistItems is not required by pop and may be deleted
type WishlistItems []WishlistItem

// String is not required by pop and may be deleted
func (w WishlistItems) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *WishlistItem) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *WishlistItem) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *WishlistItem) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
