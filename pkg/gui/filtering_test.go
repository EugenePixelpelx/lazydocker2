package gui

import (
	"testing"

	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazydocker/pkg/gui/panels"
	"github.com/stretchr/testify/assert"
)

func TestMainSearchDoesNotFilterSidePanels(t *testing.T) {
	gui := &Gui{
		State: guiState{
			Filter: filterState{
				needle:     "usage",
				searchView: &gocui.View{},
			},
		},
	}

	assert.Empty(t, gui.FilterString(&gocui.View{}))
}

func TestFilterTargetLabelUsesSidePanelTitle(t *testing.T) {
	containersView := &gocui.View{Title: "Containers"}
	containersPanel := &panels.SideListPanel[int]{
		ListPanel: panels.ListPanel[int]{View: containersView},
	}
	gui := &Gui{
		State: guiState{
			Filter: filterState{active: true, panel: containersPanel},
		},
	}

	assert.Equal(t, "Containers", gui.filterTargetLabel())
}

func TestFilterTargetLabelUsesCurrentMainTab(t *testing.T) {
	mainView := &gocui.View{Tabs: []string{"Logs", "Stats", "Env"}, TabIndex: 1}
	gui := &Gui{
		State: guiState{
			Filter: filterState{active: true, searchView: mainView},
		},
		Views: Views{Main: mainView},
	}

	assert.Equal(t, "Stats", gui.filterTargetLabel())
}
