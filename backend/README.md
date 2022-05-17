# Customer
* Catalogue
```go
guest.Get("/get", h.GetProduct)  //done
guest.Get("/getall", h.GetAllProduct) //done

customer.Get("/buy", h.BuyProduct)  //done
```
* Auctions
```go
guest.Get("/get", h.GetProductAuction)  //done
guest.Get("/getall", h.GetAllProduct) //done
```
* All shops
```go
guest.Get("/get", h.GetOwner) //done
guest.Get("/getall", h.GetAllOwner) //done
```
* My orders
* My profile

# Shop
* Dashboard
* Catalogue and Auction
```go
owner.Post("/create", h.CreateProduct) //done
owner.Delete("/delete", h.DeleteProduct) //done
owner.Get("/get", h.GetProduct) //done
owner.Get("/getall", h.GetAllProduct) //done
owner.Put("/update", h.UpdateProduct) //done
```
* Orders
* Shop settings
```go
owner.Put("/update", h.UpdateOwner)
```