package fizzbuzz

import (
	"fizzbuzz/database"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestDB(t *testing.T) {
	var conn database.Database
	db := conn.GetConnectionDB()

	var cards []CreditCard
	cards = append(cards, CreditCard{Number: "Te3_1"})
	cards = append(cards, CreditCard{Number: "Te3_2"})
	var usert UserTest
	usert.Name = "Test2"
	usert.CreditCards = cards
	db.Create(&usert)
}

func TestGetUser(t *testing.T) {
	var conn database.Database
	db := conn.GetConnectionDB()

	var users []UserTest
	db.Preload("CreditCards").Find(&users)
	for _, us := range users {
		fmt.Println("xdf", us.Name)
		us.Name = "jinzhu 2"
		db.Save(&us)

		for _, card := range us.CreditCards {
			fmt.Println("-", card.Number)
		}
	}
}
func TestGetPurchasingStatus2(t *testing.T) {
	var conn database.Database
	db := conn.GetConnectionDB()

	var ps []PurchasingStatus
	db.Find(&ps)
	for _, us := range ps {
		fmt.Println("xdf", us.Wbs)
		us.BudgetName = "jinzhu 3"
		db.Save(&us)
	}
	fmt.Println("Endf")
}

type UserTest struct {
	gorm.Model
	Name         string
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber string
}
