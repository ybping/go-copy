package encoding

import (
	"testing"
)

type location struct {
	Lat float32
	Lng float64
}

type user struct {
	Name   string
	Age    int
	Phones []string
	Addr   *location
	Work   map[string]string
}

type admin struct {
	Role   string
	Name   string
	Age    int
	Addr   *location
	Phones []string
	Work   map[string]string
}

// just basic is this working stuff
func TestSimple(t *testing.T) {
	{
		var a, b int
		a = 10
		bi, err := DeepCopy(b, a)
		if err != nil || bi.(int) != a {
			t.Error(bi, err)
		}
	}
	{
		admin1 := admin{"super", "admin1", 10, nil, []string{"100-1", "100-2"}, nil}
		user1, err := DeepCopy(user{}, admin1)
		if err != nil {
			t.Error(user1, err)
		}
	}
	{
		addr := &location{121.1, 90.2}
		work := map[string]string{"monday": "C++", "wensday": "Java"}
		admin1 := admin{"super", "admin1", 10, addr, []string{"100-1", "100-2"}, work}
		user1, err := DeepCopy(user{}, admin1)
		if err != nil {
			t.Error(user1, err)
		}
	}
}
