package rlet_test

import (
	"computedriver/pkg/rlet"
	"fmt"
	"testing"
)

func TestNewClusterIP(t *testing.T) {
	clusterIp := rlet.NewClusterIP()

	fmt.Println(clusterIp)
}
