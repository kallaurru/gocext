package gocext

import (
	"testing"
	"github.com/kallaurru/gocext"
)

func TestMakeDirs(t *testing.T) {
	root := "/media/common"
	gocext.MakeCommonInfrastructure(root)
	
}
