package tags

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// CreateTagFilter will take a single string and create an ec2 Filters struct from it
func CreateTagFilter(t []string) ([]*ec2.Filter, error) {
	var ret []*ec2.Filter
	for _, tag := range t {
		split := strings.Split(tag, "=")
		if len(split) < 2 {
			e := fmt.Errorf("incorrect tag format, should be 'key=value', got: %v", tag)
			return ret, e
		}
		key := split[:len(split)-1][0]
		value := split[len(split)-1]
		this := ec2.Filter{
			Name: aws.String(fmt.Sprintf("tag:%v", key)),
			Values: []*string{
				aws.String(value),
			},
		}
		ret = append(ret, &this)
	}
	return ret, nil
}
