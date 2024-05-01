package controllers

import (
	"net/http"

	"github.com/arif-rizal1122/go-online-shop/app/models"
	"github.com/unrolled/render"
)



func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
     render := render.New(render.Options{
		Layout: "layout",
	 })
	 
	 productModel := models.Product{}
	 products, err := productModel.GetProducts(server.DB)
     if err != nil {
		return
	 }

	 _ = render.HTML(w, http.StatusOK, "products", map[string]interface{} {
		"products": products,
	 })

}