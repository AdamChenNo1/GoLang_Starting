package controllers

import "GoChat/models"

type LongPollingController struct {
	baseController
}

func (c *LongPollingController) Join() {
	uname := c.GetString("uname")
	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	Join(uname, nil)

	c.TplName = "longpolling.html"
	c.Data["IsLongPolling"] = true
	c.Data["Username"] = uname
}

func (c *LongPollingController) Post() {
	c.TplName = "longpolling.html"

	uname := c.GetString("uname")
	content := c.GetString("content")

	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(models.EVENT_MESSAGE, uname, content)
}

func (c *LongPollingController) Fetch() {
	lastReceived, err := c.GetInt("lastReceived")
	if err != nil {
		return
	}

	events := models.GetEvents(int(lastReceived))
	if len(events) > 0 {
		c.Data["json"] = events
		c.ServeJSON()
		return
	}

	//等待新消息
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	c.Data["json"] = models.GetEvents(int(lastReceived))
	c.ServeJSON()
}
