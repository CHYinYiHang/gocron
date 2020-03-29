package bootstartp

import (
	"fmt"
	"github.com/CHYinYiHang/gocron/pkg/logging"
	"io/ioutil"
)

func InitApplication()  {
	content, err := ioutil.ReadFile("./icon")
	if err != nil {
		fmt.Println(err)
		return
	}
	logging.Info(string(content))
	//fmt.Println(string(content))
}
