// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of the specified receipt rule set. For information about
// managing receipt rule sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
// You can execute this operation no more than once per second.
func (c *Client) DescribeReceiptRuleSet(ctx context.Context, params *DescribeReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeReceiptRuleSetOutput, error) {
	if params == nil {
		params = &DescribeReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReceiptRuleSet", params, optFns, addOperationDescribeReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the details of a receipt rule set. You use
// receipt rule sets to receive email with Amazon SES. For more information, see
// the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DescribeReceiptRuleSetInput struct {

	// The name of the receipt rule set to describe.
	//
	// This member is required.
	RuleSetName *string
}

// Represents the details of the specified receipt rule set.
type DescribeReceiptRuleSetOutput struct {

	// The metadata for the receipt rule set, which consists of the rule set name and
	// the timestamp of when the rule set was created.
	Metadata *types.ReceiptRuleSetMetadata

	// A list of the receipt rules that belong to the specified receipt rule set.
	Rules []*types.ReceiptRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReceiptRuleSet{}, middleware.After)
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
	addOpDescribeReceiptRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReceiptRuleSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeReceiptRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DescribeReceiptRuleSet",
	}
}
