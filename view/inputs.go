package view

import (
	"github.com/gdamore/tcell/v2"
)

//	func (v *View) InputMounts(event *tcell.EventKey) *tcell.EventKey {
//		event = v.InputMainCommands(event)
//		return v.inputMounts(event)
//	}
func (v *View) InputMounts(event *tcell.EventKey) *tcell.EventKey {
	event = v.InputMainCommands(event)
	return v.inputMounts(event)
}

func (v *View) InputVaultPolicy(event *tcell.EventKey) *tcell.EventKey {
	event = v.InputMainCommands(event)
	return v.inputPolicy(event)
}

func (v *View) InputSecrets(event *tcell.EventKey) *tcell.EventKey {
	event = v.InputMainCommands(event)
	return v.inputSecrets(event)
}

func (v *View) InputSecret(event *tcell.EventKey) *tcell.EventKey {
	event = v.InputMainCommands(event)
	return v.inputSecret(event)
}

func (v *View) InputMainCommands(event *tcell.EventKey) *tcell.EventKey {
	if event == nil {
		return event
	}
	switch event.Key() {
	// Bug: CTRL+M key maps to enter and causes conflicts
	case tcell.KeyCtrlB:
		v.Watcher.Unsubscribe()
		v.Mounts()
	case tcell.KeyCtrlP:
		v.VPolicy()
		// Needs editing
		// case tcell.KeyCtrlJ:
		// 	v.SecretObject()
	}

	return event
}
