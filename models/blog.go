package models

type Post struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Access      int    `json:"access"`
	TimeCreated string `json:"time_created"`
	UserId      string `json:"user_id"`
	//Likes int `json:"likes"`
}

//
//var db *sql.DB
//var post Post
//
//
//
//func AddPost(c *gin.Context) {
//	err := c.BindJSON(&post)
//	if err != nil{
//		log.Println(err)
//		return
//	}
//	if post.Title == "" || post.Title == " " {
//		c.JSON(400, gin.H{"message":"enter title fields"})
//		c.Redirect(302, "...")
//	}
//	if post.Body == "" || post.Body == " " {
//		//http.Error(w, "Enter Content field", 301)
//		c.JSON(400, gin.H{"message":"enter body field"})
//		c.Redirect(302, "...")
//	} else {
//
//		add := Post{
//			Id:      uuid.New().String(),
//			Title:   post.Title,
//			Body: 	 post.Body,
//			TimeCreated:    time.Now().Format(time.RFC850),
//		}
//
//
//		stmt, er := db.Prepare("insert into news (id, title, content, time)values (?,?,?,?)")
//		defer stmt.Close()
//		if er != nil{
//			log.Println(err)
//			return
//		}
//
//		_, er = stmt.Exec(add.Id, add.Title, add.Body, add.TimeCreated)
//		if er != nil {
//			log.Println("unable to insert data ", err)
//			c.JSON(500, gin.H{"message":"unable to insert data"})
//		}
//		c.Redirect(302,"")
//	}
//
//}
//
//func Home(c *gin.Context) {
//	//result, err := db.Query("select * from news")
//	//
//	//if err != nil {
//	//	log.Println(err)
//	//	c.JSON(500, gin.H{"message":"fetch posts"})
//	//}
//	////var row []DB.Post
//	//for result.Next() {
//	//	//var r DB.Post
//	//	err = result.Scan(post.Id, &post.Title, &post.Body, &post.Likes,&post.TimeCreated)
//	//	if err != nil {
//	//		panic(err.Error())
//	//	}
//	//	//row = append(row, r)
//	//}
//}
//
//func Editpost(w http.ResponseWriter, r *http.Request) {
//	//param := chi.URLParam(r, "Id")
//	//parsedTemplate, err := template.ParseFiles("../template/editpost.html")
//	//error.ParseTempError(err)
//	//
//	//e := DB.Post{}
//
//	//db.QueryRow("select * from news where id=?", param).Scan(&e.Id, &e.Title, &e.Content, &e.Time)
//
//}
//
//func EditProcess(w http.ResponseWriter, r *http.Request) {
//	//param := chi.URLParam(r, "Id")
//	//r.ParseForm()
//	//title := r.FormValue("title")
//	//content := r.FormValue("content")
//	//
//	//if title == "" || title == " " {
//	//	http.Error(w, "Enter Title field", 301)
//	//	http.Redirect(w, r, "/edit/{Id}", http.StatusFound)
//	//}
//	//if content == "" || content == " " {
//	//	http.Error(w, "Enter Content field", 301)
//	//	http.Redirect(w, r, "/edit/{Id}", http.StatusFound)
//	//} else {
//	//	stmt, err := db.Prepare("update news set title=?, content=? where id=?")
//	//	error.ErrorChecker(err)
//	//	stmt.Exec(title, content, param)
//	//	log.Println("Table has been Upadated")
//	//}
//	//http.Redirect(w, r, "/home", http.StatusFound)
//}
//
//func DeletePost(w http.ResponseWriter, r *http.Request) {
//	//param := chi.URLParam(r, "Id")
//	//delPost, err := db.Prepare("delete from news where id=?")
//	//defer delPost.Close()
//	//error.ErrorChecker(err)
//	//delPost.Exec(param)
//	//http.Redirect(w, r, "/home", http.StatusFound)
//}
