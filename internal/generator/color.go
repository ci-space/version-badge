package generator

const DefaultColor = "#97ca00"

func ResolveColor(col string) string {
	switch col {
	case "blue":
		return "#007ec6"
	case "green":
		return "#97ca00"
	case "brightgreen":
		return "#44cc11"
	case "":
		return DefaultColor
	default:
		return col
	}
}
