package dbclient

import (
	"servicesdemo/accountservice/model"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

const accountBucket = "AccountBucket"

// IBoltClient interface for boltDb(key-value store)
type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountID string) (model.Account, error)
	Seed()
}

// BoltClient is implemetation of IBoltClient
type BoltClient struct {
	boltDb *bolt.DB
}

// OpenBoltDb open a connection with accounts.db
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDb, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Seed adding fake accounts
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

// QueryAccount query a account by id
func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltDb.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(accountBucket))

		accountBytes := b.Get([]byte(accountID))

		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}

		return json.Unmarshal(accountBytes, &account)
	})

	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(accountBucket))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) seedAccounts() {
	total := 100

	for i := 0; i < total; i++ {
		//generate a key
		key := strconv.Itoa(10000 + i)

		// create a instance of account struct
		acc := model.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}
		fmt.Println(acc)

		// serialize
		jsonBytes, _ := json.Marshal(acc)

		// write the data in accountbucket
		err := bc.boltDb.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(accountBucket))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})

		if err != nil {
			fmt.Printf("error in inserting: %v", err)
		}
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}
