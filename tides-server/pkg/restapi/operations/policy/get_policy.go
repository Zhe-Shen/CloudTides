// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPolicyHandlerFunc turns a function with the right signature into a get policy handler
type GetPolicyHandlerFunc func(GetPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPolicyHandlerFunc) Handle(params GetPolicyParams) middleware.Responder {
	return fn(params)
}

// GetPolicyHandler interface for that can handle valid get policy params
type GetPolicyHandler interface {
	Handle(GetPolicyParams) middleware.Responder
}

// NewGetPolicy creates a new http.Handler for the get policy operation
func NewGetPolicy(ctx *middleware.Context, handler GetPolicyHandler) *GetPolicy {
	return &GetPolicy{Context: ctx, Handler: handler}
}

/*GetPolicy swagger:route GET /policy/{id} policy getPolicy

get info of a policy

*/
type GetPolicy struct {
	Context *middleware.Context
	Handler GetPolicyHandler
}

func (o *GetPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetPolicyNotFoundBody get policy not found body
//
// swagger:model GetPolicyNotFoundBody
type GetPolicyNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get policy not found body
func (o *GetPolicyNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPolicyNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPolicyNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetPolicyNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetPolicyOKBody get policy o k body
//
// swagger:model GetPolicyOKBody
type GetPolicyOKBody struct {

	// deploy type
	// Enum: [K8S VM]
	DeployType string `json:"deployType,omitempty"`

	// hosts assigned
	HostsAssigned int64 `json:"hostsAssigned,omitempty"`

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

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this get policy o k body
func (o *GetPolicyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeployType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getPolicyOKBodyTypeDeployTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["K8S","VM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getPolicyOKBodyTypeDeployTypePropEnum = append(getPolicyOKBodyTypeDeployTypePropEnum, v)
	}
}

const (

	// GetPolicyOKBodyDeployTypeK8S captures enum value "K8S"
	GetPolicyOKBodyDeployTypeK8S string = "K8S"

	// GetPolicyOKBodyDeployTypeVM captures enum value "VM"
	GetPolicyOKBodyDeployTypeVM string = "VM"
)

// prop value enum
func (o *GetPolicyOKBody) validateDeployTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getPolicyOKBodyTypeDeployTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetPolicyOKBody) validateDeployType(formats strfmt.Registry) error {

	if swag.IsZero(o.DeployType) { // not required
		return nil
	}

	// value enum
	if err := o.validateDeployTypeEnum("getPolicyOK"+"."+"deployType", "body", o.DeployType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPolicyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPolicyOKBody) UnmarshalBinary(b []byte) error {
	var res GetPolicyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
