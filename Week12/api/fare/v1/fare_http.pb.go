// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.1

package fare

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type FareServiceHTTPServer interface {
	CreateFare(context.Context, *CreateFareRequest) (*CreateFareReply, error)
	DeleteFare(context.Context, *DeleteFareRequest) (*DeleteFareReply, error)
	GetFare(context.Context, *GetFareRequest) (*GetFareReply, error)
	PriceCalendar(context.Context, *PriceCalendarRequest) (*PriceCalendarResponse, error)
	Pricing(context.Context, *PricingRequest) (*PricingResponse, error)
	UpdateFare(context.Context, *UpdateFareRequest) (*UpdateFareReply, error)
}

func RegisterFareServiceHTTPServer(s *http.Server, srv FareServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/fare/", _FareService_CreateFare0_HTTP_Handler(srv))
	r.PUT("/v1/fare/{id}", _FareService_UpdateFare0_HTTP_Handler(srv))
	r.DELETE("/v1/fare/{id}", _FareService_DeleteFare0_HTTP_Handler(srv))
	r.GET("/v1/fare/{id}", _FareService_GetFare0_HTTP_Handler(srv))
	r.POST("/v1/fare/pricing", _FareService_Pricing0_HTTP_Handler(srv))
	r.POST("/v1/fare/calendar", _FareService_PriceCalendar0_HTTP_Handler(srv))
}

func _FareService_CreateFare0_HTTP_Handler(srv FareServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateFareRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fare.FareService/CreateFare")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateFare(ctx, req.(*CreateFareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateFareReply)
		return ctx.Result(200, reply)
	}
}

func _FareService_UpdateFare0_HTTP_Handler(srv FareServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateFareRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fare.FareService/UpdateFare")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateFare(ctx, req.(*UpdateFareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateFareReply)
		return ctx.Result(200, reply)
	}
}

func _FareService_DeleteFare0_HTTP_Handler(srv FareServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteFareRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fare.FareService/DeleteFare")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteFare(ctx, req.(*DeleteFareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteFareReply)
		return ctx.Result(200, reply)
	}
}

func _FareService_GetFare0_HTTP_Handler(srv FareServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFareRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fare.FareService/GetFare")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFare(ctx, req.(*GetFareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFareReply)
		return ctx.Result(200, reply)
	}
}

func _FareService_Pricing0_HTTP_Handler(srv FareServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PricingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fare.FareService/Pricing")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Pricing(ctx, req.(*PricingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PricingResponse)
		return ctx.Result(200, reply)
	}
}

func _FareService_PriceCalendar0_HTTP_Handler(srv FareServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PriceCalendarRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/fare.FareService/PriceCalendar")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PriceCalendar(ctx, req.(*PriceCalendarRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PriceCalendarResponse)
		return ctx.Result(200, reply)
	}
}

type FareServiceHTTPClient interface {
	CreateFare(ctx context.Context, req *CreateFareRequest, opts ...http.CallOption) (rsp *CreateFareReply, err error)
	DeleteFare(ctx context.Context, req *DeleteFareRequest, opts ...http.CallOption) (rsp *DeleteFareReply, err error)
	GetFare(ctx context.Context, req *GetFareRequest, opts ...http.CallOption) (rsp *GetFareReply, err error)
	PriceCalendar(ctx context.Context, req *PriceCalendarRequest, opts ...http.CallOption) (rsp *PriceCalendarResponse, err error)
	Pricing(ctx context.Context, req *PricingRequest, opts ...http.CallOption) (rsp *PricingResponse, err error)
	UpdateFare(ctx context.Context, req *UpdateFareRequest, opts ...http.CallOption) (rsp *UpdateFareReply, err error)
}

type FareServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewFareServiceHTTPClient(client *http.Client) FareServiceHTTPClient {
	return &FareServiceHTTPClientImpl{client}
}

func (c *FareServiceHTTPClientImpl) CreateFare(ctx context.Context, in *CreateFareRequest, opts ...http.CallOption) (*CreateFareReply, error) {
	var out CreateFareReply
	pattern := "/v1/fare/"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/fare.FareService/CreateFare"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FareServiceHTTPClientImpl) DeleteFare(ctx context.Context, in *DeleteFareRequest, opts ...http.CallOption) (*DeleteFareReply, error) {
	var out DeleteFareReply
	pattern := "/v1/fare/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/fare.FareService/DeleteFare"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FareServiceHTTPClientImpl) GetFare(ctx context.Context, in *GetFareRequest, opts ...http.CallOption) (*GetFareReply, error) {
	var out GetFareReply
	pattern := "/v1/fare/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/fare.FareService/GetFare"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FareServiceHTTPClientImpl) PriceCalendar(ctx context.Context, in *PriceCalendarRequest, opts ...http.CallOption) (*PriceCalendarResponse, error) {
	var out PriceCalendarResponse
	pattern := "/v1/fare/calendar"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/fare.FareService/PriceCalendar"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FareServiceHTTPClientImpl) Pricing(ctx context.Context, in *PricingRequest, opts ...http.CallOption) (*PricingResponse, error) {
	var out PricingResponse
	pattern := "/v1/fare/pricing"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/fare.FareService/Pricing"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FareServiceHTTPClientImpl) UpdateFare(ctx context.Context, in *UpdateFareRequest, opts ...http.CallOption) (*UpdateFareReply, error) {
	var out UpdateFareReply
	pattern := "/v1/fare/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/fare.FareService/UpdateFare"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
