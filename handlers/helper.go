package handlers

import (
	"chitchat/models"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

// 通过 Cookie 判断用户是否已登录
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 解析 HTML 模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成响应 HTML
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)
}

func Version() string {
	return "0.1"
}

func info(args ...interface{})  {
	logger.SetPrefix("INFO ")
	logger.Println(args)
}

func danger(args ...interface{})  {
	logger.SetPrefix("ERROR ")
	logger.Println(args)
}

func warning(args ...interface{})  {
	logger.SetPrefix("WARNING")
	logger.Println(args)
}
