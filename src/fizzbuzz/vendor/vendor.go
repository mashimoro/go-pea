package vendor

import (
	"gorm.io/gorm"
)

type VendorResponse struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Name1                 string `json:"name1"`
	Name2                 string `json:"name2"`
	IndividualIncomeTaxNo string `json:"individualIncomeTaxNo"`
	CorporateIncomeTaxNo  string `json:"corporateIncomeTaxNo"`
	VatRegister           string `json:"vatRegister"`
	VendorNo              uint   `json:"vendorNo"`
	Status                string `json:"status"`
}

type Vendor struct {
	gorm.Model
	// todo change struct to business domain by data dictionary
	ID                    string
	Name                  string
	Name1                 string
	Name2                 string
	IndividualIncomeTaxNo string
	CorporateIncomeTaxNo  string
	VatRegister           string
	VendorNo              uint
	Status                string
}

func (vendor *Vendor) ToResponse() VendorResponse {
	return VendorResponse{
		ID:                    vendor.ID,
		Name:                  vendor.Name,
		Name1:                 vendor.Name1,
		Name2:                 vendor.Name2,
		IndividualIncomeTaxNo: vendor.IndividualIncomeTaxNo,
		CorporateIncomeTaxNo:  vendor.CorporateIncomeTaxNo,
		VatRegister:           vendor.VatRegister,
		VendorNo:              vendor.VendorNo,
		Status:                vendor.Status,
	}
}

type VendorList []Vendor

func (vendorList VendorList) ToResponse() []VendorResponse {
	var responses []VendorResponse
	for _, item := range vendorList {
		responses = append(responses, item.ToResponse())
	}
	return responses
}

func (Vendor) TableName() string {
	return "vendor"
}
