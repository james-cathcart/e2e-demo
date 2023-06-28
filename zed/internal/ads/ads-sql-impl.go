package ads

import (
	"database/sql"
	"e2ezed/internal/types"
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
