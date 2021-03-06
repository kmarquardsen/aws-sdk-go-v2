// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes customizations for the QuickSight subscription on your AWS account.
func (c *Client) DeleteAccountCustomization(ctx context.Context, params *DeleteAccountCustomizationInput, optFns ...func(*Options)) (*DeleteAccountCustomizationOutput, error) {
	if params == nil {
		params = &DeleteAccountCustomizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccountCustomization", params, optFns, addOperationDeleteAccountCustomizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccountCustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccountCustomizationInput struct {

	// The ID for the AWS account that you want to delete QuickSight customizations
	// from.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace associated with the customization that you're deleting.
	Namespace *string
}

type DeleteAccountCustomizationOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccountCustomizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAccountCustomization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAccountCustomization{}, middleware.After)
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
	addOpDeleteAccountCustomizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccountCustomization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAccountCustomization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteAccountCustomization",
	}
}
