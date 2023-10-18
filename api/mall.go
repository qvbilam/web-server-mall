package api

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "mall/api/qvbilam/mall/v1"
	"mall/global"
	"mall/resource"
	"mall/validate"
	"strconv"
)

// Category 商品分类
func Category(ctx *gin.Context) {
	p, _ := ctx.GetQuery("pid")
	pid, _ := strconv.Atoi(p)

	c, err := global.MallServerClient.GetCategory(context.Background(), &proto.CategoryListRequest{
		Pid: int64(pid),
	})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	SuccessNotMessage(ctx, c)
}

// GoodsList 商品列表
func GoodsList(ctx *gin.Context) {
	request := validate.GoodsList{}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	params := &proto.GoodsListRequest{
		OnSale:      true,
		NeedProduct: true,
	}
	if request.CategoryId != nil {
		rCid := *request.CategoryId
		cid, _ := strconv.Atoi(rCid)
		params.CategoryId = int64(cid)
	}
	if request.Page != nil {
		rPage := *request.Page
		page, _ := strconv.Atoi(rPage)
		params.Page.Page = int64(page)
	}
	if request.PerPage != nil {
		rPerPage := *request.PerPage
		perPage, _ := strconv.Atoi(rPerPage)
		params.Page.PerPage = int64(perPage)
	}

	g, err := global.MallServerClient.GetGoodsList(context.Background(), params)
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	res := resource.GoodsListResource{}
	SuccessNotMessage(ctx, res.Resource(g))
	//SuccessNotMessage(ctx, g)
}

// GoodsDetail 商品详情
func GoodsDetail(ctx *gin.Context) {
	SuccessNotContent(ctx)
}

// GoodsSell 商品售卖
func GoodsSell(ctx *gin.Context) {
	paramId, _ := ctx.GetQuery("id")
	id, _ := strconv.Atoi(paramId)

	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.GoodsSell{}
	if err := ctx.ShouldBind(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}

	// 购买
	res, err := global.MallServerClient.Sell(context.Background(), &proto.SellRequest{
		Id:     int64(id),
		Count:  request.Count,
		UserId: userID,
	})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	SuccessNotMessage(ctx, res)
}
