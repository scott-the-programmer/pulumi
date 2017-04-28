// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubefission

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for HTTPTrigger resource provider */

// HTTPTriggerToken is the type token corresponding to the HTTPTrigger package type.
const HTTPTriggerToken = tokens.Type("kubefission:httptrigger:HTTPTrigger")

// HTTPTriggerProviderOps is a pluggable interface for HTTPTrigger-related management functionality.
type HTTPTriggerProviderOps interface {
    Check(ctx context.Context, obj *HTTPTrigger) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *HTTPTrigger) (string, error)
    Get(ctx context.Context, id string) (*HTTPTrigger, error)
    InspectChange(ctx context.Context,
        id string, old *HTTPTrigger, new *HTTPTrigger, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id string, old *HTTPTrigger, new *HTTPTrigger, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id string) error
}

// HTTPTriggerProvider is a dynamic gRPC-based plugin for managing HTTPTrigger resources.
type HTTPTriggerProvider struct {
    ops HTTPTriggerProviderOps
}

// NewHTTPTriggerProvider allocates a resource provider that delegates to a ops instance.
func NewHTTPTriggerProvider(ops HTTPTriggerProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &HTTPTriggerProvider{ops: ops}
}

func (p *HTTPTriggerProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *HTTPTriggerProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *HTTPTriggerProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   id,
    }, nil
}

func (p *HTTPTriggerProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    id := req.GetId()
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *HTTPTriggerProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    diff := oldprops.Diff(newprops)
    var replaces []string
    if diff.Changed("name") {
        replaces = append(replaces, "name")
    }
    more, err := p.ops.InspectChange(ctx, req.GetId(), old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *HTTPTriggerProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, req.GetId(), old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *HTTPTriggerProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(HTTPTriggerToken))
    if err := p.ops.Delete(ctx, req.GetId()); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *HTTPTriggerProvider) Unmarshal(
    v *pbstruct.Struct) (*HTTPTrigger, resource.PropertyMap, mapper.DecodeError) {
    var obj HTTPTrigger
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable HTTPTrigger structure(s) */

// HTTPTrigger is a marshalable representation of its corresponding IDL type.
type HTTPTrigger struct {
    Name string `json:"name"`
    URLPattern string `json:"urlPattern"`
    Method string `json:"method"`
    Function *resource.ID `json:"function"`
}

// HTTPTrigger's properties have constants to make dealing with diffs and property bags easier.
const (
    HTTPTrigger_Name = "name"
    HTTPTrigger_URLPattern = "urlPattern"
    HTTPTrigger_Method = "method"
    HTTPTrigger_Function = "function"
)


