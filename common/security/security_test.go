package security

import (
	"testing"
	"fmt"
)

func TestMd5(t *testing.T)  {
	str := Md5("zealous")
	fmt.Println(str)
}
