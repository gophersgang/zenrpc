// generated by zenrpc; DO NOT EDIT

package testdata

import (
	"context"
	"encoding/json"

	"github.com/sergeyfast/zenrpc"
	"github.com/sergeyfast/zenrpc/smd"
)

var RPC = struct {
	ArithService struct {
		Divide   string
		Multiply string
		Positive string
		Pow      string
		Sum      string
		SumArray string
	}

	PhoneBook struct {
		ById           string
		Delete         string
		Get            string
		Remove         string
		Save           string
		ValidateSearch string
	}
}{
	ArithService: struct {
		Divide   string
		Multiply string
		Positive string
		Pow      string
		Sum      string
		SumArray string
	}{

		Divide:   "divide",
		Multiply: "multiply",
		Positive: "positive",
		Pow:      "pow",
		Sum:      "sum",
		SumArray: "sumarray",
	},

	PhoneBook: struct {
		ById           string
		Delete         string
		Get            string
		Remove         string
		Save           string
		ValidateSearch string
	}{

		ById:           "byid",
		Delete:         "delete",
		Get:            "get",
		Remove:         "remove",
		Save:           "save",
		ValidateSearch: "validatesearch",
	},
}

func (ArithService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{}
}

// Invoke is as generated code from zenrpc cmd
func (s ArithService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}

	switch method {

	case RPC.ArithService.Divide:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.Divide(args.A, args.B))

	case RPC.ArithService.Multiply:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.Multiply(args.A, args.B))

	case RPC.ArithService.Positive:
		var args = struct {
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.Positive())

	case RPC.ArithService.Pow:
		var args = struct {
			Base float64  `json:"base"`
			Exp  *float64 `json:"exp"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		//zenrpc:exp:2 	exponent could be empty
		if args.Exp == nil {
			var v float64 = 2
			args.Exp = &v
		}

		resp.Set(s.Pow(args.Base, args.Exp))

	case RPC.ArithService.Sum:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.Sum(ctx, args.A, args.B))

	case RPC.ArithService.SumArray:
		var args = struct {
			Array *[]float64 `json:"array"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		//zenrpc:array:[]float64{1,2,4}
		if args.Array == nil {
			var v []float64 = []float64{1, 2, 4}
			args.Array = &v
		}

		resp.Set(s.SumArray(args.Array))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (PhoneBook) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{}
}

// Invoke is as generated code from zenrpc cmd
func (s PhoneBook) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}

	switch method {

	case RPC.PhoneBook.ById:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.ById(args.Id))

	case RPC.PhoneBook.Delete:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.Delete(args.Id))

	case RPC.PhoneBook.Get:
		var args = struct {
			Count  *int         `json:"count"`
			Page   *int         `json:"page"`
			Search PersonSearch `json:"search"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		//zenrpc:count:50 page size
		if args.Count == nil {
			var v int = 50
			args.Count = &v
		}

		//zenrpc:page:0 current page
		if args.Page == nil {
			var v int = 0
			args.Page = &v
		}

		resp.Set(s.Get(args.Search, args.Page, args.Count))

	case RPC.PhoneBook.Remove:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.Remove(args.Id))

	case RPC.PhoneBook.Save:
		var args = struct {
			P       Person `json:"p"`
			Replace *bool  `json:"replace"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		//zenrpc:replace:false update person if exist
		if args.Replace == nil {
			var v bool = false
			args.Replace = &v
		}

		resp.Set(s.Save(args.P, args.Replace))

	case RPC.PhoneBook.ValidateSearch:
		var args = struct {
			Search *PersonSearch `json:"search"`
		}{}

		if err := json.Unmarshal(params, &args); err != nil {
			return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, err.Error(), nil)
		}

		resp.Set(s.ValidateSearch(args.Search))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
