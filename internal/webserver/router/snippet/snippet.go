package snippet

/*
	Обработчик сниппетов сайта
*/

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type TSnippet struct {
	*http.ServeMux
}

func New(router *http.ServeMux) *TSnippet {
	page := &TSnippet{
		ServeMux: http.NewServeMux(),
	}
	router.Handle("/snippet", page)
	router.Handle("/snippet/", page)
	return &TSnippet{}
}

// Обработчик домашней страницы
func (sf *TSnippet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
		w.Write([]byte("TSnippet.ServeHTTP(): разрешён только GET-метод\n"))
		return
	}
	sf.get(w, r)
}

// Обрабатывает GET-запрос
func (sf *TSnippet) get(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	switch r.URL.Path {
	case "/snippet", "/snippet/":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			w.WriteHeader(404)
			err := fmt.Sprintf("TSnippet.get(): id=%v, err=%v\n", id, err)
			w.Write([]byte(err))
			logrus.WithField("err", err).Errorln("TSnippet.get()")
			return
		}
		msg := fmt.Sprintf("TSnippet.get(): id=%v\n", id)
		w.Write([]byte(msg))
	default:
		logrus.WithField("path", r.URL.Path).Errorln("TSnippet.get(): не найден путь домашней страницы")
		w.WriteHeader(404)
		fmt.Fprintf(w, "TSnippet.get(): не найден путь домашней страницы, path=%v\n", r.URL.Path)
	}
}
