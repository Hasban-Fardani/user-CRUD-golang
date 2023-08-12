package views

import (
	"log"
	"os"
)

func InitViews() {
	dir, _ := os.Getwd()
	log.Println("dir:", dir)
}
