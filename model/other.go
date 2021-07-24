package model

//友情链接
type Link struct {
	ID   uint   `gorm:"primarykey" form:"id" json:"id"`
	Url  string `gorm:"size:255;" form:"url" json:"url"`
	Name string `gorm:"size:255;" form:"name" json:"name"`
}

//分类页码,分类信息
type PageList struct {
	ID      int    ` form:"id" json:"id" uri:"id"`                 //分类ID
	Keyword string ` form:"keyword" json:"keyword"  uri:"keyword"` //关键词
	Page    int    ` form:"page" json:"page" uri:"page"`           //页码
	Num     int    ` form:"num" json:"num" uri:"num"`              //一页多少条,默认是10
	Sum     int64  ` form:"sum" json:"sum" uri:"sum"`              //全部多少条
	Nex     int    ` form:"nex" json:"nex"`                        //上一页
	Up      int    ` form:"up" json:"up"`                          //下一页
	SumPage int    ` form:"sumpage" json:"sumpage"`
}

//列表信息
type ListInfo struct {
	Page     PageList   //分类信息
	Views    []ViewJson //文章列表
	Listinfo Tp         //当前分类
}

//最新,推荐,tdk信息
type BaseInfo struct {
	New      []ViewJson //最新
	Tuijian  []ViewJson //推荐
	Tdk      Tdk        //tdk
	Typeinfo []Tp       //全部分类信息
}
