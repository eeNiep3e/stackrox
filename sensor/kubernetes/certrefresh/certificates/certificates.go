package certificates

import (
	"context"

	"github.com/stackrox/rox/generated/internalapi/central"
	"github.com/stackrox/rox/generated/storage"
)

// Requester defines an interface for requesting TLS certificates from Central
type Requester interface {
	Start()
	Stop()
	RequestCertificates(ctx context.Context) (*Response, error)
}

// Response represents the response to a certificate request. It contains a set of certificates or an error.
type Response struct {
	RequestId    string
	ErrorMessage *string
	Certificates *storage.TypedServiceCertificateSet
}

// NewResponseFromLocalScannerCerts creates a certificates.Response from a
// protobuf central.IssueLocalScannerCertsResponse message
func NewResponseFromLocalScannerCerts(response *central.IssueLocalScannerCertsResponse) *Response {
	if response == nil {
		return nil
	}

	res := &Response{
		RequestId: response.GetRequestId(),
	}

	if response.GetError() != nil {
		errMsg := response.GetError().GetMessage()
		res.ErrorMessage = &errMsg
	} else {
		res.Certificates = response.GetCertificates()
	}

	return res
}
