package main

import (
	"image"
	"math/rand"

	"gocv.io/x/gocv"
	cv2 "gocv.io/x/gocv"
)

type SnakeNode struct {
	x    int
	y    int
	next *SnakeNode
}

func grow(root *SnakeNode, x int, y int) {
	if root.next == nil {
		var kid SnakeNode
		kid.x = x
		kid.y = y
		kid.next = nil
		root.next = &kid
	} else {
		grow(root.next, x, y)
	}
}

func move(root *SnakeNode, x int, y int) {
	x0 := root.x
	y0 := root.y
	root.x = x
	root.y = y
	if root.next != nil {
		move(root.next, x0, y0)
	}
}

func get_len(root *SnakeNode) int {
	if root == nil {
		return 0
	} else {
		return 1 + get_len(root.next)
	}
}
func draw(root *SnakeNode, img *cv2.Mat, food *image.Point) {
	if root != nil {
		data := img.DataPtrUint8()
		w := img.Cols()
		x := root.x
		y := root.y
		if x >= 0 && y >= 0 {
			if data[y*w+x] == 127 {
				food.X = -1
				food.Y = -1
				grow(root, -1, -1)
				*food = give(img)
			}
			data[y*w+x] = 255
		}
		if root.next != nil {
			draw(root.next, img, food)
		}
	}
}

func draw_food(food image.Point, img *cv2.Mat) {
	data := img.DataPtrUint8()
	w := img.Cols()
	x := food.X
	y := food.Y
	if x >= 0 && y >= 0 {
		data[y*w+x] = 127
	}
}

func give(img *cv2.Mat) image.Point {
	rr := rand.Intn(img.Rows())
	cc := rand.Intn(img.Cols())
	var food image.Point
	food.X = cc
	food.Y = rr
	return food
}
func main() {
	window := gocv.NewWindow("Hello")
	var h int = 64
	var w int = 64
	img := cv2.NewMatWithSize(h, w, cv2.MatTypeCV8UC1)
	show := cv2.NewMatWithSize(512, 512, cv2.MatTypeCV8UC1)
	var sz image.Point
	sz.X = 512
	sz.Y = 512
	var root SnakeNode
	var snakeSpeed int = int('d')
	root.x = 10
	root.y = 10
	root.next = nil
	grow(&root, 9, 10)
	grow(&root, 8, 10)
	grow(&root, 7, 10)

	food := give(&img)
	draw_food(food, &img)
	draw(&root, &img, &food)
	for {

		gocv.Resize(img, &show, sz, 0, 0, 0)
		window.IMShow(show)
		var cmd int = window.WaitKey(200)
		var ret byte = 'q'
		// if cmd >= 0 {
		// 	fmt.Println(cmd)
		// }
		if cmd == int('d') || cmd == int('a') || cmd == int('w') || cmd == int('s') {
			snakeSpeed = cmd
		}
		// n := get_len(&root)
		// fmt.Println("len:", n)
		if snakeSpeed == int('d') {
			move(&root, root.x+1, root.y)
			img = cv2.NewMatWithSize(h, w, cv2.MatTypeCV8UC1)
			draw_food(food, &img)
			draw(&root, &img, &food)
		}
		if snakeSpeed == int('a') {
			move(&root, root.x-1, root.y)
			img = cv2.NewMatWithSize(h, w, cv2.MatTypeCV8UC1)
			draw_food(food, &img)
			draw(&root, &img, &food)
		}
		if snakeSpeed == int('w') {
			move(&root, root.x, root.y-1)
			img = cv2.NewMatWithSize(h, w, cv2.MatTypeCV8UC1)
			draw_food(food, &img)
			draw(&root, &img, &food)
		}
		if snakeSpeed == int('s') {
			move(&root, root.x, root.y+1)
			img = cv2.NewMatWithSize(h, w, cv2.MatTypeCV8UC1)
			draw_food(food, &img)
			draw(&root, &img, &food)
		}
		if cmd == int(ret) {
			break
		}
	}
}
