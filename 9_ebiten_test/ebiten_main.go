// package main

// import (
// 	"image/color"
// 	"log"
// 	"strconv"

// 	"github.com/hajimehoshi/ebiten/v2"            //ebiten本体
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil" //ebiten工具集
// )

// type Game struct {
// 	i uint8
// }

// func (g *Game) Update() error {
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "HOUZW!") //在屏幕上输出
// }

// func Hex2RGB(color16 string, alpha uint8) color.RGBA {
// 	r, _ := strconv.ParseInt(color16[:2], 16, 10)
// 	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
// 	b, _ := strconv.ParseInt(color16[4:], 16, 10)
// 	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 320, 240 //窗口分辨率
// }

// func main() {
// 	ebiten.SetWindowSize(640, 480)         //窗口大小
// 	ebiten.SetWindowTitle("Hello, World!") //窗口标题
// 	if err := ebiten.RunGame(&Game{}); err != nil {
// 		log.Fatal(err)
// 	}
// }
