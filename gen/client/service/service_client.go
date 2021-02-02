// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListTasks(params *ListTasksParams) (*ListTasksOK, *ListTasksNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListTasks retrieve list of task
*/
func (a *Client) ListTasks(params *ListTasksParams) (*ListTasksOK, *ListTasksNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listTasks",
		Method:             "GET",
		PathPattern:        "/listProducts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListTasksOK:
		return value, nil, nil
	case *ListTasksNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for service: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
