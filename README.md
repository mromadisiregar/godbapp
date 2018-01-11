## Godbapp

With golang and mongo image, this compose will build dummy counters and simple rest api service for storing user data. Next, i'll find a way to save shopping cart data. Add and remove item from nested array.

```
type Users struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nama string `bson:"nama" json:"nama"`
	Kota string `bson:"kota" json:"kota"`
	//Keranjang []Keranjang `bson:"keranjang" json:"keranjang"`
}
```

Http method implementation : 
```
r.HandleFunc("/api/counter", api.CheckCounter).Methods("GET")
r.HandleFunc("/api/users", api.CreateUser).Methods("POST")
r.HandleFunc("/api/users/{id}", api.GetUser).Methods("GET")
r.HandleFunc("/api/users/{id}", api.DeleteUser).Methods("DELETE")
r.HandleFunc("/api/users", api.GetAllUser).Methods("GET")
```
---

**How to run this code?**
```
~ docker-compose build
~ docker-compose up -d
```

**Check the counter**

![alt text](https://i.imgur.com/ZM8UbxG.png)

**Create new user**

![alt text](https://i.imgur.com/qFgzNvN.png)

**Get user by id**

![alt text](https://i.imgur.com/wn0FLc9.png)

**Delete user**

![alt text](https://i.imgur.com/vbchHjg.png)