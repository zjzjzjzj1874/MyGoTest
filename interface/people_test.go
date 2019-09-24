package people

import "testing"

func TestDoPeople(t *testing.T) {
	m := Man{Name: "Alan"}
	DoPeople(&m)

	w := Woman{Name: "Diana"}
	DoPeople(&w)
}
