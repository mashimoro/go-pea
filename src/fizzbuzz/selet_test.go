package fizzbuzz

import (
	"fizzbuzz/database"
	"fmt"
	"testing"
)

func TestSeleGroup(t *testing.T) {
	selectPurchasingStatusDetail()
}
func selectPurchasingStatusDetail() error {
	var purchasingStatusDetails []PurchasingStatusDetail
	var conn database.Database
	db := conn.GetConnectionDB()

	var ps []PurchasingStatusDetail
	var ps1 PurchasingStatusDetail

	wbsId := "I-64-B-ACTXX.MS.6200"

	db.Select("purchasing_status_detail.wbs_id", "company_code",
		"sum(purchase_requisition_amount) as purchase_requisition_amount",
		"sum(purchase_order_amount) as purchase_order_amount", "vendor.name as company_name").Where("wbs_id = ?", wbsId).Group("wbs_id").Group("company_code").Joins("LEFT JOIN vendor on purchasing_status_detail.company_code=vendor.vendor_no").Group("name").Find(&ps)

	db.Select(
		"sum(purchase_requisition_amount) - sum(purchase_order_amount) as cost_saving", "wbs_id").Where("wbs_id = ?", wbsId).Group("wbs_id").First(&ps1)
	fmt.Println("ps-----", ps1)
	fmt.Println(len(ps))
	for _, item := range ps {
		fmt.Println("-1:", item.WbsId, item.PurchaseRequisitionId, "-2:", item.PurchaseRequisitionAmount, item.PurchaseOrderId, item.PurchaseOrderAmount, item.CompanyCode, item.CompanyName)
	}

	fmt.Println("Success delete purchasing_status_detail")
	for _, propurchasingStatusDetail := range purchasingStatusDetails {
		db.Create(&propurchasingStatusDetail)
	}
	return nil
}
