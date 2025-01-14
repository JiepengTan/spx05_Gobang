// Code generated by gop (Go+); DO NOT EDIT.

package main

import "github.com/goplus/spx"

const _ = true
const (
	_Black         = 1
	_White         = 2
	_BlackAndWhite = _Black + _White
)

type Button4 struct {
	spx.Sprite
	*Game
}
type Calf1 struct {
	spx.Sprite
	*Game
}
type Chess struct {
	spx.Sprite
	*Game
}
type CurrentChess struct {
	spx.Sprite
	*Game
}
type Game struct {
	spx.Game
	Chess         Chess
	Button4       Button4
	CurrentChess  CurrentChess
	Calf1         Calf1
	gameState     spx.List
	ret           int
	col           int
	score         int
	currentCol    int
	currentPlayer int
	currentRow    int
	row           int
}
//line main.spx:22:1
func (this *Game) calcHitPosition(x float64, y float64) {
//line main.spx:23:1
	this.col = spx.Iround((x + 159) / 17.7)
//line main.spx:24:1
	this.row = spx.Iround((y - 159) / -17.7)
//line main.spx:25:1
	if this.row < 0 || this.col < 0 || this.row > 18 || this.col > 18 {
//line main.spx:26:1
		this.ret = 0
	} else {
//line main.spx:28:1
		this.ret = this.row*19 + this.col + 1
	}
}
//line main.spx:32:1
func (this *Game) initGameState() {
//line main.spx:33:1
	this.gameState.Delete(spx.All)
//line main.spx:34:1
	for
//line main.spx:34:1
	i := 0; i < 361;
//line main.spx:34:1
	i++ {
		spx.Sched()
//line main.spx:35:1
		this.gameState.Append(0)
	}
}
//line main.spx:39:1
func (this *Game) calcScore(dx int, dy int) {
//line main.spx:40:1
	this.row = this.currentRow + dy
//line main.spx:41:1
	this.col = this.currentCol + dx
//line main.spx:42:1
	for this.row > 0 && this.col > 0 && this.row < 18 && this.col < 18 && this.currentPlayer == this.gameState.At(this.row*19+this.col).Int() {
		spx.Sched()
//line main.spx:43:1
		this.score++
//line main.spx:44:1
		this.row += dy
//line main.spx:45:1
		this.col += dx
	}
}
//line main.spx:49:1
func (this *Game) checkHasWon() {
//line main.spx:50:1
	this.checkHasWonByRow(1, 1)
//line main.spx:51:1
	this.checkHasWonByRow(1, 0)
//line main.spx:52:1
	this.checkHasWonByRow(0, 1)
//line main.spx:53:1
	this.checkHasWonByRow(1, -1)
}
//line main.spx:56:1
func (this *Game) checkHasWonByRow(dx int, dy int) {
//line main.spx:57:1
	this.score = 1
//line main.spx:58:1
	this.calcScore(dx, dy)
//line main.spx:59:1
	this.calcScore(-dx, -dy)
//line main.spx:60:1
	if this.score > 4 {
//line main.spx:61:1
		this.Broadcast__1("game over", true)
	}
}
//line main.spx:65
func (this *Game) MainEntry() {
//line main.spx:65:1
	this.OnStart(func() {
//line main.spx:66:1
		this.initGameState()
//line main.spx:67:1
		this.score = 0
//line main.spx:68:1
		this.currentPlayer = _Black
	})
//line main.spx:71:1
	this.OnClick(func() {
//line main.spx:72:1
		this.calcHitPosition(this.MouseX(), this.MouseY())
//line main.spx:73:1
		if this.ret > 0 && this.gameState.At(this.ret-1).Int() == 0 {
//line main.spx:74:1
			this.currentCol = this.col
//line main.spx:75:1
			this.currentRow = this.row
//line main.spx:76:1
			this.Broadcast__0("try put chess")
		}
	})
//line main.spx:80:1
	this.OnMsg__1("confirm to put chess", func() {
//line main.spx:81:1
		if this.score < 5 {
//line main.spx:82:1
			this.gameState.Set(this.currentRow*19+this.currentCol, this.currentPlayer)
//line main.spx:83:1
			this.checkHasWon()
//line main.spx:84:1
			spx.Gopt_Sprite_Clone__0(&this.Chess)
//line main.spx:85:1
			this.Broadcast__1("put chess done", true)
		}
	})
}
func (this *Game) Main() {
	spx.Gopt_Game_Main(this, new(Button4), new(Calf1), new(Chess), new(CurrentChess))
}
//line Button4.spx:1
func (this *Button4) Main() {
//line Button4.spx:1:1
	this.OnStart(func() {
//line Button4.spx:2:1
		this.Hide()
	})
//line Button4.spx:5:1
	this.OnClick(func() {
//line Button4.spx:6:1
		this.Broadcast__0("confirm to put chess")
//line Button4.spx:7:1
		this.Hide()
	})
//line Button4.spx:10:1
	this.OnMsg__1("try put chess", func() {
//line Button4.spx:11:1
		this.Show()
	})
}
func (this *Button4) Classfname() string {
	return "Button4"
}
//line Calf1.spx:1
func (this *Calf1) Main() {
//line Calf1.spx:1:1
	this.OnStart(func() {
//line Calf1.spx:2:1
		this.Hide()
	})
//line Calf1.spx:5:1
	this.OnMsg__1("game over", func() {
//line Calf1.spx:6:1
		this.Show()
//line Calf1.spx:7:1
		if this.currentPlayer == _Black {
//line Calf1.spx:8:1
			this.Say("The black side won!", 3)
		} else {
//line Calf1.spx:10:1
			this.Say("The white side won!", 3)
		}
//line Calf1.spx:12:1
		spx.Exit__1()
	})
}
func (this *Calf1) Classfname() string {
	return "Calf1"
}
//line Chess.spx:1
func (this *Chess) Main() {
//line Chess.spx:1:1
	this.OnCloned__1(func() {
//line Chess.spx:2:1
		this.SetXYpos(float64(this.currentCol)*17.7-159, float64(this.currentRow)*-17.7+159)
//line Chess.spx:3:1
		this.SetCostume(this.currentPlayer - 1)
//line Chess.spx:4:1
		this.currentPlayer = _BlackAndWhite - this.currentPlayer
//line Chess.spx:5:1
		this.Show()
	})
}
func (this *Chess) Classfname() string {
	return "Chess"
}
//line CurrentChess.spx:1
func (this *CurrentChess) Main() {
//line CurrentChess.spx:1:1
	this.OnStart(func() {
//line CurrentChess.spx:2:1
		this.Hide()
	})
//line CurrentChess.spx:5:1
	this.OnMsg__1("put chess done", func() {
//line CurrentChess.spx:6:1
		this.Hide()
	})
//line CurrentChess.spx:9:1
	this.OnMsg__1("try put chess", func() {
//line CurrentChess.spx:10:1
		this.SetXYpos(float64(this.currentCol)*17.7-159, float64(this.currentRow)*-17.7+159)
//line CurrentChess.spx:11:1
		this.SetCostume(this.currentPlayer - 1)
//line CurrentChess.spx:12:1
		this.Show()
	})
}
func (this *CurrentChess) Classfname() string {
	return "CurrentChess"
}
func main() {
	new(Game).Main()
}
