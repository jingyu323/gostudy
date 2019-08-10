package clent

var (
	clentName string
)

//定义构造
type ClientTest struct {
	Host      string
	clentName string
}

//getter
func (c *ClientTest) ClentName() string {
	return c.clentName
}

//setter
func (c *ClientTest) SetClentName(name string) {
	c.clentName = name
}

func NewClientTest(host string) *ClientTest {
	return &ClientTest{
		Host:      host,
		clentName: "44",
	}
}
