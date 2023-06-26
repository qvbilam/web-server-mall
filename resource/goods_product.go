package resource

import proto "mall/api/qvbilam/mall/v1"

type GoodsProductResource struct {
	Id      int64            `json:"id"`
	Count   int64            `json:"count"`
	Product *ProductResource `json:"product"`
}

func (r *GoodsProductResource) Resource(p *proto.GoodsProductResponse) *GoodsProductResource {
	if p == nil {
		return nil
	}

	productRes := ProductResource{}

	var res GoodsProductResource
	res.Id = p.Id
	res.Count = p.Count
	res.Product = productRes.Resource(p.Product)
	return &res
}
