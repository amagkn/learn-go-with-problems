package get_route

import (
	"reflect"
	"testing"
)

func TestGetRoute(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		got := GetRoute([]Ticket{{From: "Astana", To: "Bali"}, {From: "Moscow", To: "Astana"}, {From: "Bali", To: "SPB"}})
		want := []Ticket{{From: "Moscow", To: "Astana"}, {From: "Astana", To: "Bali"}, {From: "Bali", To: "SPB"}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
