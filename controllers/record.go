package controllers

import (
	"crm-go/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

//  RecordController operations for Record
type RecordController struct {
	beego.Controller
}

// URLMapping ...
func (c *RecordController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Record
// @Param	body		body 	models.Record	true		"body for Record content"
// @Success 201 {int} models.Record
// @Failure 403 body is empty
// @router / [post]
func (c *RecordController) Post() {
	var v models.Record
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddRecord(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Record by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Record
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RecordController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetRecordById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Record
// @Success 200 {object} models.Record
// @Failure 403
// @router / [get]
func (c *RecordController) GetAll() {
	records, err := models.GetRecordByUserId(c.GetSession("uid").(int64))
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.Data["json"] = records
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Record
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Record	true		"body for Record content"
// @Success 200 {object} models.Record
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RecordController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Record{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateRecordById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Record
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RecordController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteRecord(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
