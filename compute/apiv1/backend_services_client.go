// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newBackendServicesClientHook clientHook

// BackendServicesCallOptions contains the retry settings for each method of BackendServicesClient.
type BackendServicesCallOptions struct {
	AddSignedUrlKey    []gax.CallOption
	AggregatedList     []gax.CallOption
	Delete             []gax.CallOption
	DeleteSignedUrlKey []gax.CallOption
	Get                []gax.CallOption
	GetHealth          []gax.CallOption
	Insert             []gax.CallOption
	List               []gax.CallOption
	Patch              []gax.CallOption
	SetSecurityPolicy  []gax.CallOption
	Update             []gax.CallOption
}

// internalBackendServicesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalBackendServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddSignedUrlKey(context.Context, *computepb.AddSignedUrlKeyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListBackendServicesRequest, ...gax.CallOption) (*computepb.BackendServiceAggregatedList, error)
	Delete(context.Context, *computepb.DeleteBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	DeleteSignedUrlKey(context.Context, *computepb.DeleteSignedUrlKeyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetBackendServiceRequest, ...gax.CallOption) (*computepb.BackendService, error)
	GetHealth(context.Context, *computepb.GetHealthBackendServiceRequest, ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error)
	Insert(context.Context, *computepb.InsertBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListBackendServicesRequest, ...gax.CallOption) (*computepb.BackendServiceList, error)
	Patch(context.Context, *computepb.PatchBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetSecurityPolicy(context.Context, *computepb.SetSecurityPolicyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Update(context.Context, *computepb.UpdateBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// BackendServicesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The BackendServices API.
type BackendServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalBackendServicesClient

	// The call options for this service.
	CallOptions *BackendServicesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BackendServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BackendServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BackendServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddSignedUrlKey adds a key for validating requests with signed URLs for this backend service.
func (c *BackendServicesClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddSignedUrlKey(ctx, req, opts...)
}

// AggregatedList retrieves the list of all BackendService resources, regional and global, available to the specified project.
func (c *BackendServicesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListBackendServicesRequest, opts ...gax.CallOption) (*computepb.BackendServiceAggregatedList, error) {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified BackendService resource.
func (c *BackendServicesClient) Delete(ctx context.Context, req *computepb.DeleteBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DeleteSignedUrlKey deletes a key for validating requests with signed URLs for this backend service.
func (c *BackendServicesClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.DeleteSignedUrlKey(ctx, req, opts...)
}

// Get returns the specified BackendService resource. Gets a list of available backend services.
func (c *BackendServicesClient) Get(ctx context.Context, req *computepb.GetBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendService, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetHealth gets the most recent health check results for this BackendService.
//
// Example request body:
//
// { “group”: “/zones/us-east1-b/instanceGroups/lb-backend-example” }
func (c *BackendServicesClient) GetHealth(ctx context.Context, req *computepb.GetHealthBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error) {
	return c.internalClient.GetHealth(ctx, req, opts...)
}

// Insert creates a BackendService resource in the specified project using the data included in the request. For more information, see  Backend services overview.
func (c *BackendServicesClient) Insert(ctx context.Context, req *computepb.InsertBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of BackendService resources available to the specified project.
func (c *BackendServicesClient) List(ctx context.Context, req *computepb.ListBackendServicesRequest, opts ...gax.CallOption) (*computepb.BackendServiceList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified BackendService resource with the data included in the request. For more information, see  Backend services overview. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *BackendServicesClient) Patch(ctx context.Context, req *computepb.PatchBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetSecurityPolicy sets the Google Cloud Armor security policy for the specified backend service. For more information, see Google Cloud Armor Overview
func (c *BackendServicesClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSecurityPolicy(ctx, req, opts...)
}

// Update updates the specified BackendService resource with the data included in the request. For more information, see Backend services overview.
func (c *BackendServicesClient) Update(ctx context.Context, req *computepb.UpdateBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type backendServicesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBackendServicesRESTClient creates a new backend services rest client.
//
// The BackendServices API.
func NewBackendServicesRESTClient(ctx context.Context, opts ...option.ClientOption) (*BackendServicesClient, error) {
	clientOpts := append(defaultBackendServicesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &backendServicesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &BackendServicesClient{internalClient: c, CallOptions: &BackendServicesCallOptions{}}, nil
}

func defaultBackendServicesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *backendServicesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *backendServicesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *backendServicesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddSignedUrlKey adds a key for validating requests with signed URLs for this backend service.
func (c *backendServicesRESTClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetSignedUrlKeyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v/addSignedUrlKey", req.GetProject(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// AggregatedList retrieves the list of all BackendService resources, regional and global, available to the specified project.
func (c *backendServicesRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListBackendServicesRequest, opts ...gax.CallOption) (*computepb.BackendServiceAggregatedList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/backendServices", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.IncludeAllScopes != nil {
		params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.BackendServiceAggregatedList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified BackendService resource.
func (c *backendServicesRESTClient) Delete(ctx context.Context, req *computepb.DeleteBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v", req.GetProject(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// DeleteSignedUrlKey deletes a key for validating requests with signed URLs for this backend service.
func (c *backendServicesRESTClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v/deleteSignedUrlKey", req.GetProject(), req.GetBackendService())

	params := url.Values{}
	if req.GetKeyName() != "" {
		params.Add("keyName", fmt.Sprintf("%v", req.GetKeyName()))
	}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get returns the specified BackendService resource. Gets a list of available backend services.
func (c *backendServicesRESTClient) Get(ctx context.Context, req *computepb.GetBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendService, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v", req.GetProject(), req.GetBackendService())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.BackendService{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// GetHealth gets the most recent health check results for this BackendService.
//
// Example request body:
//
// { “group”: “/zones/us-east1-b/instanceGroups/lb-backend-example” }
func (c *backendServicesRESTClient) GetHealth(ctx context.Context, req *computepb.GetHealthBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetResourceGroupReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v/getHealth", req.GetProject(), req.GetBackendService())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.BackendServiceGroupHealth{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a BackendService resource in the specified project using the data included in the request. For more information, see  Backend services overview.
func (c *backendServicesRESTClient) Insert(ctx context.Context, req *computepb.InsertBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetBackendServiceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// List retrieves the list of BackendService resources available to the specified project.
func (c *backendServicesRESTClient) List(ctx context.Context, req *computepb.ListBackendServicesRequest, opts ...gax.CallOption) (*computepb.BackendServiceList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.BackendServiceList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Patch patches the specified BackendService resource with the data included in the request. For more information, see  Backend services overview. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *backendServicesRESTClient) Patch(ctx context.Context, req *computepb.PatchBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetBackendServiceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v", req.GetProject(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// SetSecurityPolicy sets the Google Cloud Armor security policy for the specified backend service. For more information, see Google Cloud Armor Overview
func (c *backendServicesRESTClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetSecurityPolicyReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/backendServices/%v/setSecurityPolicy", req.GetProject(), req.GetBackendService())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Update updates the specified BackendService resource with the data included in the request. For more information, see Backend services overview.
func (c *backendServicesRESTClient) Update(ctx context.Context, req *computepb.UpdateBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetBackendServiceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	if req.GetBackendService() != "" {
		params.Add("backendService", fmt.Sprintf("%v", req.GetBackendService()))
	}
	if req.GetProject() != "" {
		params.Add("project", fmt.Sprintf("%v", req.GetProject()))
	}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}
