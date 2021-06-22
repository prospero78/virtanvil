// Package gva -- главный класс приложения
package gva

import (
	"github.com/prospero78/virtanvil/internal/webserver"
	"github.com/sirupsen/logrus"
)

// TGoVirtAnvil -- операции с приложением
type TGoVirtAnvil struct {
	webServer *webserver.TWebServer
}

// New -- возвращает новый *TGoVirtAnvil
func New() *TGoVirtAnvil {
	return &TGoVirtAnvil{
		webServer: webserver.New(),
	}
}

// Run -- запускает приложение в работу
func (sf *TGoVirtAnvil) Run() {
	logrus.Debug("TGoVirtAnvil.Run()")
	sf.webServer.Run()
}
