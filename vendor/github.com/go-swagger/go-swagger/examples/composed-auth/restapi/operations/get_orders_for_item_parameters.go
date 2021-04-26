// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetOrdersForItemParams creates a new GetOrdersForItemParams object
//
// There are no default values defined in the spec.
func NewGetOrdersForItemParams() GetOrdersForItemParams {

	return GetOrdersForItemParams{}
}

// GetOrdersForItemParams contains all the bound params for the get orders for item operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetOrdersForItem
type GetOrdersForItemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ItemID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetOrdersForItemParams() beforehand.
func (o *GetOrdersForItemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rItemID, rhkItemID, _ := route.Params.GetOK("itemID")
	if err := o.bindItemID(rItemID, rhkItemID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindItemID binds and validates parameter ItemID from path.
func (o *GetOrdersForItemParams) bindItemID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ItemID = raw

	return nil
}
