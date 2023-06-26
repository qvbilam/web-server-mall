package resource

import (
	proto "mall/api/qvbilam/mall/v1"
)

type GoodsListResource struct {
	Total int64            `json:"total"`
	List  []*GoodsResource `json:"list"`
}

type GoodsResource struct {
	ID          int64                   `json:"id"`
	CategoryID  int64                   `json:"category_id"`
	Name        string                  `json:"name"`
	Introduce   string                  `json:"introduce"`
	Icon        string                  `json:"icon"`
	PayType     string                  `json:"pay_type"`
	Price       int64                   `json:"price"`
	OriginPrice int64                   `json:"origin_price"`
	Stocks      int64                   `json:"stocks"`
	IsUnlimited bool                    `json:"is_unlimited"`
	IsHot       bool                    `json:"is_hot"`
	Products    []*GoodsProductResource `json:"products"`
}

func (r *GoodsListResource) Resource(p *proto.GoodsListResponse) *GoodsListResource {
	res := GoodsListResource{}
	res.Total = p.Total

	var goods []*GoodsResource

	if res.Total > 0 {
		for _, item := range p.List {
			var ps []*GoodsProductResource
			for _, p := range item.Products {
				var pr ProductResource
				ps = append(ps, &GoodsProductResource{
					Id:      p.Id,
					Count:   p.Count,
					Product: pr.Resource(p.Product),
				})
			}

			goods = append(goods, &GoodsResource{
				ID:          item.Id,
				CategoryID:  item.CategoryId,
				Name:        item.Name,
				Introduce:   item.Introduce,
				Icon:        item.Icon,
				PayType:     item.PayType,
				Price:       item.Price,
				OriginPrice: item.OriginalPrice,
				Stocks:      item.Stocks,
				IsUnlimited: item.IsUnlimited,
				IsHot:       item.IsHot,
				Products:    ps,
			})
		}
		res.List = goods
	}

	return &res
}
