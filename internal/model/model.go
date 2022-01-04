package model

type Model struct {
	// id  int(10) unsigned is_nullable NO
	Id uint32 `gorm:"primary_key" json:"id"`
	// 创建时间  int(10) unsigned is_nullable YES
	CreatedOn uint32 `json:"created_on"`
	// 创建人  varchar(100) is_nullable YES
	CreatedBy string `json:"created_by"`
	// 修改时间  int(10) unsigned is_nullable YES
	ModifiedOn uint32 `json:"modified_on"`
	// 修改人  varchar(100) is_nullable YES
	ModifiedBy string `json:"modified_by"`
	// 删除时间  int(10) unsigned is_nullable YES
	DeletedOn uint32 `json:"deleted_on"`
	// 是否删除 0 为未删除、1 为已删除  tinyint(3) unsigned is_nullable YES
	IsDel uint8 `json:"is_del"`
}
