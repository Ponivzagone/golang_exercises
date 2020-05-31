package main
 
import (
    "fmt"
    "os"
    "bufio"
	"strconv"
	"strings"
)

type HeightTree struct {
	tree []int64
	jump []int64
}

func MakeHeightTree(n int) *HeightTree {
	obj := new(HeightTree)
	obj.tree = make([]int64, n)
	obj.jump = make([]int64, n)
	return obj
}

func (HT *HeightTree) ReadFrom(value string) {
	for i, val := range strings.SplitN(value[0:len(value)-1], " ", -1) {
		fnum, err := strconv.ParseInt(val,10, 64)
		if err != nil { panic(err) }
		HT.tree[i] =  fnum
	}
}

func (HT *HeightTree)JumpTo(idx int64) int64 {
	
	if idx == -1 { 
		return -1
	}

	if HT.jump[idx] != 0 {
		return HT.jump[idx]
	}

	HT.jump[idx] = HT.JumpTo(HT.tree[idx]) + int64(1)
	
	return HT.jump[idx]
}

func (HT *HeightTree) ComputJump() {
	for i, val := range HT.tree {
		HT.jump[i] = HT.JumpTo(val) + int64(1)
	}
}

func (HT *HeightTree) GetHeight() int64 {
	var max int64 = HT.jump[0]

	for _, value := range HT.jump {
		if max < value {
			max = value
		}
	}

	return  max
}
 
func main()  {
    reader := bufio.NewReader(os.Stdin)
    value, _ := reader.ReadString('\n')
    n, err := strconv.Atoi(value[:len(value)-1])
	if err != nil { panic(err) }
	
	HT := MakeHeightTree(n)

	value, _ = reader.ReadString('\n')
	HT.ReadFrom(value)

	fmt.Println("Node tree: ", HT.tree)
	HT.ComputJump()
	fmt.Println("Jump for each node: ", HT.jump)

	fmt.Println("Height: ", HT.GetHeight())
}