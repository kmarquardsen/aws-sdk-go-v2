// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the specified Contributor Insights rules. When rules are enabled, they
// immediately begin analyzing log data.
func (c *Client) EnableInsightRules(ctx context.Context, params *EnableInsightRulesInput, optFns ...func(*Options)) (*EnableInsightRulesOutput, error) {
	if params == nil {
		params = &EnableInsightRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableInsightRules", params, optFns, addOperationEnableInsightRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableInsightRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableInsightRulesInput struct {

	// An array of the rule names to enable. If you need to find out the names of your
	// rules, use DescribeInsightRules
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeInsightRules.html).
	//
	// This member is required.
	RuleNames []*string
}

type EnableInsightRulesOutput struct {

	// An array listing the rules that could not be enabled. You cannot disable or
	// enable built-in rules.
	Failures []*types.PartialFailure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableInsightRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableInsightRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableInsightRules{}, middleware.After)
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
	addOpEnableInsightRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableInsightRules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableInsightRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "EnableInsightRules",
	}
}
