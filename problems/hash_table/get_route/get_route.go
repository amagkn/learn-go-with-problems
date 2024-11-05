package get_route

type Ticket struct {
	From string
	To   string
}

/*
GetRoute

Arranges the tickets one after another to create a continuous route.
*/
func GetRoute(tickets []Ticket) []Ticket {
	from := make(map[string]Ticket)
	to := make(map[string]Ticket)

	for _, ticket := range tickets {
		from[ticket.From] = ticket
		to[ticket.To] = ticket
	}

	var start Ticket

	for _, ticket := range tickets {
		if _, ok := to[ticket.From]; !ok {
			start = ticket
			break
		}
	}

	var route []Ticket
	curr := start

	for curr != (Ticket{}) {
		route = append(route, curr)
		curr = from[curr.To]
	}

	return route
}
