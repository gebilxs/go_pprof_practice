package animal

import (
	"bytedance/byte_dance_projects/day_3_go-pprof-practice/animal/canidae/dog"
	"bytedance/byte_dance_projects/day_3_go-pprof-practice/animal/canidae/wolf"
	"bytedance/byte_dance_projects/day_3_go-pprof-practice/animal/felidae/cat"
	"bytedance/byte_dance_projects/day_3_go-pprof-practice/animal/felidae/tiger"
	"bytedance/byte_dance_projects/day_3_go-pprof-practice/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
