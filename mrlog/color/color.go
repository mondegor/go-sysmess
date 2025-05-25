package color

const (
	Blue        = "\033[94m" // Blue - синий.
	Cyan        = "\033[36m" // Cyan - бирюзовый.
	LightCyan   = "\033[96m" // LightCyan - светло бирюзовый.
	Green       = "\033[32m" // Green - зелёный.
	LightGreen  = "\033[92m" // LightGreen - светло зелёный.
	Magenta     = "\033[95m" // Magenta - сиреневый.
	Red         = "\033[91m" // Red - красный.
	Yellow      = "\033[33m" // Yellow - желтый.
	LightYellow = "\033[93m" // LightYellow - светло желтый.
	Gray        = "\033[90m" // Gray - серый.
	LightGray   = "\033[37m" // LightGray - светло серый.
	White       = "\033[97m" // White - белый.
	End         = "\033[0m"  // End - завершающий символ цвета.
)

// ColorizeText - возвращает указанный текст с в указанном цвете для отображения в консоле.
func ColorizeText(colorCode, value string) string {
	return colorCode + value + End
}
