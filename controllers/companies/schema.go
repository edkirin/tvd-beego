package companies

import "tvd-beego/models"

type MetaJson struct {
	ItemsCount int64 `json:"itemsCount"`
}

type CompaniesResponseJson struct {
	Items []*models.Company `json:"items"`
	Meta  MetaJson          `json:"meta"`
}

type MachinesResponseJson struct {
	Items []*models.Machine `json:"items"`
	Meta  MetaJson          `json:"meta"`
}
