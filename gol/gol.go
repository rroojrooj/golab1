package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	IMHT := p.imageHeight
	IMWD := p.imageWidth

	newWorld := make([][]byte, IMHT)
	for i := range newWorld {
		newWorld[i] = make([]byte, IMWD)
	}

	// Copy the state from the old world
	// Need to do new state (create newWorld) since mutating along the way on the original one will change the state that will use to cal on the next iteration.
	// So that will make it wrong
	for y := 0; y < IMHT; y++ {
		for x := 0; x < IMWD; x++ {
			newWorld[y][x] = world[y][x]
		}
	}

	for y := 0; y < IMHT; y++ {
		for x := 0; x < IMWD; x++ {

			// Compute the sum of the eight neighbors using explicit modular arithmetic
			sumNeighbors :=
				(world[(y+IMHT-1)%IMHT][(x+IMWD-1)%IMWD] / 255) + // NW: Northwest neighbor
					(world[(y+IMHT-1)%IMHT][(x+IMWD)%IMWD] / 255) + // N: North neighbor
					(world[(y+IMHT-1)%IMHT][(x+IMWD+1)%IMWD] / 255) + // NE: Northeast neighbor
					(world[(y+IMHT)%IMHT][(x+IMWD-1)%IMWD] / 255) + // W: West neighbor
					(world[(y+IMHT)%IMHT][(x+IMWD+1)%IMWD] / 255) + // E: East neighbor
					(world[(y+IMHT+1)%IMHT][(x+IMWD-1)%IMWD] / 255) + // SW: Southwest neighbor
					(world[(y+IMHT+1)%IMHT][(x+IMWD)%IMWD] / 255) + // S: South neighbor
					(world[(y+IMHT+1)%IMHT][(x+IMWD+1)%IMWD] / 255) // SE: Southeast neighbor

			// Applying the rules of the game
			if world[y][x] >= 255 { // If the cell is alive
				if sumNeighbors < 2 {
					newWorld[y][x] = 0 // Rule: any live cell with fewer than two live neighbours dies
				} else if sumNeighbors == 2 || sumNeighbors == 3 {
					newWorld[y][x] = 255 // Rule: any live cell with two or three live neighbours is unaffected
				} else {
					newWorld[y][x] = 0 // Rule: any live cell with more than three live neighbours dies
				}
			} else { // If the cell is dead
				if sumNeighbors == 3 {
					newWorld[y][x] = 255 // Rule: any dead cell with exactly three live neighbours becomes alive
				} else {
					newWorld[y][x] = 0
				}
			}
		}
	}

	world = newWorld
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var aliveCells []cell // Slice to store the coordinates of alive cells

	// Iterate over the world
	for y, row := range world {
		for x, cellValue := range row {
			// If the cell is alive
			if cellValue == 255 {
				aliveCells = append(aliveCells, cell{x, y}) // Append the cell's coordinates to the slice
			}
		}
	}
	return aliveCells // Return the slice of alive cells' coordinates
}
