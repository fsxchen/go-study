package dto

import "strings"

type GeneralListDto struct {
	Skip  int    `form:"skip,default=0" json:"skip"`
	Limit int    `form:"limit,default=20" json:"limit" binding:"max=10000"`
	Order string `form:"order" json:"order"`
	Q     string `form:"q" json:"q"`
}

func TransformSearch(qs string, mapping map[string]string) (ss map[string]string) {
	ss = make(map[string]string)
	for _, v := range strings.Split(qs, ",") {
		vs := strings.Split(v, "=")
		if _, ok := mapping[vs[0]]; ok {
			ss[mapping[vs[0]]] = vs[1]
		}
	}
	return
}

type GeneralGetDto struct {
	Id string `uri:"id" json:"id" binding:"required"`
}
