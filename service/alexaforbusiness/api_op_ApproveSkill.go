// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a skill with the organization under the customer's AWS account. If a
// skill is private, the user implicitly accepts access to this skill during
// enablement.
func (c *Client) ApproveSkill(ctx context.Context, params *ApproveSkillInput, optFns ...func(*Options)) (*ApproveSkillOutput, error) {
	if params == nil {
		params = &ApproveSkillInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApproveSkill", params, optFns, addOperationApproveSkillMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApproveSkillOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApproveSkillInput struct {

	// The unique identifier of the skill.
	//
	// This member is required.
	SkillId *string
}

type ApproveSkillOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationApproveSkillMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpApproveSkill{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpApproveSkill{}, middleware.After)
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
	addOpApproveSkillValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opApproveSkill(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opApproveSkill(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ApproveSkill",
	}
}
