package models

import (
	"database/sql"
	"time"

	"example.com/eventbook/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int64     `json:"userId"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events WHERE id=$1`
	var event Event
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &event, nil
}

func (e *Event) Update() error {
	query := `UPDATE events SET name=$1, description=$2, location=$3, dateTime=$4, user_id=$5 WHERE id=$6`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.ID)
	return err
}

func DeleteEventByID(id int64) error {
	query := `DELETE FROM events WHERE id=$1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func (e *User) Register() error {
	query := "INSERT INTO registration(event_id,user_id) VALUES ($1,$2)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, )
	return err
}

func (e Event) Deregister(userId int64) error {
	query := "DELETE FROM registration WHERE event_id=? AND user_id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, e.UserId)
	return err
}

func GetRegistrationByEventID(eventId int64) ([]User, error) {
	query := `SELECT user_id FROM registration WHERE event_id=$1`
	rows, err := db.DB.Query(query, eventId)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()
	var registeredUsers []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
		registeredUsers = append(registeredUsers, user)
	}
	return registeredUsers, nil
}
