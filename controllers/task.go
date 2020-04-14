package controllers

import (
	"crm-go/models"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/astaxie/beego"
)

//  TaskController operations for Task
type TaskController struct {
	beego.Controller
}

// URLMapping ...
func (c *TaskController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Task
// @Param	body		body 	models.Task	true		"body for Task content"
// @Success 201 {int} models.Task
// @Failure 403 body is empty
// @router / [post]
func (c *TaskController) Post() {
	var v models.Task
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	user := models.User{Id: c.GetSession("uid").(int64)}
	v.User = &user
	if _, err := models.AddTask(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v.Id

		//检测用户的四个文件是否上传
		_, err := user.GetFiles()
		if err != nil {
			c.Data["json"] = err.Error()
			c.Abort("500")
		}

		err = doPython(user.Id, v.Id)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Abort("500")
		}
		c.Data["json"] = v.Id
	} else {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.ServeJSON()
}

// 执行python脚本
func doPython(id int64, taskId int64) (err error) {
	cmdStr := fmt.Sprintf("cd %srawdata/ && python3 startup.py CASData %d_CAS.csv %d_acct_coverage_by_event.txt %d_CASData.txt %d_visit_history.txt %d",
		beego.AppConfig.String("xiaodai"), id, id, id, id, taskId)
	fmt.Println(cmdStr)
	cmd := exec.Command("bash", "-c", cmdStr)
	err = cmd.Start()
	if err != nil {
		return err
	}
	// err = cmd.Wait()
	return err
}

// GetOne ...
// @Title Get One
// @Description get Task by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Task
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TaskController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTaskById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Task
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Task
// @Failure 403
// @router / [get]
func (c *TaskController) GetAll() {
	user := models.User{Id: c.GetSession("uid").(int64)}
	tasks, err := user.GetTasks()
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.Data["json"] = tasks
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Task
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Task	true		"body for Task content"
// @Success 200 {object} models.Task
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TaskController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Task{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTaskById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Task
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TaskController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteTask(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
