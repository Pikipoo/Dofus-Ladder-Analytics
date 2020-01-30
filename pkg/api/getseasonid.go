package api

// GetSeasonID takes the aimed season to analyze for and returns the right
// season ID, correcting the season ID offset created by Ankama
func GetSeasonID(season int) int {
	if season < 2 {
		return season
	}
	return season + 1
}
