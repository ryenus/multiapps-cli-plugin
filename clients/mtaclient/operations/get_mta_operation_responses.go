// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	baseclient "github.com/cloudfoundry/multiapps-cli-plugin/clients/baseclient"
	"github.com/cloudfoundry/multiapps-cli-plugin/clients/models"
)

// GetMtaOperationReader is a Reader for the GetMtaOperation structure.
type GetMtaOperationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMtaOperationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMtaOperationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, baseclient.BuildErrorResponse(response, consumer, o.formats)
	}
}

// NewGetMtaOperationOK creates a GetMtaOperationOK with default headers values
func NewGetMtaOperationOK() *GetMtaOperationOK {
	return &GetMtaOperationOK{}
}

/*GetMtaOperationOK handles this case with default header values.

OK
*/
type GetMtaOperationOK struct {
	Payload *models.Operation
}

func (o *GetMtaOperationOK) Error() string {
	return fmt.Sprintf("[GET /operations/{operationId}][%d] getMtaOperationOK  %+v", 200, o.Payload)
}

func (o *GetMtaOperationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Operation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
