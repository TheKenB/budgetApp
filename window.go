package main

type resolution struct {
	x int
	y int
}

func res720() resolution {
	return resolution{x: 1280, y: 720}
}

func res1080() resolution {
	return resolution{x: 1920, y: 1080}
}

func res1440() resolution {
	return resolution{x: 2560, y: 1440}
}
