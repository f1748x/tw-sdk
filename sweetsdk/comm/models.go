package comm

import "fmt"

type Com struct {
	AppId string
	Secrt string
	State bool
	// Name  string
}

func (m *Com) Authentication() string {
	fmt.Println("执行------------Authentication")
	return m.AppId
}
