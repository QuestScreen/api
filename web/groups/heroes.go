package groups

// Hero describes a hero (player character).
type Hero interface {
	// Name of the hero
	Name() string
	// ID of the hero
	ID() string
	// Short description (e.g. class/race/etc)
	Description() string
}

// HeroList describes the list of heroes.
type HeroList interface {
	Hero(index int) Hero
	NumHeroes() int
}
