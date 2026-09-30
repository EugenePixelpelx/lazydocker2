package gui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
)

func (gui *Gui) handleOpenFilter() error {
	panel, ok := gui.currentListPanel()
	if !ok {
		return nil
	}

	if panel.IsFilterDisabled() {
		return nil
	}

	gui.State.Filter.active = true
	gui.State.Filter.panel = panel

	if err := gui.updateFilterPrompt(); err != nil {
		return err
	}

	return gui.switchFocus(gui.Views.Filter)
}

func (gui *Gui) handleOpenMainSearch() error {
	gui.State.Filter.active = true
	gui.State.Filter.panel = nil
	gui.State.Filter.searchView = gui.Views.Main
	gui.State.Filter.needle = ""
	gui.Views.Filter.ClearTextArea()

	if err := gui.updateFilterPrompt(); err != nil {
		return err
	}

	return gui.switchFocus(gui.Views.Filter)
}

func (gui *Gui) onNewFilterNeedle(value string) error {
	gui.State.Filter.needle = value

	if gui.State.Filter.searchView != nil {
		if value == "" {
			gui.State.Filter.searchView.ClearSearch()
			return gui.updateFilterPrompt()
		}

		return gui.State.Filter.searchView.Search(value)
	}

	gui.ResetOrigin(gui.State.Filter.panel.GetView())
	return gui.State.Filter.panel.RerenderList()
}

func (gui *Gui) wrapEditor(f func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool) func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
	return func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
		matched := f(v, key, ch, mod)
		if matched {
			if err := gui.onNewFilterNeedle(v.TextArea.GetContent()); err != nil {
				gui.Log.Error(err)
			}
		}
		return matched
	}
}

func (gui *Gui) escapeFilterPrompt() error {
	if err := gui.clearFilter(); err != nil {
		return err
	}

	return gui.returnFocus()
}

func (gui *Gui) clearFilter() error {
	if gui.State.Filter.searchView != nil {
		return gui.clearMainSearch()
	}

	gui.State.Filter.needle = ""
	gui.State.Filter.active = false
	panel := gui.State.Filter.panel
	gui.State.Filter.panel = nil
	gui.Views.Filter.ClearTextArea()

	if panel == nil {
		return nil
	}

	gui.ResetOrigin(panel.GetView())

	return panel.RerenderList()
}

func (gui *Gui) clearMainSearch() error {
	if gui.State.Filter.searchView == nil {
		return nil
	}

	gui.State.Filter.searchView.ClearSearch()
	gui.State.Filter.searchView = nil
	gui.State.Filter.needle = ""
	gui.State.Filter.active = false
	gui.Views.Filter.ClearTextArea()

	return gui.updateFilterPrompt()
}

func (gui *Gui) onMainSearchResult(_ int, _ int, _ int) error {
	return gui.updateFilterPrompt()
}

func (gui *Gui) selectNextMainSearchResult() error {
	current, total := gui.Views.Main.GetSearchStatus()
	if total == 0 {
		return nil
	}

	return gui.Views.Main.SelectSearchResult((current + 1) % total)
}

func (gui *Gui) updateFilterPrompt() error {
	return gui.setViewContent(gui.Views.FilterPrefix, gui.filterPrompt())
}

func (gui *Gui) filterTargetView() *gocui.View {
	if !gui.State.Filter.active {
		return nil
	}

	if gui.State.Filter.searchView != nil {
		return gui.State.Filter.searchView
	}

	if gui.State.Filter.panel != nil {
		return gui.State.Filter.panel.GetView()
	}

	return nil
}

func (gui *Gui) filterTargetLabel() string {
	targetView := gui.filterTargetView()
	if targetView == nil {
		return ""
	}

	if targetView == gui.Views.Main && targetView.TabIndex >= 0 && targetView.TabIndex < len(targetView.Tabs) {
		return targetView.Tabs[targetView.TabIndex]
	}

	if targetView.Title != "" {
		return targetView.Title
	}

	return targetView.Name()
}

// returns to the list view with the filter still applied
func (gui *Gui) commitFilter() error {
	if gui.State.Filter.needle == "" {
		if err := gui.clearFilter(); err != nil {
			return err
		}
	}

	return gui.returnFocus()
}

func (gui *Gui) filterPrompt() string {
	targetLabel := gui.filterTargetLabel()

	if gui.State.Filter.searchView != nil {
		current, total := gui.State.Filter.searchView.GetSearchStatus()
		if gui.State.Filter.needle != "" {
			if total == 0 {
				return fmt.Sprintf("%s %s (0/0): ", gui.Tr.SearchPrompt, targetLabel)
			}

			return fmt.Sprintf("%s %s (%d/%d): ", gui.Tr.SearchPrompt, targetLabel, current+1, total)
		}

		return fmt.Sprintf("%s %s: ", gui.Tr.SearchPrompt, targetLabel)
	}

	if targetLabel == "" {
		return fmt.Sprintf("%s: ", gui.Tr.FilterPrompt)
	}

	return fmt.Sprintf("%s %s: ", gui.Tr.FilterPrompt, targetLabel)
}
