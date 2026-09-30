package gui

import (
	"testing"

	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazydocker/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestShouldHighlightView(t *testing.T) {
	mainView := &gocui.View{}
	filterView := &gocui.View{}
	containersView := &gocui.View{}
	imagesView := &gocui.View{}
	gui := &Gui{Views: Views{Main: mainView, Filter: filterView}}

	assert.True(t, gui.shouldHighlightView(containersView, containersView))
	assert.False(t, gui.shouldHighlightView(mainView, mainView))
	assert.False(t, gui.shouldHighlightView(containersView, mainView))

	mainView.ParentView = containersView

	assert.True(t, gui.shouldHighlightView(containersView, mainView))
	assert.False(t, gui.shouldHighlightView(imagesView, mainView))
	assert.False(t, gui.shouldHighlightView(containersView, filterView))

	gui.State.Filter.searchView = mainView
	gui.State.Filter.active = true

	assert.True(t, gui.shouldHighlightView(containersView, filterView))
	assert.False(t, gui.shouldHighlightView(imagesView, filterView))
}

func TestMainSearchKeepsMainPanelFullscreen(t *testing.T) {
	mainView := &gocui.View{}
	userConfig := config.GetDefaultConfig()
	gui := &Gui{
		Config: &config.AppConfig{UserConfig: &userConfig},
		State: guiState{
			ViewStack:  []string{"containers"},
			ScreenMode: SCREEN_FULL,
			Filter: filterState{
				active:     true,
				searchView: mainView,
			},
		},
		Views: Views{Main: mainView},
	}

	sideWeight, mainWeight := gui.getMidSectionWeights()

	assert.Equal(t, 0, sideWeight)
	assert.Equal(t, 1, mainWeight)
}

func TestToggleMainPanelFullscreenRestoresPreviousMode(t *testing.T) {
	gui := &Gui{State: guiState{ScreenMode: SCREEN_HALF}}

	assert.NoError(t, gui.toggleMainPanelFullscreen())
	assert.Equal(t, SCREEN_FULL, gui.State.ScreenMode)
	assert.True(t, gui.State.MainPanelFullscreen)

	assert.NoError(t, gui.toggleMainPanelFullscreen())
	assert.Equal(t, SCREEN_HALF, gui.State.ScreenMode)
	assert.False(t, gui.State.MainPanelFullscreen)
}

func TestToggleMainPanelFullscreenLeavesLegacyFullscreen(t *testing.T) {
	gui := &Gui{State: guiState{ScreenMode: SCREEN_FULL}}

	assert.NoError(t, gui.toggleMainPanelFullscreen())
	assert.Equal(t, SCREEN_NORMAL, gui.State.ScreenMode)
	assert.False(t, gui.State.MainPanelFullscreen)
}

func TestEscapeLeavesMainPanelFullscreenWithoutReturningFocus(t *testing.T) {
	mainView := &gocui.View{}
	parentView := &gocui.View{}
	mainView.ParentView = parentView
	gui := &Gui{
		State: guiState{
			ScreenMode:          SCREEN_FULL,
			MainPanelFullscreen: true,
			PreviousScreenMode:  SCREEN_HALF,
		},
	}

	assert.NoError(t, gui.handleExitMain(nil, mainView))
	assert.Equal(t, SCREEN_HALF, gui.State.ScreenMode)
	assert.False(t, gui.State.MainPanelFullscreen)
	assert.Same(t, parentView, mainView.ParentView)
}

func TestTabDoesNothingInMainPanelFullscreen(t *testing.T) {
	gui := &Gui{State: guiState{MainPanelFullscreen: true}}

	assert.NoError(t, gui.handleNextViewFromMain(nil, &gocui.View{}))
	assert.NoError(t, gui.handlePreviousViewFromMain(nil, &gocui.View{}))
}
