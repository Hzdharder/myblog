package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"myblog/dao"
	"myblog/model"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.PostForm("rusername")
	password := c.PostForm("rpassword")

	user := model.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.Register(&user)
	c.Redirect(301, "/")
}
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("lusername")
	password := c.PostForm("lpassword")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在")
		fmt.Println("用户名不存在")
	} else {
		if u.Password != password {
			c.HTML(200, "login.html", "密码错误")
			fmt.Println("密码错误")
		} else {
			fmt.Println("登入成功")
			c.Redirect(301, "/")
		}
	}
}

// 博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)

}

// 添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(302, "/post_index")
}

// 跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}
func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)
	content := blackfriday.Run([]byte(p.Content))
	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}

func ShowFace(c *gin.Context) {
	c.HTML(200, "face.html", nil)
}

func ShowStar(c *gin.Context) {
	c.HTML(200, "starsky.html", nil)
}

func DeletePost(c *gin.Context) {
	s := c.Query("dpid")
	pid, _ := strconv.Atoi(s)
	dao.Mgr.DeletePost(pid)
	c.Redirect(301, "/post_index")
}
func GoModifyPost(c *gin.Context) {
	s := c.Query("mpid")
	pid, _ := strconv.Atoi(s)
	post := dao.Mgr.GetPost(pid)
	c.HTML(200, "modify.html", post)
}

func ModifyPost(c *gin.Context) {
	s := c.Query("kpid")
	title := c.PostForm("mtitle")
	tag := c.PostForm("mtag")
	content := c.PostForm("mcontent")
	pid, _ := strconv.Atoi(s)
	dao.Mgr.ModifyPost(pid, title, content, tag)
	c.Redirect(301, "/post_index")
}
