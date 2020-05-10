package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (this *baseController) Prepare() {
	// Reset language option.
	this.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	// 2. Default language is Chinese.
	if len(this.Lang) == 0 {
		this.Lang = "zh-CN"
	}

	// Set template level language option.
	this.Data["Lang"] = this.Lang
}

type MainController struct {
	baseController
}

func (c *MainController) Get() {
	c.TplName = "welcome.html"
}

func (c *MainController) Join() {
	uname := c.GetString("uname")
	tech := c.GetString("tech")

	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		c.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		c.Redirect("/ws?uname="+uname, 302)
	default:
		c.Redirect("/"+uname, 302)
	}

	return
}
