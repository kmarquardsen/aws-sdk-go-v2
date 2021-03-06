// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all supported Elasticsearch versions
func (c *Client) ListElasticsearchVersions(ctx context.Context, params *ListElasticsearchVersionsInput, optFns ...func(*Options)) (*ListElasticsearchVersionsOutput, error) {
	if params == nil {
		params = &ListElasticsearchVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListElasticsearchVersions", params, optFns, addOperationListElasticsearchVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListElasticsearchVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListElasticsearchVersions operation. Use
// MaxResults to control the maximum number of results to retrieve in a single
// call. Use NextToken in response to retrieve more results. If the received
// response does not contain a NextToken, then there are no more results to
// retrieve.
type ListElasticsearchVersionsInput struct {

	// Set this value to limit the number of results returned. Value provided must be
	// greater than 10 else it wont be honored.
	MaxResults *int32

	// Paginated APIs accepts NextToken input to returns next page results and provides
	// a NextToken output in the response which can be used by the client to retrieve
	// more results.
	NextToken *string
}

// Container for the parameters for response received from
// ListElasticsearchVersions operation.
type ListElasticsearchVersionsOutput struct {

	// List of supported elastic search versions.
	ElasticsearchVersions []*string

	// Paginated APIs accepts NextToken input to returns next page results and provides
	// a NextToken output in the response which can be used by the client to retrieve
	// more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListElasticsearchVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListElasticsearchVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListElasticsearchVersions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListElasticsearchVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListElasticsearchVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListElasticsearchVersions",
	}
}
