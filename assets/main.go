package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// for the white theme
type whiteTheme struct {
	fyne.Theme
}

func (w whiteTheme) TextColor() color.Color {
	return color.White
}

func (w whiteTheme) BackgroundColor() color.Color {
	return color.Black
}

func getDimensions(input string) (int, int) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(lines) == 0 || lines[0] == "" {
		return 0, 0
	}
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != width {
			return 0, 0
		}
	}
	return width, len(lines)
}

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	cmd := exec.Command("./quadA", fmt.Sprint(x), fmt.Sprint(y))
	output, err := cmd.CombinedOutput()
	if err != nil {
		// You can choose to handle the error differently if needed
		return ""
	}
	return strings.TrimRight(string(output), "\n")
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			switch {
			case i == 1 && j == 1:
				sb.WriteByte('/')
			case (i == 1 && j == x) || (i == y && j == 1):
				sb.WriteByte('\\')
			case i == y && j == x:
				sb.WriteByte('/')
			case i == 1 || i == y || j == 1 || j == x:
				sb.WriteByte('*')
			default:
				sb.WriteByte(' ')
			}
		}
		if i != y {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == 1 && j == x) {
				sb.WriteByte('A')
			} else if (i == y && j == 1) || (i == y && j == x) {
				sb.WriteByte('C')
			} else if i == 1 || i == y || j == 1 || j == x {
				sb.WriteByte('B')
			} else {
				sb.WriteByte(' ')
			}
		}
		if i != y {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == y && j == 1) {
				sb.WriteByte('A')
			} else if (i == 1 && j == x) || (i == y && j == x) {
				sb.WriteByte('C')
			} else if i == 1 || i == y || j == 1 || j == x {
				sb.WriteByte('B')
			} else {
				sb.WriteByte(' ')
			}
		}
		if i != y {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var sb strings.Builder
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				sb.WriteByte('A')
			} else if (i == 1 && j == x) || (i == y && j == 1) {
				sb.WriteByte('C')
			} else if i == y && j == x {
				sb.WriteByte('A')
			} else if i == 1 || i == y || j == 1 || j == x {
				sb.WriteByte('B')
			} else {
				sb.WriteByte(' ')
			}
		}
		if i != y {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func checkQuad(input string) string {
	width, height := getDimensions(input)
	if width == 0 || height == 0 {
		return "❌ Not a quad function ❌"
	}

	quads := []struct {
		name string
		fn   func(int, int) string
	}{
		{"quadA", quadA},
		{"quadB", quadB},
		{"quadC", quadC},
		{"quadD", quadD},
		{"quadE", quadE},
	}

	var matches []string
	for _, q := range quads {
		if q.fn(width, height) == input {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", q.name, width, height))
		}
	}
	sort.Strings(matches)
	if len(matches) == 0 {
		return "❌ Not a quad function ❌"
	}
	return strings.Join(matches, " || ")
}

func main() {
	audiopath := "./assats/pixify-230092.mp3"
	f, err := os.Open(audiopath)
	if err != nil {
		log.Println("Could not open music file:", err)
	} else {
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Println("Could not decode mp3:", err)
		} else {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			go func() {
				for {
					// Reset streamer to the beginning
					streamer.Seek(0)
					speaker.Play(streamer)
					// Wait for streamer to finish
					select {
					case <-time.After(time.Duration(float64(streamer.Len())/float64(format.SampleRate)) * time.Second):
					}
				}
			}()
		}
		// Don't close streamer or file here, let speaker use them
	}
	a := app.New()
	w := a.NewWindow("QUAD CHECKER")

	// 🎨 ASCII Logo (multi-line)
	logo := `
 ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓███████▓▒░        ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓███████▓▒░  
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░     ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░      ░▒▓███████▓▒░░▒▓██████▓▒░ ░▒▓███████▓▒░  
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░     ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
 ░▒▓██████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓███████▓▒░        ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░ 
 											𝗆𝖺𝖽𝖾 𝖻𝗒 𝗆𝗋𝖺𝗀𝖺𝖻 && 𝖸𝖾𝗅𝗍𝗎𝗐𝖺𝗂 
   ░▒▓█▓▒░                                                                                                                                             
	░▒▓██▓▒░                                                                                                                                      

`
	logoLines := strings.Split(logo, "\n")
	var logoTexts []*canvas.Text
	for _, line := range logoLines {
		lineText := canvas.NewText(line, theme.PrimaryColor())
		lineText.TextSize = 10
		lineText.TextStyle.Monospace = true
		logoTexts = append(logoTexts, lineText)
	}
	logoBox := container.NewVBox()
	for _, t := range logoTexts {
		logoBox.Add(t)
	}

	// Animation: move logo from left to right, then disappear letter by letter from the right, then reappear letter by letter from the left
	go func() {
		shift := 0
		maxShift := 30 // how far to move right
		logoLen := 0
		if len(logoLines) > 0 {
			for _, l := range logoLines {
				if len([]rune(l)) > logoLen {
					logoLen = len([]rune(l))
				}
			}
		}
		for {
			// Move right
			for shift = 0; shift <= maxShift; shift++ {
				for i, t := range logoTexts {
					orig := strings.TrimRight(logoLines[i], " ")
					t.Text = strings.Repeat(" ", shift) + orig
					canvas.Refresh(t)
				}
				time.Sleep(40 * time.Millisecond)
			}
			// Disappear letter by letter from right
			for cut := logoLen; cut >= 0; cut-- {
				for i, t := range logoTexts {
					orig := strings.TrimRight(logoLines[i], " ")
					runes := []rune(orig)
					n := len(runes)
					if cut > 0 {
						if cut > n {
							t.Text = strings.Repeat(" ", maxShift) + orig
						} else {
							t.Text = strings.Repeat(" ", maxShift) + string(runes[:cut])
						}
					} else {
						t.Text = ""
					}
					canvas.Refresh(t)
				}
				time.Sleep(50 * time.Millisecond)
			}
			// Appear letter by letter from left
			for cut := 1; cut <= logoLen; cut++ {
				for i, t := range logoTexts {
					orig := strings.TrimRight(logoLines[i], " ")
					runes := []rune(orig)
					n := len(runes)
					if cut > n {
						t.Text = strings.Repeat(" ", maxShift) + orig
					} else {
						t.Text = strings.Repeat(" ", maxShift) + string(runes[:cut])
					}
					canvas.Refresh(t)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	commandEntry := widget.NewEntry()
	commandEntry.SetPlaceHolder("Type a command like: ./quadA 3 3")

	outputBox := widget.NewMultiLineEntry()
	outputBox.SetPlaceHolder("Quad output will appear here...")
	outputBox.SetMinRowsVisible(10)
	resultBox := widget.NewMultiLineEntry()
	resultBox.SetPlaceHolder("Result will appear here...")
	resultBox.SetMinRowsVisible(5)

	quadFuncs := map[string]func(int, int) string{
		"quadA": quadA,
		"quadB": quadB,
		"quadC": quadC,
		"quadD": quadD,
		"quadE": quadE,
	}

	runCommand := func() {
		cmd := strings.TrimSpace(commandEntry.Text)
		fields := strings.Fields(cmd)
		if len(fields) != 3 || !strings.HasPrefix(fields[0], "./quad") {
			outputBox.SetText("❗ Input a quad example like ./quadA 3 3")
			resultBox.SetText("")
			return
		}
		quadName := strings.TrimPrefix(fields[0], "./")
		fn, ok := quadFuncs[quadName]
		if !ok {
			outputBox.SetText("⚠️ Unknown Quad Function" + quadName)
			resultBox.SetText("")
			return
		}
		var x, y int
		_, err1 := fmt.Sscan(fields[1], &x)
		_, err2 := fmt.Sscan(fields[2], &y)
		if err1 != nil || err2 != nil || x <= 0 || y <= 0 {
			outputBox.SetText("⚠️ Incorrect Dimensions.")
			resultBox.SetText("")
			return
		}
		quadOutput := fn(x, y)
		outputBox.SetText(quadOutput)
		result := checkQuad(quadOutput)
		resultBox.SetText(result)
	}

	runButton := widget.NewButton("Checker", runCommand)
	commandEntry.OnSubmitted = func(_ string) { runCommand() }

	content := container.NewVBox(
		logoBox,
		widget.NewSeparator(),
		commandEntry,
		runButton,
		widget.NewSeparator(),
		outputBox,
		widget.NewSeparator(),
		resultBox,
		layout.NewSpacer(),
	)

	w.SetContent(container.NewPadded(content))
	w.Resize(fyne.NewSize(700, 600))
	w.CenterOnScreen()
	w.ShowAndRun()
}
