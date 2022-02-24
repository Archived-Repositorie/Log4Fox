package log4fox

type color string

type Font struct {
	BLACK       color
	RED         color
	GREEN       color
	YELLOW      color
	BLUE        color
	PURPLE      color
	CYAN        color
	WHITE       color
	BLACK_BOLD  color
	RED_BOLD    color
	GREEN_BOLD  color
	YELLOW_BOLD color
	BLUE_BOLD   color
	PURPLE_BOLD color
	CYAN_BOLD   color
	WHITE_BOLD  color
	BLACK_UL    color
	RED_UL      color
	GREEN_UL    color
	YELLOW_UL   color
	BLUE_UL     color
	PURPLE_UL   color
	CYAN_UL     color
	WHITE_UL    color
}

type BG struct {
	BLACK     color
	RED       color
	GREEN     color
	YELLOW    color
	BLUE      color
	PURPLE    color
	CYAN      color
	WHITE     color
	BLACK_HI  color
	RED_HI    color
	GREEN_HI  color
	YELLOW_HI color
	BLUE_HI   color
	PURPLE_HI color
	CYAN_HI   color
	WHITE_HI  color
}

type Colors struct {
	RESET string
	Font  Font
	BG    BG
}

func Color() Colors {
	font := Font{
		BLACK:       "\033[0;30m",
		RED:         "\033[0;31m",
		GREEN:       "\033[0;32m",
		YELLOW:      "\033[0;33m",
		BLUE:        "\033[0;34m",
		PURPLE:      "\033[0;35m",
		CYAN:        "\033[0;36m",
		WHITE:       "\033[0;37m",
		BLACK_BOLD:  "\033[1;30m",
		RED_BOLD:    "\033[1;31m",
		GREEN_BOLD:  "\033[1;32m",
		YELLOW_BOLD: "\033[1;33m",
		BLUE_BOLD:   "\033[1;34m",
		PURPLE_BOLD: "\033[1;35m",
		CYAN_BOLD:   "\033[1;36m",
		WHITE_BOLD:  "\033[1;37m",
		BLACK_UL:    "\033[4;30m",
		RED_UL:      "\033[4;31m",
		GREEN_UL:    "\033[4;32m",
		YELLOW_UL:   "\033[4;33m",
		BLUE_UL:     "\033[4;34m",
		PURPLE_UL:   "\033[4;35m",
		CYAN_UL:     "\033[4;36m",
		WHITE_UL:    "\033[4;37m",
	}

	bg := BG{
		BLACK:     "\033[40m",
		RED:       "\033[41m",
		GREEN:     "\033[42m",
		YELLOW:    "\033[43m",
		BLUE:      "\033[44m",
		PURPLE:    "\033[45m",
		CYAN:      "\033[46m",
		WHITE:     "\033[47m",
		BLACK_HI:  "\033[0;100m",
		RED_HI:    "\033[0;101m",
		GREEN_HI:  "\033[0;102m",
		YELLOW_HI: "\033[0;103m",
		BLUE_HI:   "\033[0;104m",
		PURPLE_HI: "\033[0;105m",
		CYAN_HI:   "\033[0;106m",
		WHITE_HI:  "\033[0;107m",
	}
	c := Colors{
		RESET: "\033[m",
		Font:  font,
		BG:    bg,
	}

	return c
}
