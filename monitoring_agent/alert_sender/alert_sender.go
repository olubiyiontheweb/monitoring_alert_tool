package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/joho/godotenv"
)

const (
    
    // Replace recipient@example.com with a "To" address. If your account 
    // is still in the sandbox, this address must be verified.
    Recipient = "	olubiyiontheweb@gmail.com"
    
    // The HTML body for the email.
    HtmlBody =  "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
                "<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
                "<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
    
    //The email body for recipients with non-HTML email clients.
    TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
)

func main() {
	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}
    // Create a new session in the us-west-2 region.
    // Replace us-west-2 with the AWS Region you're using for Amazon SES.
    sess, err := session.NewSession(&aws.Config{
        Region:aws.String(os.Getenv("AWS_REGION")),
		Credentials:credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")},
    )
    
    // Create an SES session.
    svc := ses.New(sess)
    
    // Assemble the email.
    input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            CcAddresses: []*string{
            },
            ToAddresses: []*string{
                aws.String(Recipient),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Html: &ses.Content{
                    Charset: aws.String(os.Getenv("CHARSET")),
                    Data:    aws.String(HtmlBody),
                },
                Text: &ses.Content{
                    Charset: aws.String(os.Getenv("CHARSET")),
                    Data:    aws.String(TextBody),
                },
            },
            Subject: &ses.Content{
                Charset: aws.String(os.Getenv("CHARSET")),
                Data:    aws.String(os.Getenv("SUBJECT")),
            },
        },
        Source: aws.String(os.Getenv("FROM_EMAIL")),
            // Uncomment to use a configuration set
            //ConfigurationSetName: aws.String(ConfigurationSet),
    }

    // Attempt to send the email.
    result, err := svc.SendEmail(input)
    
    // Display error messages if they occur.
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case ses.ErrCodeMessageRejected:
                fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
            case ses.ErrCodeMailFromDomainNotVerifiedException:
                fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
            case ses.ErrCodeConfigurationSetDoesNotExistException:
                fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
    
        return
    }
    
    fmt.Println("Email Sent to address: " + Recipient)
    fmt.Println(result)
}