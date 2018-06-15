package ass

import (
	"bufio"
	"fmt"
	"io"
)

// percentGraphColored prints with ansii a bar graph of win percentage and total game percentage
func percentGraphColored(w io.Writer, p1name, p2name string, totalWidth int, p1, p2, done, total float64) {
	p1x := p1 / done
	p2x := p2 / done
	completed := done / total

	var totalPercent float64

	// compare header
	buff := bufio.NewWriter(w)
	buff.WriteString("\033[2J\033[0;0H")

	buff.WriteString("\033[0m")
	buff.WriteString(p1name)
	buff.WriteString(fmt.Sprint("[", int(p1x*100), "]"))

	// compare ratio graph
	for i := 0; i < totalWidth; i++ {
		totalPercent = float64(i) / float64(totalWidth)
		if totalPercent < p1x {
			buff.WriteString("\033[34m░")
		} else if totalPercent < p1x+p2x {
			buff.WriteString("\033[31m░")
		}
	}
	// compare footer
	buff.WriteString(fmt.Sprint("\033[0m[", int(p2x*100), "]"))
	buff.WriteString(p2name)
	buff.WriteString("\n")

	// total header
	buff.WriteString(fmt.Sprint("\033[0mC%[", int(completed*100), "]"))
	// total bar graph
	for i := 0; i < totalWidth; i++ {
		totalPercent = float64(i) / float64(totalWidth)
		if totalPercent < completed {
			buff.WriteString("\033[33m░")
		} else {
			buff.WriteString("\033[30m░")
		}
	}
	// total footer
	buff.WriteString(fmt.Sprint("\033[0m[", int(done), "]"))
	buff.WriteString("Round")
	buff.WriteString("\033[0m\n")

	buff.Flush()
}
