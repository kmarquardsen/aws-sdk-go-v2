// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a credential report for the AWS account. For more information about
// the credential report, see Getting Credential Reports
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html) in
// the IAM User Guide.
func (c *Client) GenerateCredentialReport(ctx context.Context, params *GenerateCredentialReportInput, optFns ...func(*Options)) (*GenerateCredentialReportOutput, error) {
	if params == nil {
		params = &GenerateCredentialReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateCredentialReport", params, optFns, addOperationGenerateCredentialReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateCredentialReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateCredentialReportInput struct {
}

// Contains the response to a successful GenerateCredentialReport request.
type GenerateCredentialReportOutput struct {

	// Information about the credential report.
	Description *string

	// Information about the state of the credential report.
	State types.ReportStateType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGenerateCredentialReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGenerateCredentialReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGenerateCredentialReport{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateCredentialReport(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGenerateCredentialReport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GenerateCredentialReport",
	}
}
