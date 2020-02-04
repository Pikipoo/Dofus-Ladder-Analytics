package api

import (
	"errors"
	"strconv"
)

type character struct {
	ranking   int
	name      string
	server    string
	class     characterClass
	isOmega   bool
	level     int
	points    int
	victories int
}

func newCharater(ranking int, name string, server string, class characterClass, isOmega bool,
	level int, points int, victories int) *character {
	character := character{
		ranking:   ranking,
		name:      name,
		server:    server,
		class:     class,
		isOmega:   isOmega,
		level:     level,
		points:    points,
		victories: victories}

	return &character
}

func getCharacterClass(class string) characterClass {
	characterClass := characterClass(class)
	_, ok := CharacterClassMap[characterClass]
	if ok != true {
		return NONE
	}
	return characterClass
}

func newCharaterFromArray(dataArray []string) (character *character, err error) {
	if len(dataArray) < 7 {
		return nil, errors.New("Invalid data provided")
	}

	ranking, err := strconv.Atoi(dataArray[0])
	if err != nil {
		return nil, err
	}
	name, server, class, isOmega := dataArray[1], dataArray[2], getCharacterClass(dataArray[3]), false
	i := 4
	if len(dataArray) > 7 {
		isOmega = true
		i++
	}
	level, err := strconv.Atoi(dataArray[i])
	if err != nil {
		return nil, err
	}
	points, err := strconv.Atoi(dataArray[i+1])
	if err != nil {
		return nil, err
	}
	victories, err := strconv.Atoi(dataArray[i+2])
	if err != nil {
		return nil, err
	}

	return newCharater(ranking, name, server, class, isOmega, level, points, victories), nil
}
