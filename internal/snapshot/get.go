package snapshot

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lucasreed/snapman/internal/tags"
)

// Get will return a list of Snap structs
func Get(e *ec2.EC2, t []string) []*ec2.Snapshot {
	ret := []*ec2.Snapshot{}
	for _, tag := range t {
		if err := tags.Validate(&tag); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	filters, err := tags.CreateTagFilter(t)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	params := &ec2.DescribeSnapshotsInput{
		Filters: filters,
	}
	err = e.DescribeSnapshotsPages(params,
		func(page *ec2.DescribeSnapshotsOutput, lastPage bool) bool {
			for _, v := range page.Snapshots {
				ret = append(ret, v)
			}
			return !lastPage
		})
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}
	return ret
}
