# Snapman - EBS snapshot manager
Tool that makes it easier to manage EBS snapshots in AWS assuming that there are useful tags on the snapshots to filter on.

It is a work in progress and the only functioning command at this point is to list snapshots based on tag filters.

## Why?
Most of this can be done with the [AWS CLI tool](https://aws.amazon.com/cli/), but it can be cumbersome to create filters. This tool is an opinionated way to get the most useful information.

The original intent of creating this tool was to practice golang. If it becomes useful for someone else, even better!
