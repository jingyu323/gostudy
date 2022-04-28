package best_pricatice

import (
	"encoding/binary"
	"io"
)

//先处理错误避免嵌套
type Gopher struct {
	Name     string
	AgeYears int
}

func main() {
	gofer := &Gopher{
		"", 9,
	}
	io.WriteString()
	gofer.WriteTo()

}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4
	n, err := w.Write([]byte(g.Name))
	size += int64(n)
	if err != nil {
		return
	}
	err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
	if err == nil {
		size += 4
	}
	return
}
