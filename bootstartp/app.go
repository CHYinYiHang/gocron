package bootstartp

import (
	"fmt"
	"io/ioutil"
)

func InitApplication()  {
	content, err := ioutil.ReadFile("./icon")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
