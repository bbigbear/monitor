package controllers

type GrowController struct {
	BaseController
}

func (this *GrowController) Get() {
	this.TplName = "index.tpl"
}
