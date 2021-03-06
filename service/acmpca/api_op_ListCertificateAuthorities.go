// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the private certificate authorities that you created by using the
// CreateCertificateAuthority action.
func (c *Client) ListCertificateAuthorities(ctx context.Context, params *ListCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListCertificateAuthoritiesOutput, error) {
	if params == nil {
		params = &ListCertificateAuthoritiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificateAuthorities", params, optFns, addOperationListCertificateAuthoritiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificateAuthoritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCertificateAuthoritiesInput struct {

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int32

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string
}

type ListCertificateAuthoritiesOutput struct {

	// Summary information about each certificate authority you have created.
	CertificateAuthorities []*types.CertificateAuthority

	// When the list is truncated, this value is present and should be used for the
	// NextToken parameter in a subsequent pagination request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCertificateAuthoritiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCertificateAuthorities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCertificateAuthorities{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificateAuthorities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCertificateAuthorities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "ListCertificateAuthorities",
	}
}
