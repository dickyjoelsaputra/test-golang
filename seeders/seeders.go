package seeders

import (
	"fmt"
	"test-golang/config/mysql"
	"test-golang/models"
)

func Seed() {
	user := []models.User{
		{Name: "admin", Email: "email@gmail.com", Password: "123"},
	}

	products := []models.Product{
		{Name: "Product A", Description: "This is product A", Price: 1000, Image: "image_a.png"},
		{Name: "Product B", Description: "This is product B", Price: 2000, Image: "image_b.png"},
		{Name: "Product C", Description: "This is product C", Price: 3000, Image: "image_c.png"},
		{Name: "Product D", Description: "This is product D", Price: 4000, Image: "image_d.png"},
		{Name: "Product E", Description: "This is product E", Price: 5000, Image: "image_e.png"},
		{Name: "Product F", Description: "This is product F", Price: 6000, Image: "image_f.png"},
		{Name: "Product G", Description: "This is product G", Price: 7000, Image: "image_g.png"},
		{Name: "Product H", Description: "This is product H", Price: 8000, Image: "image_h.png"},
		{Name: "Product I", Description: "This is product I", Price: 9000, Image: "image_i.png"},
		{Name: "Product J", Description: "This is product J", Price: 10000, Image: "image_j.png"},
		{Name: "Product K", Description: "This is product K", Price: 11000, Image: "image_k.png"},
		{Name: "Product L", Description: "This is product L", Price: 12000, Image: "image_l.png"},
		{Name: "Product M", Description: "This is product M", Price: 13000, Image: "image_m.png"},
		{Name: "Product N", Description: "This is product N", Price: 14000, Image: "image_n.png"},
		{Name: "Product O", Description: "This is product O", Price: 15000, Image: "image_o.png"},
		{Name: "Product P", Description: "This is product P", Price: 16000, Image: "image_p.png"},
		{Name: "Product Q", Description: "This is product Q", Price: 17000, Image: "image_q.png"},
		{Name: "Product R", Description: "This is product R", Price: 18000, Image: "image_r.png"},
		{Name: "Product S", Description: "This is product S", Price: 19000, Image: "image_s.png"},
		{Name: "Product T", Description: "This is product T", Price: 20000, Image: "image_t.png"},
		{Name: "Product U", Description: "This is product U", Price: 21000, Image: "image_u.png"},
		{Name: "Product V", Description: "This is product V", Price: 22000, Image: "image_v.png"},
		{Name: "Product W", Description: "This is product W", Price: 23000, Image: "image_w.png"},
		{Name: "Product X", Description: "This is product X", Price: 24000, Image: "image_x.png"},
		{Name: "Product Y", Description: "This is product Y", Price: 25000, Image: "image_y.png"},
		{Name: "Product Z", Description: "This is product Z", Price: 26000, Image: "image_z.png"},
	}

	for i := 0; i < 3; i++ {
		products = append(products, products...)
	}

	for _, p := range products {
		err := mysql.DB.Create(&p).Error
		if err != nil {
			fmt.Println(err)
			panic("Failed to seed products")
		}
	}

	mysql.DB.Create(user)

	fmt.Println("Seed data success")
}
