package models

import (
	db "chitchat/databases"
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// Check if session is valid
func (session *Session) Check() (valid bool, err error) {
	err = db.Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where uuid=?", session.Uuid).Scan(&session.Id,
		&session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// delete session by uuid
func (session *Session) DeleteByUuid() (err error) {
	sql := "delete from sessions where uuid=?"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	return
}

// get user from session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = db.Db.QueryRow("select id,uuid,name,email,created_at from users where id=?", session.UserId).Scan(&user.Id,
		&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// delete all sessions
func SessionDeleteAll() (err error) {
	sql := "delete from sessions"
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(stmt)
	return
}
