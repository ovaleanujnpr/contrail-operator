package e2e

import "time"

var (
	retryInterval          = time.Second * 5
	waitTimeout            = time.Second * 400
	cleanupRetryInterval   = time.Second * 1
	cleanupTimeout         = time.Second * 5
	waitForOperatorTimeout = time.Minute * 10
)
