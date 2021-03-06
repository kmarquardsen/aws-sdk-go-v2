// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about one or more specified events. Information
// includes standard event data (region, service, and so on, as returned by
// DescribeEvents), a detailed event description, and possible additional metadata
// that depends upon the nature of the event. Affected entities are not included;
// to retrieve those, use the DescribeAffectedEntities operation. If a specified
// event cannot be retrieved, an error message is returned for that event.
func (c *Client) DescribeEventDetails(ctx context.Context, params *DescribeEventDetailsInput, optFns ...func(*Options)) (*DescribeEventDetailsOutput, error) {
	if params == nil {
		params = &DescribeEventDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventDetails", params, optFns, addOperationDescribeEventDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventDetailsInput struct {

	// A list of event ARNs (unique identifiers). For example:
	// "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	//
	// This member is required.
	EventArns []*string

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string
}

type DescribeEventDetailsOutput struct {

	// Error messages for any events that could not be retrieved.
	FailedSet []*types.EventDetailsErrorItem

	// Information about the events that could be retrieved.
	SuccessfulSet []*types.EventDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventDetails{}, middleware.After)
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
	addOpDescribeEventDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventDetails(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventDetails",
	}
}
