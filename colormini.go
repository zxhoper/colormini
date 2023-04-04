package colormini

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Example() {
	NoteGreen("The color print setup is: OK")
	NoteTYellowf("Color Notef is %s!\n", "OK")

	Notefn("Format line with %s appended.", "NEWLINE")

	_______________Note_______________("hi")
	PrintAllColorNotef()

	PrintAllColor()

	PrintAllColorTitle()

	NotePrefix = "PERFIX "
	PrintAllColor()
	PrintAllColorTitle()

	ln := "line"
	nl := "newline"

	Notefn("Format %s with %s appended.", ln, nl)

	NotePrefix = ""
}

func PrintAllColor() {
	NoteBlockFirst()

	Note("Default Note")
	Notef("Default Notef\n")

	NoteGreen("Green")
	NoteRed("Red")
	NoteYellow("Yellow")
	NoteBlue("Blue")
	NotePurple("Purple")
	NoteBlockLast()

	NotePrtHr()
}
func PrintAllColorNotef() {
	NoteBlockFirst()

	Notef("Default Notef\n")

	NoteGreenf("Green")
	NoteRedf("Red")
	NoteYellowf("Yellow")
	NoteBluef("Blue")
	NotePurplef("Purple")

	NoteBlockLast()

	NotePrtHr()
}
func PrintAllColorTitle() {
	NoteBlockFirst()
	NoteT("Default NoteT")
	NoteTf("Default Note Title \n")
	NoteTGreenf("Green\n")
	NoteTRedf("Red\n")
	NoteTYellowf("Yellow\n")
	NoteTBluef("Blue\n")
	NoteTPurplef("Purple\n")
	NoteBlockLast()
}

// NotePrefix ColorMini ===============================================================================================
var NotePrefix = ""
var clr = map[string]string{"Non": "0", "Red": "31", "Green": "32", "Yellow": "33", "Blue": "34", "Purple": "35"}

func _______________Note_______________(a ...interface{}) { NoteColor(getCC(clr["Red"]), a...) }
func GetColorType() int                                   { return len(clr) }
func getCC(c string) string                               { return fmt.Sprintf("\033[%sm", c) }
func NoteColor(c string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, c)
	eewP(os.Stdout, NotePrefix, a...)
	_, _ = fmt.Fprintf(os.Stdout, "\033[0m")
}
func NoteGreen(a ...interface{})  { NoteColor(getCC(clr["Green"]), a...) }
func NoteRed(a ...interface{})    { NoteColor(getCC(clr["Red"]), a...) }
func NoteYellow(a ...interface{}) { NoteColor(getCC(clr["Yellow"]), a...) }
func NoteBlue(a ...interface{})   { NoteColor(getCC(clr["Blue"]), a...) }
func NotePurple(a ...interface{}) { NoteColor(getCC(clr["Purple"]), a...) }
func Note(a ...interface{})       { eewP(os.Stdout, NotePrefix, a...) }

// eewP  wraps fmt.Println with prefix
func eewP(w io.Writer, pfix string, a ...interface{}) {
	var prefixedAny []interface{}
	if pfix != "" {
		prefixedAny = append(prefixedAny, pfix)
	}
	prefixedAny = append(prefixedAny, a...)
	_, _ = fmt.Fprintln(w, prefixedAny...)
}
func NoteT(s string) { eewPT(os.Stdout, NotePrefix, s) }
func eewPT(w io.Writer, pfix string, s string) {
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
	_, _ = fmt.Fprintf(w, pfix+"== %-45s ==\n", s)
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
}
func Notef(format string, a ...interface{}) {
	eewPf(os.Stdout, NotePrefix, format, a...)
}
func NoteColorf(colorCode string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s", colorCode, format, getCC(clr["Non"]))
	eewPf(os.Stdout, NotePrefix, formatColor, a...)
}
func NoteGreenf(format string, a ...interface{})  { NoteColorf(getCC(clr["Green"]), format, a...) }
func NoteRedf(format string, a ...interface{})    { NoteColorf(getCC(clr["Red"]), format, a...) }
func NoteYellowf(format string, a ...interface{}) { NoteColorf(getCC(clr["Yellow"]), format, a...) }
func NoteBluef(format string, a ...interface{})   { NoteColorf(getCC(clr["Blue"]), format, a...) }
func NotePurplef(format string, a ...interface{}) { NoteColorf(getCC(clr["Purple"]), format, a...) }
func eewPf(w io.Writer, pfix string, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, pfix+format, a...)
}
func Notefn(format string, a ...interface{}) { eewPfn(os.Stdout, format, NotePrefix, a...) }
func eewPfn(w io.Writer, pfix string, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, pfix+format+"\n", a...)
}
func NoteTColorf(c string, format string, a ...interface{}) {
	formatColor := fmt.Sprintf("%s%s%s%s", c, "== ", format, getCC(clr["Non"]))
	prefixedFormatColor := NotePrefix + formatColor
	hrColor := fmt.Sprintf("%s%s%s", c, StringRepeat("=", 51), getCC(clr["Non"]))
	_, _ = fmt.Fprintln(os.Stdout, NotePrefix+hrColor)
	_, _ = fmt.Fprintf(os.Stdout, prefixedFormatColor, a...)
	_, _ = fmt.Fprintln(os.Stdout, NotePrefix+hrColor)
}
func NoteTRedf(format string, a ...interface{})    { NoteTColorf(getCC(clr["Red"]), format, a...) }
func NoteTGreenf(format string, a ...interface{})  { NoteTColorf(getCC(clr["Green"]), format, a...) }
func NoteTYellowf(format string, a ...interface{}) { NoteTColorf(getCC(clr["Yellow"]), format, a...) }
func NoteTBluef(format string, a ...interface{})   { NoteTColorf(getCC(clr["Blue"]), format, a...) }
func NoteTPurplef(format string, a ...interface{}) { NoteTColorf(getCC(clr["Purple"]), format, a...) }
func NoteTf(format string, a ...interface{})       { eewPft(os.Stdout, NotePrefix, format, a...) }
func eewPft(w io.Writer, pfix string, format string, a ...interface{}) {
	prefixedFormat := pfix + "== " + format
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
	_, _ = fmt.Fprintf(w, prefixedFormat, a...)
	_, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 51))
}
func NotePrtHr()                         { eewPhr(os.Stdout, NotePrefix) }
func eewPhr(w io.Writer, pfix string)    { _, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 72)) }
func NoteBlockFirst()                    { eewPFirst(os.Stdout, NotePrefix) }
func NoteBlockLast()                     { eewPLast(os.Stdout, NotePrefix) }
func eewPFirst(w io.Writer, pfix string) { _, _ = fmt.Fprintln(w, pfix+StringRepeat("-", 64)+"\\\\") }
func eewPLast(w io.Writer, pfix string)  { _, _ = fmt.Fprintln(w, pfix+StringRepeat("=", 64)+"//") }
func StringRepeat(s string, repeatTimes int) string {
	var sb strings.Builder
	for i := 0; i < repeatTimes; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

// ===============================================================================================
