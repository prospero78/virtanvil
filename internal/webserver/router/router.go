package router

import (
	"net/http"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/virtanvil/internal/webserver/router/index"
	"github.com/prospero78/virtanvil/internal/webserver/router/snippet"
)

/*
	Пакет  группирует в себе обработчики для различных страниц
*/

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}

type TRouter struct {
	*http.ServeMux
	pIndex   *index.TIndex
	pSnippet *snippet.TSnippet
}

func New() *TRouter {
	router := &TRouter{
		ServeMux: http.NewServeMux(),
	}
	router.pIndex = index.New(router.ServeMux)
	router.pSnippet = snippet.New(router.ServeMux)

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	if logrus.GetLevel() == logrus.DebugLevel {
		fileServer := http.FileServer(neuteredFileSystem{http.Dir("../src")})
		// Используем функцию mux.Handle() для регистрации обработчика для
		// всех запросов, которые начинаются с "/static/". Мы убираем
		// префикс "/static" перед тем как запрос достигнет http.FileServer
		router.Handle("/static", http.NotFoundHandler())
		router.Handle("/static/", http.StripPrefix("/static", fileServer))
	}

	return router
}
