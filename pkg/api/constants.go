package api

type characterClass string
type queueType string
type leagueName int

// Character classes enum
const (
	NONE       characterClass = "Unknown"
	FECA       characterClass = "Féca"
	OSAMODAS   characterClass = "Osamodas"
	ENUTROF    characterClass = "Enutrof"
	SRAM       characterClass = "Sram"
	XELOR      characterClass = "Xelor"
	ECAFLIP    characterClass = "Ecaflip"
	ENIRIPSA   characterClass = "Eniripsa"
	IOP        characterClass = "Iop"
	CRA        characterClass = "Cra"
	SADIDA     characterClass = "Sadida"
	SACRIEUR   characterClass = "Sacrieur"
	PANDAWA    characterClass = "Pandawa"
	ROUBLARD   characterClass = "Roublard"
	ZOBAL      characterClass = "Zobal"
	STEAMER    characterClass = "Steamer"
	ELIOTROPE  characterClass = "Eliotrope"
	HUPPERMAGE characterClass = "Huppermage"
	OUGINAK    characterClass = "Ouginak"
)

// CharacterClassMap map of all classes in the game related to their ID
var CharacterClassMap = map[characterClass]int{
	FECA:       1,
	OSAMODAS:   2,
	ENUTROF:    3,
	SRAM:       4,
	XELOR:      5,
	ECAFLIP:    6,
	ENIRIPSA:   7,
	IOP:        8,
	CRA:        9,
	SADIDA:     10,
	SACRIEUR:   11,
	PANDAWA:    12,
	ROUBLARD:   13,
	ZOBAL:      14,
	STEAMER:    15,
	ELIOTROPE:  16,
	HUPPERMAGE: 17,
	OUGINAK:    18,
}

// PvP queues enum
const (
	SOLO  queueType = "solo" // 3v3 Solo
	GROUP queueType = "team" // 3v3 Team
	DUEL  queueType = "duel" // 1v1 Duel
)

// PvP ranking leagues enum
const (
	LEGEND leagueName = 47
)

const (
	baseURL     = "https://www.dofus.com/fr/mmorpg/communaute/ladder/kolizeum?"
	paramQueue  = "&type=%s"
	paramSeason = "&season=%d"
	paramLeague = "&league=%d"
	paramClass  = "&breeds=%d"
)
