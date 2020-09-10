// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Composes an email message and immediately queues it for sending. In order to
// send email using the SendEmail operation, your message must meet the following
// requirements:  <ul> <li> <p>The message must be sent from a verified email
// address or domain. If you attempt to send email using a non-verified address or
// domain, the operation will result in an "Email address not verified" error. </p>
// </li> <li> <p>If your account is still in the Amazon SES sandbox, you may only
// send to verified addresses or domains, or to email addresses associated with the
// Amazon SES Mailbox Simulator. For more information, see <a
// href="https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html">Verifying
// Email Addresses and Domains</a> in the <i>Amazon SES Developer Guide.</i> </p>
// </li> <li> <p>The maximum message size is 10 MB.</p> </li> <li> <p>The message
// must include at least one recipient email address. The recipient address can be
// a To: address, a CC: address, or a BCC: address. If a recipient email address is
// invalid (that is, it is not in the format
// <i>UserName@[SubDomain.]Domain.TopLevelDomain</i>), the entire message will be
// rejected, even if the message contains other recipients that are valid.</p>
// </li> <li> <p>The message may not include more than 50 recipients, across the
// To:, CC: and BCC: fields. If you need to send an email message to a larger
// audience, you can divide your recipient list into groups of 50 or fewer, and
// then call the <code>SendEmail</code> operation several times to send the message
// to each group.</p> </li> </ul> <important> <p>For every message that you send,
// the total number of recipients (including each recipient in the To:, CC: and
// BCC: fields) is counted against the maximum number of emails you can send in a
// 24-hour period (your <i>sending quota</i>). For more information about sending
// quotas in Amazon SES, see <a
// href="https://docs.aws.amazon.com/ses/latest/DeveloperGuide/manage-sending-limits.html">Managing
// Your Amazon SES Sending Limits</a> in the <i>Amazon SES Developer Guide.</i>
// </p> </important>
func (c *Client) SendEmail(ctx context.Context, params *SendEmailInput, optFns ...func(*Options)) (*SendEmailOutput, error) {
	stack := middleware.NewStack("SendEmail", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSendEmailMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpSendEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendEmail(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "SendEmail",
			Err:           err,
		}
	}
	out := result.(*SendEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a single formatted email using Amazon SES. For more
// information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-formatted.html).
type SendEmailInput struct {
	// The reply-to email address(es) for the message. If the recipient replies to the
	// message, each reply-to address will receive the reply.
	ReplyToAddresses []*string
	// The email address that is sending the email. This email address must be either
	// individually verified with Amazon SES, or from a domain that has been verified
	// with Amazon SES. For information about verifying identities, see the Amazon SES
	// Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html).
	// If you are sending on behalf of another user and have been permitted to do so by
	// a sending authorization policy, then you must also specify the SourceArn
	// parameter. For more information about sending authorization, see the Amazon SES
	// Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	// Amazon SES does not support the SMTPUTF8 extension, as described in RFC6531
	// (https://tools.ietf.org/html/rfc6531). For this reason, the local part of a
	// source email address (the part of the email address that precedes the @ sign)
	// may only contain 7-bit ASCII characters
	// (https://en.wikipedia.org/wiki/Email_address#Local-part). If the domain part of
	// an address (the part after the @ sign) contains non-ASCII characters, they must
	// be encoded using Punycode, as described in RFC3492
	// (https://tools.ietf.org/html/rfc3492.html). The sender name (also known as the
	// friendly name) may contain non-ASCII characters. These characters must be
	// encoded using MIME encoded-word syntax, as described in RFC 2047
	// (https://tools.ietf.org/html/rfc2047). MIME encoded-word syntax uses the
	// following form: =?charset?encoding?encoded-text?=.
	Source *string
	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to send for the email address specified in the Source parameter. For
	// example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to send from user@example.com, then you would specify the
	// SourceArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com, and the
	// Source to be user@example.com. For more information about sending authorization,
	// see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	SourceArn *string
	// The name of the configuration set to use when you send an email using SendEmail.
	ConfigurationSetName *string
	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send using SendEmail. Tags correspond to characteristics of the email that you
	// define, so that you can publish email sending events.
	Tags []*types.MessageTag
	// The email address that bounces and complaints will be forwarded to when feedback
	// forwarding is enabled. If the message cannot be delivered to the recipient, then
	// an error message will be returned from the recipient's ISP; this message will
	// then be forwarded to the email address specified by the ReturnPath parameter.
	// The ReturnPath parameter is never overwritten. This email address must be either
	// individually verified with Amazon SES, or from a domain that has been verified
	// with Amazon SES.
	ReturnPath *string
	// The message to be sent.
	Message *types.Message
	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the ReturnPath parameter. For example,
	// if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to use feedback@example.com, then you would specify the
	// ReturnPathArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com, and
	// the ReturnPath to be feedback@example.com. For more information about sending
	// authorization, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	ReturnPathArn *string
	// The destination for this email, composed of To:, CC:, and BCC: fields.
	Destination *types.Destination
}

// Represents a unique message ID.
type SendEmailOutput struct {
	// The unique message identifier returned from the SendEmail action.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSendEmailMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSendEmail{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSendEmail{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendEmail",
	}
}