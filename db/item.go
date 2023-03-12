package db

import (
	"database/sql"
	"time"

	"github.com/rizperdana/prototype-golang-rest/models"
)

func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC;")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

func (db Database) AddItem(item *models.Item) error {
	var id int
	var createdAt string
	var updatedAt string
	query := `INSERT INTO item (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at;`
	err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return err
	}
	item.ID = id
	item.CreatedAt = createdAt
	item.UpdatedAt = updatedAt
	return nil
}

func (db Database) GetItemById(itemId int) (models.Item, error) {
	item := models.Item{}
	query := `SELECT * FROM items WHERE id = $1;`
	row := db.Conn.QueryRow(query, itemId)
	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM items WHERE id = $1;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	now := time.Now()
	query := `UPDATE items SET name=$1, description=$2, updated_at=$3 WHERE id=$4 RETURNING id, name, description, created_at, updated_at;`
	err := db.Conn.QueryRow(query, itemData.Name, itemData.Description, now, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}
	return item, nil
}
