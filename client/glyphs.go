package client

const (
	unicodeRule     = "\u2500" // BOX DRAWINGS LIGHT HORIZONTAL
	unicodeEllipsis = "\u2026" // HORIZONTAL ELLIPSIS
	asciiRule       = "-"
	asciiEllipsis   = "..."
)

// consoleProbe is the console check ResolveBoxGlyphs and EncodeForConsole ask, a var so a test can pin it.
var consoleProbe = consoleIsUnicode

// ResolveBoxGlyphs returns the rule and ellipsis this console can carry.
func ResolveBoxGlyphs() (rule, ellipsis string) {
	if consoleProbe() {
		return unicodeRule, unicodeEllipsis
	}
	return asciiRule, asciiEllipsis
}
