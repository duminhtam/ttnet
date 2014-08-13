package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Google() revel.Result {
	greeting := "gplus"
	return c.Render(greeting)
}
func (c App) Facebook() revel.Result {
	greeting := "Hello Tam"
	return c.Render(greeting)
}
func (c App) Index() revel.Result {
	greeting := "Hello Tam"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("myName is required")
	c.Validation.MinSize(myName, 4).Message("4 len!!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	var data = map[string]map[string]string{
		"account":{
			"name":myName,
			"age":"27",
		},
		"ads":{
			"title":"Ads Title",
			"body":"Sell s.th",
		},
	}

	return c.Render(data)
}

