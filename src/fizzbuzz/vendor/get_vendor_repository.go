package vendor

import (
	"context"

	"gorm.io/gorm"
)

func Get(db *gorm.DB) getVendorFunc {
	return func(context.Context) (VendorList, error) {
		var vendorList VendorList
		err := db.Find(&vendorList).Error
		return vendorList, err
	}
}

func GetByID(db *gorm.DB) func(context.Context, string) (Vendor, error) {
	return func(ctx context.Context, ID string) (Vendor, error) {
		var vendor Vendor
		err := db.First(&vendor, "id = ?", ID).Error
		return vendor, err
	}
}
