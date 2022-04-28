package restaurantmodel

type Filter struct {
	Name string `json:"name,omitempty" form:"name"`
	Addr string `json:"addr,omitempty" form:"addr"`
}
