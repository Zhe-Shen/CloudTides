// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListPolicyHandlerFunc turns a function with the right signature into a list policy handler
type ListPolicyHandlerFunc func(ListPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPolicyHandlerFunc) Handle(params ListPolicyParams) middleware.Responder {
	return fn(params)
}

// ListPolicyHandler interface for that can handle valid list policy params
type ListPolicyHandler interface {
	Handle(ListPolicyParams) middleware.Responder
}

// NewListPolicy creates a new http.Handler for the list policy operation
func NewListPolicy(ctx *middleware.Context, handler ListPolicyHandler) *ListPolicy {
	return &ListPolicy{Context: ctx, Handler: handler}
}

/*ListPolicy swagger:route GET /policy policy listPolicy

list all available policies

*/
type ListPolicy struct {
	Context *middleware.Context
	Handler ListPolicyHandler
}

func (o *ListPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListPolicyOKBody list policy o k body
//
// swagger:model ListPolicyOKBody
type ListPolicyOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// results
	Results []*ListPolicyOKBodyResultsItems0 `json:"results"`
}

// Validate validates this list policy o k body
func (o *ListPolicyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPolicyOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listPolicyOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPolicyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPolicyOKBody) UnmarshalBinary(b []byte) error {
	var res ListPolicyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ListPolicyOKBodyResultsItems0 list policy o k body results items0
//
// swagger:model ListPolicyOKBodyResultsItems0
type ListPolicyOKBodyResultsItems0 struct {

	// deploy type
	// Enum: [K8S VM]
	DeployType string `json:"deployType,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// idle policy
	IdlePolicy string `json:"idlePolicy,omitempty"`

	// is destroy
	IsDestroy bool `json:"isDestroy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// threshold policy
	ThresholdPolicy string `json:"thresholdPolicy,omitempty"`
}

// Validate validates this list policy o k body results items0
func (o *ListPolicyOKBodyResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeployType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listPolicyOKBodyResultsItems0TypeDeployTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["K8S","VM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listPolicyOKBodyResultsItems0TypeDeployTypePropEnum = append(listPolicyOKBodyResultsItems0TypeDeployTypePropEnum, v)
	}
}

const (

	// ListPolicyOKBodyResultsItems0DeployTypeK8S captures enum value "K8S"
	ListPolicyOKBodyResultsItems0DeployTypeK8S string = "K8S"

	// ListPolicyOKBodyResultsItems0DeployTypeVM captures enum value "VM"
	ListPolicyOKBodyResultsItems0DeployTypeVM string = "VM"
)

// prop value enum
func (o *ListPolicyOKBodyResultsItems0) validateDeployTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listPolicyOKBodyResultsItems0TypeDeployTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListPolicyOKBodyResultsItems0) validateDeployType(formats strfmt.Registry) error {

	if swag.IsZero(o.DeployType) { // not required
		return nil
	}

	// value enum
	if err := o.validateDeployTypeEnum("deployType", "body", o.DeployType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPolicyOKBodyResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPolicyOKBodyResultsItems0) UnmarshalBinary(b []byte) error {
	var res ListPolicyOKBodyResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
