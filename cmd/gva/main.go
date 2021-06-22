// Package main -- пускач для сервера
package main

import (
	"github.com/prospero78/virtanvil/internal/gva"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Infoln("main()")
	app := gva.New()
	app.Run()
}
