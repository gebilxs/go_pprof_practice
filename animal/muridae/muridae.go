package muridae

import "bytedance/byte_dance_projects/day_3_go-pprof-practice/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
