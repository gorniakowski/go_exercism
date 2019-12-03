package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var max = 26 * 26 * 10 * 10 * 10
var issuedRobots = make(map[string]bool, 1)
var err error

//Robot represents robot
type Robot struct {
	name string
}

//Name return robot's name or generates new name
func (r *Robot) Name() (string, error) {

	if issuedRobots[r.name] {
		return r.name, nil
	}

	r.name, err = genUniqName()
	if err == nil {
		issuedRobots[r.name] = true
	}
	return r.name, err
}

//Reset sets new name to a robot
func (r *Robot) Reset() {

	r.name, err = genUniqName()
	if err != nil {
		issuedRobots[r.name] = true
	}
}

func genUniqName() (string, error) {
	if len(issuedRobots) == max {
		return "", errors.New("nameSpace is exausecd")
	}
	rand.Seed(time.Now().UnixNano())
	var name string
	for ok := true; ok; ok = issuedRobots[name] {
		name = ""
		for i := 0; i < 5; i++ {
			val := rand.Intn(26)
			if i < 2 {
				name += string(letters[val])
				continue
			}
			name += strconv.Itoa(val % 10)

		}
	}
	return name, nil
}
