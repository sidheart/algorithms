package ctci

type Pixel struct {
	r,g,b,a byte
}

type image [][]Pixel

// Equal returns true if image i is equal to other
// Two images are equal if the contain the same Pixels in the same order
func (i image) Equal(other image) bool {
	if (i == nil) != (other == nil) {
		return false
	}
	if len(i) != len(other) {
		return false
	}
	if len(i) == 0 {
		return true
	}
	if len(i[0]) != len(other[0]) {
		return false
	}
	for row, _ := range i {
		for col, _ := range i[0] {
			if i[row][col] != other[row][col] {
				return false
			}
		}
	}
	return true
}

// rotateCW rotates an image clockwise
func rotateCW(im image) image {
	if len(im) == 0 || len(im[0]) == 0 {
		return nil
	}
	m, n := len(im), len(im[0])
	var rotIm image
	for col := 0; col < n; col++ {
		var newRow []Pixel
		for row := m - 1; row >= 0; row-- {
			newRow = append(newRow, im[row][col])
		}
		rotIm = append(rotIm, newRow)
	}
	return rotIm
}
