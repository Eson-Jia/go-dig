package pearls

import (
	"fmt"
	"testing"
)

func Test_Set(t *testing.T) {
	maxElements, maxVal := 1000, 10000
	// source :=
	// dst := []int{1, 2, 3, 4, 5, 6}
	suits := []struct {
		Source []int
		Dst    []int
	}{
		{Source: []int{1, 2, 3, 5, 6, 4, 4},
			Dst: []int{1, 2, 3, 4, 5, 6},
		},
	}

	sets := []struct {
		IntSet
		Name string
	}{
		{
			IntSet: NewIntSetArr(maxElements, maxVal),
			Name:   "IntSetArr",
		},
		{
			IntSet: NewIntSetSTL(maxElements, maxVal),
			Name:   "IntSetSTL",
		},
		{
			IntSet: NewIntSetList(maxElements, maxVal),
			Name:   "IntSetList",
		},
		{
			IntSet: NewIntSetBST(maxElements, maxVal),
			Name:   "IntSetBST",
		},
	}
	for _, set := range sets {
		for _, suit := range suits {
			for _, element := range suit.Source {
				set.Insert(element)
			}
			dist := set.Report()
			for i := 0; i < len(suit.Dst); i++ {
				if dist[i] != suit.Dst[i] {
					t.Errorf("test faied in source:%v,set type %s", suit.Source, set.Name)
					return
				}
			}
		}
		set.Report()
	}

}

// TestSlice make([]int,len,cap)
func TestSlice(t *testing.T) {
	array := make([]int, 0, 10)
	func(theDst []int) {
		theDst = append(theDst, 123)
		showSlice("first inner slice:", theDst)
	}(array)
	func(theDst []int) {
		theDst = append(theDst, 456)
		showSlice("second inner slice:", theDst)
	}(array)
	showSlice("outter slice:", array)
	showSlice("[:10] slice:", array[:10])
}

func showSlice(description string, s []int) {
	fmt.Println(description, "  length:", len(s), "capacity:", cap(s), "value:", s)
}

func TestMap(t *testing.T) {
	array := make(map[int]int)
	func(theDst map[int]int) {
		theDst[1] = 1
		fmt.Println(theDst)
	}(array)
	fmt.Println(array)
}
