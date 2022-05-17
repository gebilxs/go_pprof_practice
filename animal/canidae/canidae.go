package canidae

import "bytedance/byte_dance_projects/day_3_go-pprof-practice/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
