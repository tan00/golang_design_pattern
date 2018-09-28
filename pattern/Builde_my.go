package pattern

import (
	"fmt"
	"time"
)

func BuilderClient() {

	fullMode := NewFullMode()
	simpleMode := NewSimpleMode()

	player := new(tvideoPlaer)

	player.SetMode(fullMode)
	player.Play()

	player.SetMode(simpleMode)
	player.Play()

	fmt.Println("call BuilderClient")
}

type playMode interface {
	showMode()
}

type tvideoPlaer struct {
	mode playMode
}

func (vp *tvideoPlaer) SetMode(ins playMode) {
	vp.mode = ins
}

func (vp *tvideoPlaer) Play() {
	vp.mode.showMode()
}

//mode of videoPlaer //
type tplayMode struct {
	playList   []string
	mainWindow string
	controlBar time.Time
}

func (pm *tplayMode) showMode() {
	fmt.Printf("playList=%v mainWindow=%v controlBar=%v \n", pm.playList, pm.mainWindow, pm.controlBar)
}

type fullMode struct {
	tplayMode
}
type simpleMode struct {
	tplayMode
}

func NewFullMode() playMode {
	newmode := new(fullMode)
	newmode.playList = make([]string, 1)
	newmode.playList = append(newmode.playList, "empty list")

	newmode.mainWindow = "full mode mainwindow"
	newmode.controlBar = time.Now()
	return newmode
}

func NewSimpleMode() playMode {
	newmode := new(simpleMode)
	newmode.playList = make([]string, 1)
	newmode.playList = append(newmode.playList, "empty list")

	newmode.mainWindow = "simple mode mainwindow"
	newmode.controlBar = time.Now()
	return newmode
}
