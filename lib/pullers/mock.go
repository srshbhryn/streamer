package pullers

import (
	"strconv"
	"time"
)

type MockPuller struct{}

func CreateMock() MockPuller {
	return MockPuller{}
}

var counter int

func (_ MockPuller) Read() (string, error) {
	time.Sleep(300 * time.Millisecond)
	counter += 1
	if counter%2 == 0 {
		return "a," + strconv.Itoa(counter), nil
	}
	return "b," + strconv.Itoa(counter), nil

}
