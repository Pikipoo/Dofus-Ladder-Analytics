package api

type characterClass int
type queueType string
type leagueName int

// Character classes enum
const (
	Feca       characterClass = 1
	Osamodas   characterClass = 2
	Enutrof    characterClass = 3
	Sram       characterClass = 4
	Xelor      characterClass = 5
	Ecaflip    characterClass = 6
	Eniripsa   characterClass = 7
	Iop        characterClass = 8
	Cra        characterClass = 9
	Sadida     characterClass = 10
	Sacrieur   characterClass = 11
	Pandawa    characterClass = 12
	Roublard   characterClass = 13
	Zobal      characterClass = 14
	Steamer    characterClass = 15
	Eliotrope  characterClass = 16
	Huppermage characterClass = 17
	Ouginak    characterClass = 18
)

// PvP queues enum
const (
	Solo  queueType = "solo" // 3v3 Solo
	Group queueType = "team" // 3v3 Team
	Duel  queueType = "duel" // 1v1 Duel
)

// PvP ranking leagues enum
const (
	Legend leagueName = 47
)

const (
	baseURL     = "https://www.dofus.com/fr/mmorpg/communaute/ladder/kolizeum?"
	paramQueue  = "&type=%s"
	paramSeason = "&season=%d"
	paramLeague = "&league=%d"
	paramClass  = "&breeds=%d"
)
