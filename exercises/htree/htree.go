package htree
 
import (
	"strconv"
	"strings"
)

const dsc =`Требуется посчитать высоту дерева										
																					
Правила составления корневого дерева:												
    - На вход подается массив переходов между узлами дерева в формате: 4 -1 4 1 1 2 5
    - -1: корень дерева															
																					
Пример: 																			
            t[1] = -1 																
           /         \ 															
      t[3] = 1     t[4] = 1														
                    /     \														
              t[0] = 4   t[2] = 4													
                             |														
                         t[5] = 2													
                             |														
                         t[6] = 5													
																					
Высота для такого дерева будет = 4													`

func Description() string {
	return dsc
}

type HeightTree struct {
	tree []int64 // Переходы узлов дерева
	jump []int64 // Колчиство переходов к корню дерева для каждого узла
}

func (HT *HeightTree) Tree() []int64 {
	return HT.tree
}

func (HT *HeightTree) Jump() []int64 {
	return HT.jump
}

func MakeHeightTree(n int) *HeightTree {
	obj := new(HeightTree)
	obj.tree = make([]int64, n)
	obj.jump = make([]int64, n)
	return obj
}

func (HT *HeightTree) ReadFrom(value string) {
	for i, val := range strings.Fields(value) {
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