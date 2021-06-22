package index

/*
	Обработчик главной страницы сайта
*/

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

type TIndex struct {
	*http.ServeMux
}

func New(router *http.ServeMux) *TIndex {
	page := &TIndex{
		ServeMux: http.NewServeMux(),
	}
	router.Handle("/", page)
	router.Handle("/index", page)
	router.Handle("/index/", page)
	return &TIndex{}
}

// Обработчик домашней страницы
func (sf *TIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
		w.Write([]byte("TIndex.ServeHTTP(): разрешён только GET-метод\n"))
		return
	}
	sf.get(w, r)
}

// Обрабатывает GET-запрос
func (sf *TIndex) get(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	switch r.URL.Path {
	case "/", "/index", "/index/":
		// Используем функцию template.ParseFiles() для чтения файла шаблона.
		// Если возникла ошибка, мы запишем детальное сообщение ошибки и
		// используя функцию http.Error() мы отправим пользователю
		// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
		var (
			ts    *template.Template
			err   error
			files []string
		)
		if logrus.GetLevel() == logrus.DebugLevel {
			// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
			// файл tmpl-index.html должен быть *первым* файлом в срезе.
			files = []string{
				"../src/html/tmpl-index.html",
				"../src/html/tmpl-base.html",
				"../src/html/tmpl-footer.html",
			}
			ts, err = template.ParseFiles(files...)
		}
		if err != nil {
			logrus.WithError(err).Errorln("TIndex.get(): при загрузке домашней страницы")
			w.WriteHeader(405)
			fmt.Fprintf(w, "TIndex.get(): при загрузке домашней страницы, err=%v\n", err)
			return
		}
		// Затем мы используем метод Execute() для записи содержимого
		// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
		// возможность отправки динамических данных в шаблон.
		err = ts.Execute(w, nil)
		if err != nil {
			logrus.WithError(err).Errorln("TIndex.get(): при обработке шаблона домашней страницы")
			w.WriteHeader(405)
			fmt.Fprintf(w, "TIndex.get(): при обработке шаблона домашней страницы, err=%v\n", err)
			return
		}
	default:
		logrus.WithField("path", r.URL.Path).Errorln("TIndex.get(): не найден путь домашней страницы")
		w.WriteHeader(404)
		fmt.Fprintf(w, "TIndex.get(): не найден путь домашней страницы, path=%v\n", r.URL.Path)
	}
}
