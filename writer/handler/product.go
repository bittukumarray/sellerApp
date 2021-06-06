package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"writer/db"
	"writer/model"
)

func Create(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res model.ResponseModel

	err = json.Unmarshal(body, &res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong data format"))
		return
	}

	client, err := db.GetMongoClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("database connection error"))
		return
	}

	res.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	collection := client.Database(db.DB).Collection(db.ISSUES)

	_, err = collection.InsertOne(context.TODO(), res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while inserting data"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("data inserted successfully"))
}
