package singleton

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetInstance(t *testing.T) {
	convey.Convey("singleton test", t, func() {
		for i := 0; i < 5; i++ {
			tempInstance := GetInstance()
			fmt.Println(&tempInstance)
		}
	})
}
