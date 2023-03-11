package tracker

import (
	"github.com/aabri-assignments/assignment-golang/helper"
)

type Tracker interface {
	Track() Flight
}

type tracker struct {
	flightTracker []Flight
	sources       []string
	destinations  []string
}

func New(flights []Flight) Tracker {
	t := &tracker{
		flightTracker: flights,
	}

	starts := make([]string, len(t.flightTracker))

	ends := make([]string, len(t.flightTracker))

	t.sources = starts

	t.destinations = ends

	return t
}

func (t *tracker) Track() Flight {
	if len(t.flightTracker) == 1 {
		return t.flightTracker[0]
	}

	for i, v := range t.flightTracker {
		t.sources[i], t.destinations[i] = v.Source, v.Destination
	}

	return Flight{
		Source:      helper.FindFirstDifference(t.sources, t.destinations),
		Destination: helper.FindFirstDifference(t.destinations, t.sources),
	}
}
