package vendor

import (
	"context"
	"net/http"
	"pst-backend/domain/logger"

	"github.com/labstack/echo/v4"
)

type getVendorFunc func(context.Context) (VendorList, error)

func (fn getVendorFunc) Get(ctx context.Context) (VendorList, error) {
	return fn(ctx)
}

func GetHandler(svc getVendorFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Unwrap(c)

		vendor, err := svc.Get(c.Request().Context())

		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, vendor.ToResponse())
	}
}

type getVendorByIDFunc func(context.Context, string) (Vendor, error)

func (fn getVendorByIDFunc) GetByID(ctx context.Context, ID string) (Vendor, error) {
	return fn(ctx, ID)
}

func GetByIDHandler(svc getVendorByIDFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := c.Param("id")
		log := logger.Unwrap(c)

		vendors, err := svc.GetByID(c.Request().Context(), ID)
		if err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, vendors.ToResponse())
	}
}
