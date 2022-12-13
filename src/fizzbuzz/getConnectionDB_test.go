package fizzbuzz

import (
	"fizzbuzz/database"
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestGetConnectionDB(t *testing.T) {
	var conn database.Database
	db := conn.GetConnectionDB()
	fmt.Println(db)
	createProductTable(db)
}

type PurchasingStatusDetail struct {
	gorm.Model
	ID                        uint `gorm:"primarykey"`
	WbsId                     string
	MaterialId                string
	Material                  string
	PurchaseRequisitionId     string
	ProjectId                 string
	BidId                     string
	PurchaseRequisitionAmount float64
	PurchaseOrderId           string
	CompanyName               string
	CompanyCode               string
	PurchaseOrderAmount       float64
	PurchaseStatus            string
	Remark                    string
	CostSaving                float64
	DisbursementDate          time.Time
	GrIrDoc                   string
	PmtDoc                    string
	PmtDate                   time.Time
}

func (PurchasingStatusDetail) TableName() string {
	return "purchasing_status_detail"
}
func createProductTable(db *gorm.DB) error {
	db.AutoMigrate(&PurchasingStatusDetail{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})
	// db.Create(&PurchasingStatusDetail{WbsId: "D42", PurchaseRequisitionAmount: 100})

	// db.Exec("DROP TABLE purchasing_status_detail")
	// db.Exec("DROP TABLE purchasing_status_detail")
	return nil
}
