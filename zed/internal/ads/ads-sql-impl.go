package ads

import (
	"database/sql"
	"e2ezed/internal/types"
	"errors"
	"log"
)

type SQLImpl struct {
	conn *sql.DB
}

func New(conn *sql.DB) DAO {

	return &SQLImpl{
		conn: conn,
	}
}

func (dao *SQLImpl) GetAdsByCustomer(customer string) ([]types.Ad, error) {

	log.Printf("getting ads for customer: %s", customer)

	query := `SELECT rowid, customer, info FROM ads WHERE customer=$1 ORDER BY rowid DESC`

	response, err := dao.conn.Query(query, customer)
	if err != nil {
		return nil, err
	}

	defer response.Close()

	var ads []types.Ad
	var tmp types.Ad
	for response.Next() {

		err := response.Scan(&tmp.ID, &tmp.Customer, &tmp.Info)
		if err != nil {
			break
		}

		ads = append(ads, tmp)

	}

	return ads, err

}

func (dao *SQLImpl) GetAdByID(id int64) ([]types.Ad, error) {

	log.Printf("getting ads for customer: %d", id)

	query := `SELECT rowid, customer, info FROM ads WHERE rowid=$1 LIMIT 1`

	response := dao.conn.QueryRow(query, id)

	var tmp types.Ad
	err := response.Scan(&tmp.ID, &tmp.Customer, &tmp.Info)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Printf("error: %s\n", err.Error())
		return nil, err
	}

	var adSlice []types.Ad
	adSlice = append(adSlice, tmp)

	return adSlice, nil

}

func (dao *SQLImpl) GetAllAds() ([]types.Ad, error) {

	log.Println(`getting all ads`)

	query := `SELECT rowid, customer, info FROM ads ORDER BY customer ASC`
	response, err := dao.conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer response.Close()

	var ads []types.Ad
	var tmp types.Ad
	for response.Next() {

		err := response.Scan(&tmp.ID, &tmp.Customer, &tmp.Info)
		if err != nil {
			break
		}

		ads = append(ads, tmp)

	}

	return ads, err

}

func (dao *SQLImpl) CreateAd(ad types.Ad) (int64, error) {

	log.Printf("adding new Ad record for customer: %s\n", ad.Customer)

	query := `INSERT INTO ads (customer, info) VALUES ($1, $2)`
	stmt, err := dao.conn.Prepare(query)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return -1, err
	}

	result, err := stmt.Exec(ad.Customer, ad.Info)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return -1, err
	}
	defer stmt.Close()

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return -1, err
	}

	return lastID, nil

}

func (dao *SQLImpl) Update(ad types.Ad) error {

	log.Printf("updating ad with ID: %d", ad.ID)

	query := `UPDATE ads SET customer=$1, info=$2 WHERE rowid=$3`
	stmt, err := dao.conn.Prepare(query)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(ad.Customer, ad.Info, ad.ID)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	if rowsAffected < 1 {
		err = errors.New("no rows were affected")
		log.Printf("error: %s\n", err.Error())
		return err
	}

	return nil
}

func (dao *SQLImpl) DeleteAdByID(id int64) error {

	log.Printf("deleting add with ID: %d", id)

	query := `DELETE FROM ads WHERE rowid=$1`
	stmt, err := dao.conn.Prepare(query)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	if numRowsAffected < 1 {
		err = errors.New("error: no rows affected")
		log.Printf("error: %s\n", err)
		return err
	}

	return nil

}
