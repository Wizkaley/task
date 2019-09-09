package main

import(

	bolt "go.etcd.io/bbolt"
	"time"	
)


func main(){
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second}) ; if err!=nil{
		panic(err)
	}
	defer db.Close()
}