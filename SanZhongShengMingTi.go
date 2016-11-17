// SanZhongShengMingTi
package main

//"fmt"

const (
	LIMBO = iota
	LIVE
	DEAD

	KICK
	KILL
)

func recv() int {
	return KILL
}

func main() {
	state := LIMBO
	for {
		switch note := recv(); state {
		case LIMBO:
			switch note {
			case KICK:
				state = LIVE
			case KILL:
				state = DEAD
			}
		case LIVE:
			switch note {
			case KICK:
				state = LIMBO
			case KILL:
				state = DEAD
			}
		case DEAD:
			switch note {
			case KICK:
				state = LIMBO
			case KILL:
				panic("double kill")
			}
		}
	}
}
