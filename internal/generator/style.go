package generator

type Style string

const (
	StyleUnspecified Style = ""
	StyleFlat        Style = "flat"
	StyleFlagSquare  Style = "flat-square"
	StylePlastic     Style = "plastic"
)

func (s Style) Valid() bool {
	switch s {
	case StyleUnspecified:
		return false
	case StyleFlat:
		return true
	case StyleFlagSquare:
		return true
	case StylePlastic:
		return true
	}

	return false
}
