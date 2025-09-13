package color

// Коды цветов используемых в консоли.
const (
	Blue        = "\033[94m" // синий
	Cyan        = "\033[36m" // бирюзовый
	LightCyan   = "\033[96m" // светло бирюзовый
	Green       = "\033[32m" // зелёный
	LightGreen  = "\033[92m" // светло зелёный
	Magenta     = "\033[95m" // сиреневый
	Red         = "\033[91m" // красный
	Yellow      = "\033[33m" // желтый
	LightYellow = "\033[93m" // светло желтый
	Gray        = "\033[90m" // серый
	LightGray   = "\033[37m" // светло серый
	White       = "\033[97m" // белый
	End         = "\033[0m"  // завершающий символ цвета
)

// ColorizeText - возвращает указанный текст в указанном цвете для отображения в консоли.
func ColorizeText(colorCode, value string) string {
	return colorCode + value + End
}
