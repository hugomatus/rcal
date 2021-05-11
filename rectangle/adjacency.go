package rectangle

// isAdjacent determines whether Rectangles R1 and R2 are adjacent
func (r1 Rectangle) isAdjacent(r2 Rectangle) bool {

	if !(len(r1.Intersection(r2)) == 0) || r1.Containment(r2) {
		return false
	}
	if r1.proper(r2) {
		return true
	}

	if r1.partial(r2) {
		return true
	}

	return false
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
