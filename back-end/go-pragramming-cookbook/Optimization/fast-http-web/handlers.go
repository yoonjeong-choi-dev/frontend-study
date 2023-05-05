package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (c *ItemController) GetItemsHandler(ctx *fasthttp.RequestCtx) {
	encoder := json.NewEncoder(ctx)
	items := c.readItems()
	encoder.Encode(&items)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (c *ItemController) AddItemsHandler(ctx *fasthttp.RequestCtx) {
	item, ok := ctx.UserValue("item").(string)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	c.addItem(item)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
