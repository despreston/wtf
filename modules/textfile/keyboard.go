package textfile

import (
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(nil)

	widget.SetKeyboardChar("l", widget.NextSource, "Select next file")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous file")
	widget.SetKeyboardChar("o", widget.openFile, "Open file")

	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next file")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous file")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openFile, "Open file")
}

func (widget *Widget) openFile() {
	src := widget.CurrentSource()
	widget.tviewApp.Suspend(func() {
		cmd := exec.Command("nvim", src)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	})
}
