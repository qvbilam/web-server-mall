package resource

import proto "mall/api/qvbilam/mall/v1"

type ProductResource struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Type  string `json:"type"`
	Tag   string `json:"tag"`
}

func (r *ProductResource) Resource(p *proto.ProductResponse) *ProductResource {
	if p == nil {
		return nil
	}
	var res ProductResource
	res.Id = p.Id
	res.Name = p.Name
	res.Price = p.Price
	res.Type = p.Type
	res.Tag = p.Tag

	return &res
}
