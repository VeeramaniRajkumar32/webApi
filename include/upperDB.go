package include

// import (
// 	"fmt"
// 	"log"

// 	"github.com/go-sql-driver/mysql"
// 	"github.com/upper/db/v4/adapter/mysql"
// )

// var dbConn = mysql.ConnectionURL{
// 	Database: `blog`,
// 	Host:     `localhost`,
// 	username: `root`,
// 	Password: ``,
// }

// func see() {
// 	sess, err := mysql.Open(dbConn)
// 	if err != nil {
// 		log.Fatal("Open: ", err)
// 	}
// 	defer sess.Close()

// 	fmt.Printf("Connected to %q with DSN:\n\t%q", sess.Name(), dbConn)
// }
