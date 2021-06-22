// Package webserver -- главный тип веб-сервера
package webserver

import (
	"net/http"
	

	"github.com/sirupsen/logrus"

	"github.com/prospero78/virtanvil/internal/webserver/router"
)

// TWebServer -- операци ис веб-сервером
type TWebServer struct {
	router  *http.ServeMux
	router2 *router.TRouter
}

// New -- возвращает новый *TWebServer
func New() *TWebServer {
	srv := &TWebServer{
		router:  http.NewServeMux(),
		router2: router.New(),
	}
	srv.router.HandleFunc("/snippet/create", srv._createSnippet)

	

	return srv
}

// Run -- запускает веб-сервер в работу
func (sf *TWebServer) Run() {
	strAdr := "localhost:8200"
	logrus.WithField("adr", strAdr).Debug("TWebServer.Run()")
	err := http.ListenAndServe(strAdr, sf.router2)
	logrus.WithError(err).Panicln("TWebServer.Run(): ошибка при работе сервера")
}

// Обработчик для создания новой заметки.
func (sf *TWebServer) _createSnippet(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
	// что http.MethodPost является строкой и содержит текст "POST".
	if r.Method != http.MethodPost {
		// Если это не так, то вызывается метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".

		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.

		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		//      // Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением.
		//      // http.Error(w, "Метод запрещен!", 405)
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!\n"))
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}
