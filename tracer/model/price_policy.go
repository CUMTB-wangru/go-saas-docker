package model

import "gorm.io/gorm"

type PricePolicy struct {
	gorm.Model
	Catogory      string `gorm:"type:varchar(64)" json:"catogory" label:"收费类型"`
	Title         string `gorm:"type:varchar(64)" json:"title" label:"标题"`
	Price         uint64 `gorm:"type:int(5)" json:"price" label:"价格"`
	ProjectNum    uint64 `json:"project_num" label:"项目数量"`
	ProjectMember uint64 `json:"project_member" label:"项目成员人数"`
	ProjectSpace  uint64 `json:"project_space" label:"每个项目空间" help_text:"单位是M"`
	PerFileSize   uint64 `json:"per_file_size" label:"单文件大小" help_text:"单位是M"`
}

// GetAllBlog 查询所有博客信息
func GetAllBlog() PricePolicy {
	var allBlog PricePolicy
	DB.Find(&allBlog)
	return allBlog
}

// TypeBlog 根据类型查找博客
func TypeBlog(tyb string) PricePolicy {
	var typeBlog PricePolicy
	DB.Where("type=?", tyb).Find(&typeBlog)
	return typeBlog
}

// TopBlog 置顶博客查询
func TopBlog(top string) PricePolicy {
	var topBlog PricePolicy
	DB.Where("top=?", top).Find(&topBlog)
	return topBlog
}
