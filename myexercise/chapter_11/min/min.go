package min

type Miner interface {
	Less(i, j int) bool
	Len() int
	At(i int) any
}

func Min(data Miner) (res any) {
	res = data.At(0)
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			res = data.At(i)
		}
	}
	return
}

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) At(i int) any       { return p[i] }
