package purchasing_status_detail

import (
	"time"

	"gorm.io/gorm"
)

type PurchasingStatusDetailResponse struct {
	PurchaseRequisitionId     uint      `json:"purchaseRequisitionId"`
	ProjectId                 uint      `json:"projectId"`
	BidId                     string    `json:"bidId"`
	PurchaseRequisitionAmount float64   `json:"purchaseRequisitionAmount"`
	PurchaseOrderId           uint      `json:"purchaseOrderId"`
	CompanyName               string    `json:"companyName"`
	PurchaseOrderAmount       float64   `json:"purchaseOrderAmount"`
	PuchaseStatus             string    `json:"puchaseStatus"`
	Remark                    string    `json:"remark"`
	DisbursementDate          time.Time `json:"disbursementDate"`
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

// func MockDetail() []PurchasingStatusDetail {
// 	var purchasingDetailList []PurchasingStatusDetail
// 	purchasingCustomerDetail := PurchasingStatusDetail{
// 		WbsId:                     "I-64-A-ACTXX.6000",
// 		PurchaseRequisitionId:     1212354,
// 		ProjectId:                 123456,
// 		BidId:                     "BID-005",
// 		PurchaseRequisitionAmount: 50000.34,
// 		PurchaseOrderId:           1234634,
// 		CompanyName:               "Odd-e",
// 		PurchaseOrderAmount:       40398.50,
// 		PuchaseStatus:             "finish",
// 		Remark:                    "remark",
// 		DisbursementDate:          time.Now(),
// 	}
// 	purchasingInvestmentDetail := PurchasingStatusDetail{
// 		WbsId:                     "P-64-A-ACTXX.6000",
// 		PurchaseRequisitionId:     1,
// 		ProjectId:                 1,
// 		BidId:                     "1",
// 		PurchaseRequisitionAmount: 1.34,
// 		PurchaseOrderId:           1,
// 		CompanyName:               "Odd-e1",
// 		PurchaseOrderAmount:       1.50,
// 		PuchaseStatus:             "finish1",
// 		Remark:                    "remark1",
// 		DisbursementDate:          time.Now(),
// 	}
// 	purchasingProjectDetail := PurchasingStatusDetail{
// 		WbsId:                     "TC-64-C-ACTXX.6000",
// 		PurchaseRequisitionId:     2,
// 		ProjectId:                 3,
// 		BidId:                     "4",
// 		PurchaseRequisitionAmount: 5.34,
// 		PurchaseOrderId:           6,
// 		CompanyName:               "Odd-e2",
// 		PurchaseOrderAmount:       2.50,
// 		PuchaseStatus:             "finish2",
// 		Remark:                    "remark2",
// 		DisbursementDate:          time.Now(),
// 	}
// 	purchasingDetailList = append(purchasingDetailList, purchasingCustomerDetail, purchasingInvestmentDetail, purchasingProjectDetail)

// 	return purchasingDetailList
// }
