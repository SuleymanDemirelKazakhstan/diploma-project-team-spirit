# Admin
```go
// Admin
basicauth{"admin": "admin"}
//Create Shop Owner
Post("/admin/create", {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    Phone    string `json:"phone" validate:"required"`
    Address  string `json:"address" validate:"required"`
})
// Delete Shop Owner
Delete("/admin/delete", {
    Id int // shop id
})
// Get Shop Owner
Get("/admin/get", {
    Id int // shop id
}) return{
    "status":  true,
    "message": "success",
    "user":    {
        Id       int    `json:"id"`
        Name     string `json:"name" validate:"required"`
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required"`
        Phone    string `json:"phone" validate:"required"`
        Address  string `json:"address" validate:"required"`
        Image    string `json:"image"       form:"image"`
    },
}
// Get all Shop Owner
Get("/admin/getall")
return{
    "status":  true,
    "message": "success",
    "users": [
        {
            Id       int    `json:"id"`
            Name     string `json:"name" validate:"required"`
            Email    string `json:"email" validate:"required,email"`
            Password string `json:"password" validate:"required"`
            Phone    string `json:"phone" validate:"required"`
            Address  string `json:"address" validate:"required"`
            Image    string `json:"image"       form:"image"`
        },
    ],
}
// Update Shop Owner
Put("/admin/update", {
    Id       int
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    Phone    string `json:"phone" validate:"required"`
    Address  string `json:"address" validate:"required"`
})
// Save Image
Post("/admin/saveimage", {
    Id int // shop id + image
})
// Delete Image
Delete("/admin/deleteimage", {
    Name string // file name
    Id   int    // shop id
})
```

# Shop
```go
// Owner JWT token
Post("/auth/owner", {
    Email    string `json:"email,omitempty" validate:"required,email"`
    Password string `json:"password,omitempty" validate:"required"`
})return{
    "status": "success",
    "token":  t,
}
owner.Post("/owner/create", {
    OwnerId     int       `json:"shop_id,omitempty" validate:"required"`
    Price       float64   `json:"price,omitempty" validate:"required"`
    Name        string    `json:"name,omitempty" validate:"required"`
})
owner.Delete("/owner/delete", {
    Id int
})
owner.Get("/owner/get", {
    Id int
})return{
    "status":  true,
    "message": "success",
    "product": product,
}
owner.Get("/owner/getall")
return{
		"status":   true,
		"message":  "success",
		"products": products,
	}
owner.Put("/owner/update", {
    Id          int
    Price       float64   `json:"price,omitempty" validate:"required"`
    Name        string    `json:"name,omitempty" validate:"required"`
    Description string    `json:"description,omitempty"`
    Discount    int       `json:"discount,omitempty"`
    Auction     bool
})
owner.Post("/owner/saveimage", {
    Id int // shop id + image
})
owner.Delete("/owner/deleteimage", {
    Name string // file name
    Id   int    // shop id
})
owner.Get("/owner/order", {
    Id int // shop_id
})return{
		"status":   true,
		"message":  "success",
		"products": products,
	}
owner.Put("/owner/issued", {
    Id int
})
owner.Get("/owner/verify", {
    Email string
})return{
		"status":  true,
		"message": "success",
		"code":    code,
	}
owner.Put("/owner/updateinfo", {
    Id       int
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    Phone    string `json:"phone" validate:"required"`
    Address  string `json:"address" validate:"required"`
})
```

# Customer
```go
get "/" return {
    "status": true, 
    "message": "success"
}
// JWT
auth.Post("/auth/customer", h.Customer.Login)

// Unauthorized customer
guest.Post("/g/signup", {
    Name     string `json:"name,omitempty" validate:"required"`
    Email    string `json:"email,omitempty" validate:"required,email"`
    Password string `json:"password,omitempty" validate:"required"`
    Phone    string `json:"phone,omitempty" validate:"required"`
})
guest.Get("/g/get", {
    Id int
})return{
    "status":  true,
    "message": "success",
    "product": product,
}
guest.Get("/g/allproduct", 
return{
		"status":   true,
		"message":  "success",
		"products": products,
	})
guest.Get("/g/verify", {
    Email string
})return{
    "status":  true,
    "message": "success",
    "code":    code,
}

// Authorized customer
customer.Post("/c/buy", {
    Customer_id int 
    Product_id  int     
    Shop_id     int 
})
// return auction goods price
customer.Get("/c/getter", {
    Id string
})return{
		"status":  true,
		"message": "success",
		"Value":   v,
	}
// Set price for auction goods
customer.Post("/c/setter", h.Customer.Setter)
customer.Post("/c/saveimage", {
    Id int // shop id + image
})
customer.Delete("/c/deleteimage", {
    Name string // file name
    Id   int    // shop id
})
customer.Get("/c/order", {
    Id int // customer_id
})return{
		"status":   true,
		"message":  "success",
		"products": products,
	}
customer.Get("/c/verify", {
    Email string
})return{
    "status":  true,
    "message": "success",
    "code":    code,
}
```