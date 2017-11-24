package toolkit

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type randomizer struct {
	sync.Mutex
	r *rand.Rand
}

var (
	randomizerInitialized uint32
	r                     *randomizer
	rmu                   = new(sync.Mutex)
)

func (r *randomizer) Intn(limit int) int {
	defer r.Unlock()
	r.Lock()
	return r.r.Intn(limit)
}

func initRandomSource() {
	if atomic.LoadUint32(&randomizerInitialized) == 1 {
		return
	}

	rmu.Lock()
	defer rmu.Unlock()

	if r == nil {
		src := rand.NewSource(time.Now().UnixNano())
		r = new(randomizer)
		r.r = rand.New(src)
	}
}

func RandInt(limit int) int {
	initRandomSource()
	return r.Intn(limit)
}

func RandFloat(limit int, decimal int) float64 {
	initRandomSource()
	return float64(r.Intn(limit+decimal)) / float64(10*decimal)
}

func Div(f1, f2 float64) float64 {
	if f2 == 0 {
		return 0
	}

	return f1 / f2
}
