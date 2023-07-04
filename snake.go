package main

type SnakeSeg struct {
	X int
	Y int
}

type SnakeBody struct {
	Segments  []SnakeSeg
	direction int //0=up, 1=right, 2=down, 3=left
}

func (s *SnakeBody) update(width int, height int, foodEaten bool) {
	s.Segments = append(s.Segments, s.Segments[len(s.Segments)-1].getUpdatedSeg(s, width, height))
	if !foodEaten {
		s.Segments = s.Segments[1:]
	}
}

func (ss *SnakeSeg) getUpdatedSeg(s *SnakeBody, width int, height int) SnakeSeg {
	newSeg := *ss
	if s.direction == 0 {
		newSeg.Y -= 1
		if newSeg.Y < 0 {
			newSeg.Y = height
		}
	} else if s.direction == 1 {
		newSeg.X = (newSeg.X + 1) % width
	} else if s.direction == 2 {
		newSeg.Y = (newSeg.Y + 1) % height
	} else if s.direction == 3 {
		newSeg.X -= 1
		if newSeg.X < 0 {
			newSeg.X = width
		}
	}
	return newSeg
}

func (s *SnakeBody) changeDirection(dir int) {
	s.direction = dir
}

func (s *SnakeSeg) display() rune {
	return ' '
}

func (s *SnakeBody) reset() {
	snakeParts := []SnakeSeg{
		{
			X: 5,
			Y: 10,
		},
		{
			X: 6,
			Y: 10,
		},
		{
			X: 7,
			Y: 10,
		},
	}

	s.Segments = snakeParts
	s.direction = 1
}