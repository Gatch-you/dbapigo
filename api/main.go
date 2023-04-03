package main

import (
	connect "api/pkg/db"
	openingdao "api/pkg/opening"
	"fmt"

	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/", showUserIndex)
	r.HandleFunc("/api/user/create", createUser).Methods("POST")
	r.HandleFunc("/api/user/{id:[0-9]+}", deleteUser).Methods("DELETE")
	r.HandleFunc("/api/user/update/{id:[0-9]+}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func showUserIndex(w http.ResponseWriter, r *http.Request) {
	//openingdaoファイル上のFetchIndexを呼び出し、userとして宣言している。
	user := openingdao.FetchIndex()

	fmt.Println("\nuserの内容\n", user)
	//userの内容をjson形式に変換し、bytesに初期化、変数として設定する。
	bytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	//先ほどのbytesをstring形式に変換し、その後
	w.Write(([]byte(string(bytes))))
	fmt.Println("\n最終的な出力\n", ([]byte(string(bytes))))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	defer db.Close()

	var user openingdao.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO user (id, firstname, lastname, age, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.ID, user.FirstName, user.LastName, user.Age, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(([]byte(string(bytes))))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Insert a New User Profile!"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := connect.Connect()
	defer db.Close()

	var user openingdao.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	updt, err := db.Prepare("UPDATE user SET firstname = ?, lastname = ?, age = ?, email = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = updt.Exec(user.FirstName, user.LastName, user.Age, user.Email, id)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(([]byte(string(bytes))))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Update a User Profile!"))

	// update, err := db.Prepare("")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := connect.Connect()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = delete.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID = %s has been deleted sucsessfuly\n", id)

}
