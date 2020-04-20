package handlers

import (
	"chitchat/models"
	"fmt"
	"net/http"
)

// POST /thread/post
// 在指定群组下创建新主题
func PostThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("Cannot get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := models.ThreadByUuid(uuid)
		if err != nil {
			error_message(w, r, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			fmt.Println(err)
			fmt.Println("Cannot create post")
		}
		url := fmt.Sprintf("/thread/read?id=", uuid)
		http.Redirect(w, r, url, 302)
	}
}
