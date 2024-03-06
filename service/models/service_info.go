package models

type ServiceInfo struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Mac              string `json:"mac"`
	Domain           string `json:"domain"`
	Ip               string `json:"ip"`
	IsPing           bool   `json:"isPing"`
	IsManuallyModify bool   `json:"isManuallyModify"`
}

func (s ServiceInfo) CompareContentsEqual(o ServiceInfo) bool {
	if s.Id != o.Id ||
		s.Name != o.Name ||
		s.Mac != o.Mac ||
		s.Domain != o.Domain ||
		s.Ip != o.Ip {
		return false
	}
	return true
}
