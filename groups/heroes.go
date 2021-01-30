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

// HeroChangeAction is an enum describing a change in the list of heroes
type HeroChangeAction int

const (
	// HeroAdded describes the action of adding a hero to the list of heroes
	HeroAdded HeroChangeAction = iota
	// HeroModified describes the action of modifying a hero's data
	HeroModified
	// HeroDeleted describes the action of deleting a hero from the list of heroes
	HeroDeleted
)
