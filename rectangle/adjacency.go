package rectangle

type Adjacency string

const (
	Proper     Adjacency = "Proper"
	Partial    Adjacency = "Partial"
	SubLine    Adjacency = "Sub-line"
	NoAjacency Adjacency = "None identified"
)

// isAdjacent determines whether Rectangles R1 and R2 are adjacent
func (r1 Rectangle) isAdjacent(r2 Rectangle) (bool, Adjacency) {

	if r1.proper(r2) {
		return true, Proper
	}

	if r1.partial(r2) {
		return true, Partial
	}

	if r1.subline(r2) {
		return true, SubLine
	}

	return false, NoAjacency
}

// proper
func (r1 Rectangle) proper(r2 Rectangle) bool {

	//top or bottom
	if r1.xMin() == r2.xMin() {
		//top
		if r1.yMax() == r2.yMin() {
			LoadFigure(19)
			return true
		}
		//bottom
		if r1.yMin() == r2.yMax() {
			LoadFigure(18)
			return true
		}
	}

	return false
}

// partial
func (r1 Rectangle) partial(r2 Rectangle) bool {

	//left
	if r1.xMin() < r2.xMax() && r1.xMin() > r2.xMin() {
		//bottom
		if r1.yMin() == r2.yMax() {
			LoadFigure(23)
			return true
		}

		//top
		if r1.yMax() == r2.yMin() {

			LoadFigure(21)
			return true
		}

	}

	//right
	if r1.xMax() > r2.xMin() && r1.xMax() < r2.xMax() {
		//bottom
		if r1.yMin() == r2.yMax() {
			LoadFigure(22)
			return true
		}
		//top
		if r1.yMax() == r2.yMin() {

			LoadFigure(23)
			return true
		}
	}

	return false
}

//TODO LoadFig

// subline
func (r1 Rectangle) subline(r2 Rectangle) bool {
	//left or right
	if r1.yMax() > r2.yMax() {
		if r1.yMin() < r2.yMin() {
			//right
			if r1.xMax() == r2.xMin() {
				//LoadFig
				return true
			}

			//left
			if r1.xMin() == r2.xMax() {
				//LoadFig
				return true
			}
		}
	}

	//top or bottom
	if r1.xMax() > r2.xMax() && r1.xMin() < r2.xMin() {

		//bottom
		if r1.yMin() == r2.yMax() {
			//LoadFig
			return true
		}

		//Top
		if r1.yMax() == r2.yMin() {
			//LoadFig
			return true
		}
	}

	return false

}
