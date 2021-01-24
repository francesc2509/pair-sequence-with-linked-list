package sequence

import "github.com/francesc2509/pair-sequence-with-linked-list/entities"

func join(pairs []map[string]int, node *entities.Node) *entities.Node {
	if pairs == nil || len(pairs) == 0 {
		return node
	}

	pair := pairs[0]
	pairs = pairs[1:]

	if node == nil {
		return join(pairs, entities.NewNode(pair, nil))
	}

	if n := combine(pair, node); n != nil {
		return join(pairs, n)
	}

	if pair["start"] < node.Start() {
		tmp := node
		n := entities.NewNode(pair, tmp)

		return join(pairs, n)
	}

	node.SetNext(join([]map[string]int{pair}, node.Next()))

	return join(pairs, node)
}

func combine(pair map[string]int, node *entities.Node) *entities.Node {
	start := pair["start"]
	end := pair["end"]

	nStart := node.Start()
	nEnd := node.End()

	if start >= nStart && end <= nEnd {
		return node
	}

	min := start
	minHigh := nStart

	if nStart < min {
		min = nStart
		minHigh = start
	}

	max := end
	maxLow := nEnd
	if nEnd > end {
		max = nEnd
		maxLow = end
	}

	canCombine := ((minHigh >= min && minHigh <= maxLow) && (maxLow >= min && maxLow <= max)) || minHigh == maxLow+1
	if canCombine {
		value := map[string]int{"start": min, "end": max}
		n := join([]map[string]int{value}, node.Next())

		return n
	}

	return nil
}
