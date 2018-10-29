package sess

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetClient takes a region string and returns an ec2 client
func GetClient(r *string, p *string) *ec2.EC2 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(*r),
		Credentials: credentials.NewSharedCredentials("", *p),
	}))
	return ec2.New(sess)
}
