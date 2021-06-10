package luffy

type Payload struct {
	BoxCode         string       `json:"box_code"`
	BoxRouteID      string       `json:"box_route_id"`
	PartnerCode     string       `json:"partner_code"`
	BusinessType    BusinessType `json:"business_type"`
	ServiceCode     string       `json:"service_code"`
	RefType         string       `json:"ref_type"`
	OldRouteID      string       `json:"old_route_id"`
	StandardBoxCode string       `json:"standard_box_code"`
}
