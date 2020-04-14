package controllers

import (
	"crm-go/models"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//  FileController operations for File
type FileController struct {
	beego.Controller
}

// URLMapping ...
func (c *FileController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description 上传文件接口
// @Param	name		formData 	string	true		"上传文件名"
// @Param	file		formData 	file	true		"上传的文件"
// @Success 200 {int} file的ID
// @Failure 400 用户错误
// @Failure 500 服务端错误
// @router / [post]
func (c *FileController) Post() {
	name := c.GetString("name")
	f, _, err := c.GetFile("file")
	//获取文件失败
	if err != nil {
		c.Data["json"] = "获取文件失败" + err.Error()
		c.Abort("400")
	}
	defer f.Close()
	uid := c.GetSession("uid").(int64)
	//文件存储失败
	if err = c.SaveToFile("file", fmt.Sprintf("%srawdata/CASData/%d_%s", beego.AppConfig.String("xiaodai"), uid, name)); err != nil {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	o := orm.NewOrm()
	user := models.User{Id: uid}
	file := models.File{Name: name, User: &user}
	if o.Read(&file, "Name", "User") == nil {
		file.Time = time.Now()
		_, err := o.Update(&file, "Time")
		if err != nil {
			c.Data["json"] = err.Error()
			c.Abort("500")
		}
		c.Data["json"] = file.Id
	} else {
		file.Time = time.Now()
		_, err := o.Insert(&file)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Abort("500")
		}
		c.Data["json"] = file.Id
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get File by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.File
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FileController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetFileById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get File
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.File
// @Failure 403
// @router / [get]
func (c *FileController) GetAll() {
	user := models.User{Id: c.GetSession("uid").(int64)}
	files, err := user.GetFiles()
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.Data["json"] = files
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the File
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.File	true		"body for File content"
// @Success 200 {object} models.File
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FileController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// v := models.File{Id: id}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// if err := models.UpdateFileById(&v); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the File
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FileController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteFile(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.ServeJSON()
}
