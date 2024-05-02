package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/arif-rizal1122/go-online-shop/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func GetShoppingCartID(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionShoppingCart)
	if session.Values["cart_id"] == nil {
		session.Values["cart_id"] = uuid.New().String()
		_ = session.Save(r, w)
	}
    // ini hasil session yg dibuat
	return fmt.Sprintf("%v", session.Values["cart_id"])
}




func GetShoppingCart(db *gorm.DB, cartID string) (*models.Cart, error) {
	var cart models.Cart

	existCart, err := cart.GetCart(db, cartID)
	if err != nil {
		existCart, _ = cart.CreateCart(db, cartID)
	}
	fmt.Println(existCart)
	return existCart, nil
}




func (server *Server) GetCarts(w http.ResponseWriter, r *http.Request) {
    var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	cart, _ = GetShoppingCart(server.DB, cartID)
	fmt.Println("my cart id" , cart.ID)
}


func (server *Server) AddItemCart(w http.ResponseWriter, r *http.Request) {
	// diambil dari tmpl form product
	productID := r.FormValue("product_id")
	qty, _    := strconv.Atoi(r.FormValue("qty"))

	productModel := models.Product{}
	product, err := productModel.FindByID(server.DB, productID)

	if err != nil {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	} 
	if qty > product.Stock {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	var cart *models.Cart
	cartID  := GetShoppingCartID(w, r)
	cart, _  = GetShoppingCart(server.DB, cartID) 
	fmt.Println("cart id " ,cart.ID)
	http.Redirect(w, r, "/carts", http.StatusSeeOther)
} 