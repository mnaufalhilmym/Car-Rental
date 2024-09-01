package util

func ParseComparisonFilter(comparison string) string {
	switch comparison {
	case "gt":
		return ">"
	case "gte":
		return ">="
	case "lt":
		return "<"
	case "lte":
		return "<="
	default:
		return ""
	}
}
