package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ここ
type Client struct {
	*http.Client
}

// ここ
func NewClient(c *http.Client) Client {
	return Client{c}
}

type User struct {
	ID       string
	Location string
}

func (c Client) FetchUser(id string) (user *User) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/users/"+id, nil)

	resp, _ := c.Do(req) // ここ
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &user)
	return user
}
func main() {
	// ここ
	qiitaClient := NewClient(http.DefaultClient)
	user := qiitaClient.FetchUser("yyh-gl")
	fmt.Println("============ RESULT ============")
	fmt.Printf("%+v\n", user)
	fmt.Println("============ RESULT ============")
}

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func main() {
// 	dsn := "gin:gin1234@tcp(172.23.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	r := gin.Default()

// 	// r.GET("/get-product", func(c *gin.Context) {
// 	// 	var product Product
// 	// 	db.First(&product, 1)                 // find product with integer primary key
// 	// 	db.First(&product, "code = ?", "D42") // find product with code D42
// 	// 	fmt.Printf("%v\n", *result)
// 	// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	// 	"status": "success
// 	// 	// })
// 	// 	// result := db.Select("price").Find(&product)
// 	// 	// c.JSONP(http.StatusOK, gin.H{
// 	// 	// 	"message": "ok",
// 	// 	// 	"data":    result,
// 	// 	// })
// 	// 	// c.String(http.StatusOK, "Hello world")
// 	// 	// db.Create(&Product{Code: "D42", Price: 100})
// 	// })

//   // GetAll is get all Post
// type PostRepository struct{}
// func (_ PostRepository) GetAll() ([]Product, error) ;{
//   db := db.GetDB()
//   var u []Product
//   if err := db.Table("products").Select("price").Scan(&u).Error; err != nil {
//       return nil, err
//   }
//   return u, nil
// }

// type PostController struct{}
//   func (pc PostController) Index(c *gin.Context); {
//     // var u repository.PostRepository
//     p, err := GetAll()
//     if err != nil {
//         c.AbortWithStatus(404)
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//     } else {
//         c.JSON(200, p)
//     }
// }

// func() Init() ;{
//   r := router() *gin.Engine {
//     r := gin.Default()

//    u := r.Group("/product")
//    {
//        // ctrl := controllers.UserController{}
//        u.GET("", Index)
//        // u.POST("", ctrl.Create)
//        // u.GET("/:id", ctrl.Show)
//        // u.PUT("/:id", ctrl.Update)
//        // u.DELETE("/:id", ctrl.Delete)
//    }

//    return r
// }
//   r.Run()
// }
// Init()
// 	r.Run(":8080")
// }
