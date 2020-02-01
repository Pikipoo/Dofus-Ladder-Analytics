package api

type character struct {
	name      string
	server    string
	class     characterClass
	isOmega   bool
	level     int
	points    int
	victories int
}

func newCharater(name string, server string, class characterClass, isOmega bool,
	level int, points int, victories int) *character {
	character := character{
		name:      name,
		server:    server,
		class:     class,
		isOmega:   isOmega,
		level:     level,
		points:    points,
		victories: victories}

	return &character
}
