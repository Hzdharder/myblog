package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"myblog/model"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User

	//博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
	DeletePost(pid int)
	ModifyPost(pid int, title string, content string, tag string)
}
type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("链接失败")
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}
func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(&user)
}
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(&post)
}
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
func (mgr *manager) DeletePost(pid int) {
	var post model.Post
	mgr.db.Where("ID=?", pid).Delete(&post)
}
func (mgr *manager) ModifyPost(pid int, title string, content string, tag string) {

	mgr.db.Table("posts").Where("ID IN ?", []int{pid}).Updates(map[string]interface{}{"Title": title, "Content": content, "Tag": tag})
}
