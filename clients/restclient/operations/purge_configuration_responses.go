// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	baseclient "github.com/cloudfoundry/multiapps-cli-plugin/clients/baseclient"
)

// PurgeConfigurationReader is a Reader for the PurgeConfiguration structure.
type PurgeConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurgeConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPurgeConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, baseclient.BuildErrorResponse(response, consumer, o.formats)
	}
}

// NewPurgeConfigurationNoContent creates a PurgeConfigurationNoContent with default headers values
func NewPurgeConfigurationNoContent() *PurgeConfigurationNoContent {
	return &PurgeConfigurationNoContent{}
}

/*PurgeConfigurationNoContent handles this case with default header values.

No Content
*/
type PurgeConfigurationNoContent struct {
}

func (o *PurgeConfigurationNoContent) Error() string {
	return fmt.Sprintf("[POST /configuration-entries/purge][%d] purgeConfigurationNoContent ", 204)
}

func (o *PurgeConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
