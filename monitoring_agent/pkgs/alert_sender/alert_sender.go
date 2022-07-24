package alert_sender

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendAlert(recipient string, 
                subject string, 
                message string,
                html string,
                creds map[string]string) (int, string) {

    // Create a new session in the us-west-2 region.
    // Replace us-west-2 with the AWS Region you're using for Amazon SES.
    sess, err := session.NewSession(&aws.Config{
        Region:aws.String(creds["aws_region"]),
		Credentials:credentials.NewStaticCredentials(
            creds["aws_access_key"],
            creds["aws_secret_key"],
            "")},
    )

    if err != nil {
        return http.StatusUnprocessableEntity, fmt.Sprintf("Error creating session: %s", err)
    }
    
    // Create an SES session.
    svc := ses.New(sess)
    
    // Assemble the email.
    input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            CcAddresses: []*string{
            },
            ToAddresses: []*string{
                aws.String(recipient),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{                
                Html: &ses.Content{
                    Charset: aws.String(creds["charset"]),
                    Data:    aws.String(html),
                },
                Text: &ses.Content{
                    Charset: aws.String(creds["charset"]),
                    Data:    aws.String(message),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String(creds["charset"]),
                Data:    aws.String(subject),
            },
        },
        Source: aws.String(creds["from_address"]),
    }

    // Attempt to send the email.
    result, err := svc.SendEmail(input)
    
    // Display error messages if they occur.
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case ses.ErrCodeMessageRejected:
                return http.StatusUnprocessableEntity, fmt.Sprintln(ses.ErrCodeMessageRejected, aerr.Error())
            case ses.ErrCodeMailFromDomainNotVerifiedException:
                return http.StatusUnprocessableEntity, fmt.Sprintln(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
            case ses.ErrCodeConfigurationSetDoesNotExistException:                
                return http.StatusUnprocessableEntity, fmt.Sprintln(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
            default:
                return http.StatusUnprocessableEntity, aerr.Error()
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            return http.StatusUnprocessableEntity, err.Error()
        }
    }

    return http.StatusOK, fmt.Sprintf("Email successfully Sent to address: %s", recipient) + fmt.Sprintf(" with: %s", result)
}