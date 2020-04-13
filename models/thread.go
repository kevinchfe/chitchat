package models

import (
	db "chitchat/databases"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := db.Db.Query("select count(*) from posts where thread_id=?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := db.Db.Query("select id,uuid,body,user_id,thread_id,created_at from posts where thread_id=?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get all threads
func Threads() (threads []Thread, err error) {
	rows, err := db.Db.Query("select id,uuid,body,user_id,thread_id,created_at from threads order by created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Id, &thread.UserId, &thread.Topic, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, thread)
	}
	rows.Close()
	return
}

// Get a thread by the UUID
func ThreadByUuid(uuid string) (thread Thread, err error) {
	thread = Thread{}
	db.Db.QueryRow("select id,uuid,topic,user_id,created_at from threads where uuid=?", uuid).Scan(&thread.Id,
		&thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	db.Db.QueryRow("select id,uuid,name,email,created_at from users where user_id=?", thread.UserId).Scan(&user.Id,
		&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}

