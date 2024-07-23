package comm

import "fmt"

type Com struct {
	AppId string
	Secrt string
	State bool
	// Name  string
}
type Playse struct {
	Id   int
	Name string
}

func (m *Com) Authentication() string {
	fmt.Println("执行------------Authentication")
	return m.AppId
}

func (m *Com) GetPclas() *Playse {
	c := new(Playse)
	c.Id = 1111
	c.Name = "测试内容"
	fmt.Println("执行------------Authentication")
	return c
}
