package db

func Main() (n int) {

	n = 10
	return // return n や return 10と書いても問題ありません。

}

// "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
// "github.com/katsuomi/gin-like-twitter-api/models"
// func Db() (
// 	a := "aa"
// )
// var (
// 	db  *gorm.DB
// 	err error
// )

// // Init is initialize db from main function
// func Init() {
// 	db, err = gorm.Open("mysql", "host=db port=8080 user=gin dbname=gin password=gin1234 sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	autoMigration()
// 	user := models.User{
// 		ID:    1,
// 		Name:  "aoki",
// 		Posts: []models.Post{{ID: 1, Content: "tweet1"}, {ID: 2, Content: "tweet2"}},
// 	}
// 	db.Create(&user)
// }

// // GetDB is called in models
// func GetDB() *gorm.DB {
// 	return db
// }

// // Close is closing db
// func Close() {
// 	if err := db.Close(); err != nil {
// 		panic(err)
// 	}
// }

// func autoMigration() {
// 	db.AutoMigrate(&models.User{})
// 	db.AutoMigrate(&models.Post{})
// }
