package gui

import (
	"testing"

	"github.com/jesseduffield/gocui"
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
