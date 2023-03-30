package main

import (
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

	r.HandleFunc("/user/", showUserIndex)
	r.HandleFunc("/user/{id: [0-9]+}", showUserByKey).Methods("GET")
	r.HandleFunc("/user/create/", showCreateUser)
	// r.HandleFunc("/user/insert/", inputUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// 引数として渡したv (any or interface) をjson形式のデータとして
// string形式で取得する関数
func jsonEncode(v interface{}) string {
	//vをjson形式に変換して返す関数jsonパッケージ内のMarshal関数をbytesに代入している
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err.Error())
	}
	//返すbytesはjson形式のデータで、これをstring形式で返している
	return string(bytes)
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

func showUserByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte(string(jsonEncode(openingdao.FetchByKey(id)))))
}

func showCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("showCreateUserの動作確認")
	update := []byte("Update")
	_, err := w.Write(update)
	if err != nil {
		log.Fatal(err)
	}
}

// func inputUser(w http.ResponseWriter, r *http.Request) {
// 	user := openingdao.Insert(id, firstname, lastname, age, email)

// 	fmt.Println("\nuser内容の確認\n", user)

// 	fmt.Scan()
// 	w.Write([]byte(jsonEncode(openingdao.Insert(id, firstname, lastname, age, email))))
// }
