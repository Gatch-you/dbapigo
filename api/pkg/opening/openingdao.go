package openingdao

import (
	"api/pkg/db"
	"fmt"
	"log"
)

// User型の構造体
// json形式にするために、jsonのタグを追加している
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func FetchIndex() []User {
	//dbと接続し、一連の処理が終わったらdbを閉じる
	db := db.Connect()
	defer db.Close()

	// fmt.Printf("FetchIndexの内容")
	// fmt.Println(db)

	//db.Query(" ")にて" "内のクエリを実行して、その結果をrowsに代入している。
	//よって、rowsにはuser内の全てのデータが代入されている。
	//rowsは&{0x1400013e6c0 0x104dca5a0 0x1400006edc0 <nil> <nil> {{0 0} 0 0 {{} 0} {{} 0}} false <nil> []}
	//↑なんだこれ？
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	//rowsの確認コード
	// fmt.Println("\n\nrowsの内容\n", rows)

	//userArgsというUser構造体を持つ配列スライスを初期化、作成する。
	userArgs := make([]User, 0)
	// fmt.Println("\nこの時点でのuseArgsの内容\n", userArgs)
	//rows.Next()にて次の行(データ)がなくなるまで処理を実行する
	for rows.Next() {
		var user User
		//なぜerrに代入している？このあとerrを出力しても<nill>である
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		//append関数を使ってuserArgsに取得したuser情報を入れている
		userArgs = append(userArgs, user)

	}

	// fmt.Println("\n最終的なrows\n", rows)
	// fmt.Println("\n最終的なuseArgs\n", userArgs)
	return userArgs
}

func FetchByKey(id string) []User {
	db := db.Connect()
	defer db.Close()

	////db.Query(" ")にて" "内のクエリを実行して、その結果をrowsに代入している。
	rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		log.Fatal(err.Error())
	}
	userArgs := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.Email)
		if err != nil {
			log.Fatal(err.Error())
		}
		userArgs = append(userArgs, user)
	}
	return userArgs
}

// 失敗！ローカルサーバーが応答しなかった...
func Create(id int, firstname string, lastname string, age int, email string) []User {
	db := db.Connect()
	defer db.Close()

	var user User

	stmt, err := db.Prepare("INSERT INTO user (id, firstname, lastname, age, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.ID, "userID")
	//この下でエラーが出てるっぽい
	// var number int
	// number, _ = strconv.Atoi(id)

	// var ageNumber int
	// ageNumber, _ = strconv.Atoi(age)

	_, err = stmt.Exec(id, firstname, lastname, age, email)
	if err != nil {
		log.Fatal(err)
	}

	f := FetchIndex()
	return f
}

// func Update(id int, firstname string, lastname string, age int, email string) []User {
// 	db := db.Connect()
// 	defer db.Close()

// 	updt, err := db.Prepare("UPDATE user SET firstname = ?, lastname = ?, age = ?, email = ? WHERE id = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = updt.Exec(id, firstname, lastname, age, email)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	f := FetchByKey()
// 	return f
// }

// openingdaoと処理を分割したコード
func Delete(id string) []User {
	db := db.Connect()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = delete.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User ID = %v has been deleted sucsessfuly", id)
	f := FetchByKey(id)
	return f
}
