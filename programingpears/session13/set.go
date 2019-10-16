package pearls

type IntSet interface {
	Insert(int)
	Size() int
	Report() []int
}
