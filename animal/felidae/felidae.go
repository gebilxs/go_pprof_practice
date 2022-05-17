package felidae

import "bytedance/byte_dance_projects/day_3_go-pprof-practice/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
