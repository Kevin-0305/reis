package es

type EsCluterGroup struct {
	EsCluterId     uint `gorm:"column:es_cluter_id"`
	ProjectGroupId uint `gorm:"column:project_group_id"`
}

func (s *EsCluterGroup) TableName() string {
	return "es_cluter_group"
}
