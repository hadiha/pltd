// Code generated by candi v1.14.3.

package resthandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo"

	"PLTD3/internal/modules/business/domain"
	"PLTD3/pkg/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/wrapper"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(uc usecase.Usecase, deps dependency.Dependency) *RestHandler {
	return &RestHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root *echo.Group) {
	v1Root := root.Group(candihelper.V1)

	business := v1Root.Group("/business", echo.WrapMiddleware(h.mw.HTTPBearerAuth))
	business.GET("", h.getAllBusiness, echo.WrapMiddleware(h.mw.HTTPPermissionACL("getAllBusiness")))
	business.GET("/condition", h.getConditionUnit, echo.WrapMiddleware(h.mw.HTTPPermissionACL("getConditionUnit")))

	// Laporan
	business.GET("/report", h.getBusinessReport, echo.WrapMiddleware(h.mw.HTTPPermissionACL("getBusinessReport")))

	// hapus jika tidak perlu
	// business.GET("/:id", h.getDetailBusinessByID, echo.WrapMiddleware(h.mw.HTTPPermissionACL("getDetailBusiness")))
	// business.POST("", h.createBusiness, echo.WrapMiddleware(h.mw.HTTPPermissionACL("createBusiness")))
	// business.PUT("/:id", h.updateBusiness, echo.WrapMiddleware(h.mw.HTTPPermissionACL("updateBusiness")))
	// business.DELETE("/:id", h.deleteBusiness, echo.WrapMiddleware(h.mw.HTTPPermissionACL("deleteBusiness")))
}

func (h *RestHandler) getConditionUnit(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:GetConditionUnit")
	defer trace.Finish()

	var filter domain.FilterBusiness
	if err := candihelper.ParseFromQueryParam(c.Request().URL.Query(), &filter); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(c.Response())
	}

	// if err := h.validator.ValidateDocument("business/get_all", filter); err != nil {
	// 	return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(c.Response())
	// }

	data, err := h.uc.Business().GetConditionUnit(ctx, &filter)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(c.Response())
}


func (h *RestHandler) getBusinessReport(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:GetBusinessReport")
	defer trace.Finish()

	var filter domain.FilterBusiness
	if err := candihelper.ParseFromQueryParam(c.Request().URL.Query(), &filter); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(c.Response())
	}

	// if err := h.validator.ValidateDocument("business/get_all", filter); err != nil {
	// 	return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(c.Response())
	// }

	data, meta, err := h.uc.Business().GetBusinessReport(ctx, &filter)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success", meta, data).JSON(c.Response())
}

func (h *RestHandler) getAllBusiness(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:GetAllBusiness")
	defer trace.Finish()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using HTTPBearerAuth in middleware for this handler

	var filter domain.FilterBusiness
	if err := candihelper.ParseFromQueryParam(c.Request().URL.Query(), &filter); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(c.Response())
	}

	// ini dikomen dlu
	// if err := h.validator.ValidateDocument("business/get_all", filter); err != nil {
	// 	return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(c.Response())
	// }

	data, meta, err := h.uc.Business().GetAllBusiness(ctx, &filter)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	message := "Success, with your user id (" + tokenClaim.Subject + ") and role (" + tokenClaim.Role + ")"
	return wrapper.NewHTTPResponse(http.StatusOK, message, meta, data).JSON(c.Response())
}






// hapus kalo gak ada CRUD
func (h *RestHandler) getDetailBusinessByID(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:GetDetailBusinessByID")
	defer trace.Finish()

	data, err := h.uc.Business().GetDetailBusiness(ctx, c.Param("id"))
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(c.Response())
}

func (h *RestHandler) createBusiness(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:CreateBusiness")
	defer trace.Finish()

	body, _ := io.ReadAll(c.Request().Body)
	if err := h.validator.ValidateDocument("business/save", body); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(c.Response())
	}

	var payload domain.RequestBusiness
	if err := json.Unmarshal(body, &payload); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	err := h.uc.Business().CreateBusiness(ctx, &payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusCreated, "Success").JSON(c.Response())
}

func (h *RestHandler) updateBusiness(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:UpdateBusiness")
	defer trace.Finish()

	body, _ := io.ReadAll(c.Request().Body)
	if err := h.validator.ValidateDocument("business/save", body); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(c.Response())
	}

	var payload domain.RequestBusiness
	if err := json.Unmarshal(body, &payload); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	payload.ID = c.Param("id")
	err := h.uc.Business().UpdateBusiness(ctx, &payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(c.Response())
}

func (h *RestHandler) deleteBusiness(c echo.Context) error {
	trace, ctx := tracer.StartTraceWithContext(c.Request().Context(), "BusinessDeliveryREST:DeleteBusiness")
	defer trace.Finish()

	if err := h.uc.Business().DeleteBusiness(ctx, c.Param("id")); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(c.Response())
}
