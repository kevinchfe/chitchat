package models

import (
	db "chitchat/databases"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	sql := "insert into sessions (uuid,email,user_id,created_at) values (?,?,?,?)"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	uuid := createUUID()
	stmt.Exec(uuid, user.Email, user.Id, time.Now())

	stmtu, err := db.Db.Prepare("select id,uuid,email,user_id,created_at from sessions where uuid=?")
	if err != nil {
		return
	}
	defer stmtu.Close()

	err = stmtu.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = db.Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where user_id=?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	sql := "insert into users (uuid,name,email,password,created_at) values (?,?,?,?,?)"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	stmt.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())
	stmtu, err := db.Db.Prepare("select id,uuid,name,email,password,created_at from users where uuid=?")
	if err != nil {
		return
	}
	defer stmtu.Close()
	err = stmtu.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// delete user
func (user *User) Delete() (err error) {
	sql := "delete from users where id=?"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// update user
func (user *User) Update() (err error) {
	sql := "update users set name=?, email=? where id=?"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return
}

// delete all users
func (user *User) UserDeleteAll() (err error) {
	sql := "delete from users"
	_, err = db.Db.Exec(sql)
	return
}

// get all users
func Users() (users []User, err error) {
	rows, err := db.Db.Query("select id,uuid,name,email,password,created_at from users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// get user by email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = db.Db.QueryRow("select id,uuid,name,email,password,created_at from users where email=?", email).Scan(&user.Id,
		&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// get user by uuid
func UserByUuid(uuid string) (user User, err error) {
	user = User{}
	err = db.Db.QueryRow("select id,uuid,name,email,password,created_at where uuid=?", uuid).Scan(&user.Id,
		&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (thread Thread, err error) {
	sql := "insert into threads (uuid,topic,user_id,created_at) values (?,?,?,?)"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	stmt.Exec(uuid, topic, user.Id, time.Now())

	stmtt, err := db.Db.Prepare("select id,uuid,topic,user_id,created_at from threads where uuid=?")
	if err != nil {
		return
	}
	defer stmtt.Close()

	err = stmtt.QueryRow(uuid).Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

// Create a new post to a thread
func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	sql := "insert into posts (uuid,body,user_id,thread_id,created_at) values (?,?,?,?,?)"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	stmt.Exec(uuid, body, user.Id, thread.Id, time.Now())

	stmtp, err := db.Db.Prepare("select id,uuid,body,user_id,thread_id,created_at from posts where uuid=?")
	if err != nil {
		return
	}
	defer stmtp.Close()

	err = stmtp.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)

	return
}
