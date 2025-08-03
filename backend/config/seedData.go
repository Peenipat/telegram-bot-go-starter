package config

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	model "github.com/Peenipat/telegram-bot-go-starter/model"
)

func ptrInt(v int) *int {
	return &v
}

func SeedData() {
	menuItems := []model.MenuItem{
		{Name: "ข้าวมันไก่", Category: "จานเดียว", Price: 50},
		{Name: "ก๋วยเตี๋ยวเรือ", Category: "เส้น", Price: 45},
		{Name: "ข้าวราดแกงเขียวหวาน", Category: "แกง", Price: 60},
		{Name: "ข้าวหมูแดง", Category: "จานเดียว", Price: 55},
		{Name: "ส้มตำ", Category: "ยำ", Price: 40},
		{Name: "ข้าวขาหมู", Category: "จานเดียว", Price: 65},
		{Name: "ลาบหมู", Category: "ยำ", Price: 60},
		{Name: "ต้มแซ่บกระดูกอ่อน", Category: "ต้ม", Price: 75},
		{Name: "หมูทอดกระเทียม", Category: "จานเดียว", Price: 55},
		{Name: "ข้าวผัดกุ้ง", Category: "จานเดียว", Price: 70},
	}
	for _, m := range menuItems {
		DB.FirstOrCreate(&m, model.MenuItem{Name: m.Name})
	}

	ingredients := []model.Ingredient{
		{Name: "หมู", Unit: "kg", StockLevel: 20},
		{Name: "ไก่", Unit: "kg", StockLevel: 15},
		{Name: "ข้าวสาร", Unit: "kg", StockLevel: 30},
		{Name: "ผักชี", Unit: "bunch", StockLevel: 10},
		{Name: "พริก", Unit: "g", StockLevel: 500},
		{Name: "กระเทียม", Unit: "g", StockLevel: 1000},
		{Name: "หอมแดง", Unit: "g", StockLevel: 800},
		{Name: "น้ำมันพืช", Unit: "bottle", StockLevel: 25},
		{Name: "เกลือ", Unit: "kg", StockLevel: 5},
		{Name: "น้ำตาล", Unit: "kg", StockLevel: 10},
	}
	for _, i := range ingredients {
		DB.FirstOrCreate(&i, model.Ingredient{Name: i.Name})
	}

	statuses := []string{"pending", "confirmed", "cancelled"}

	for i := 0; i < 10; i++ {
		reservation := model.Reservation{
			CustomerName: fmt.Sprintf("ลูกค้า %d", i+1),
			Phone:        fmt.Sprintf("08000000%d", i+1),
			TableNumber:  (i % 5) + 1,
			ReservedTime: time.Now().Add(time.Duration(i) * time.Hour),
			Status:       statuses[rand.Intn(len(statuses))],
		}

		DB.FirstOrCreate(&reservation, model.Reservation{
			CustomerName: reservation.CustomerName,
			Phone:        reservation.Phone,
		})
	}

	menuIDs := []uint{1, 2, 3, 4, 5}
	ingredientIDs := []uint{1, 2, 3, 4, 5}
	paymentMethods := []string{"cash", "credit", "mobile"}
	orderStatuses := []string{"completed", "pending", "cancelled"}
	for i := 0; i < 5; i++ {
		order := model.Order{
			TableNumber:   ptrInt((i % 5) + 1),
			TotalAmount:   0,
			PaymentMethod: paymentMethods[rand.Intn(len(paymentMethods))],
			Status:        orderStatuses[rand.Intn(len(orderStatuses))],
		}
		DB.Create(&order)

		var total float64
		numItems := rand.Intn(3) + 1 // 1–3 รายการ
		for j := 0; j < numItems; j++ {
			menuID := menuIDs[rand.Intn(len(menuIDs))]
			quantity := rand.Intn(3) + 1
			price := float64(40 + rand.Intn(81)) // ราคาสุ่ม 40–120

			orderItem := model.OrderItem{
				OrderID:    order.ID,
				MenuItemID: menuID,
				Quantity:   quantity,
				Price:      price * float64(quantity),
			}
			total += orderItem.Price
			DB.Create(&orderItem)
		}

		// อัปเดตราคารวม
		DB.Model(&order).Update("total_amount", total)

		// ใส่ feedback บางรายการ
		if rand.Intn(2) == 0 {
			feedback := model.Feedback{
				OrderID:      &order.ID,
				CustomerName: fmt.Sprintf("ลูกค้า %d", i+1),
				Comment:      "อาหารอร่อย บริการดี",
				Rating:       rand.Intn(3) + 3, // 3-5
			}
			DB.Create(&feedback)
		}
	}

	// Seed stock movements
	for i := 0; i < 5; i++ {
		mov := model.StockMovement{
			IngredientID: ingredientIDs[rand.Intn(len(ingredientIDs))],
			Amount:       float64(rand.Intn(5)+1) * -1, // ออกสต็อก
			Type:         "out",
			CreatedAt:    time.Now().Add(-time.Duration(i) * time.Hour),
		}
		DB.Create(&mov)
	}

	log.Println("Seed data completed")
}
