package controllers

import (
	"encoding/json"
	"strconv"
	"todo/models"

	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) ListTasks() {
	res := struct {
		Tasks []*models.Task
	}{models.DefaultTaskList.All()}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *TaskController) NewTask() {
	req := struct {
		Title string
	}{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("空标题"))
		return
	}

	t, err := models.NewTask(req.Title)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	models.DefaultTaskList.Save(t)
}

func (c *TaskController) GetTask() {
	id := c.Ctx.Input.Param(":id")
	beego.Info("任务号：")
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("已找到", ok)
	if !ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("任务未找到"))
		return
	}
	c.Data["json"] = t
	c.ServeJSON()
}

func (c *TaskController) UpdateTask() {
	id := c.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &t); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}
