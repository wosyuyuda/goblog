package server

import (
	"goblog/config"
	"goblog/dao"
	"goblog/model"
	"goblog/util"
)

/* var JoinDao *gorm.DB

func init() {
	db := dao.MDB.Table("views").Select("views.id,views.title,views.click,views.created_at,views.pic,views.typeid,views.content, tps.name as Typename")
	JoinDao = db.Joins("left join tps on tps.id = views.typeid")
} */

//获取列表,恶心死了gorm- 字段无法进入join关联查询结果...调了我半天.
func GetViewlist(id interface{}, page int) (vi []model.ViewJson) {
	db := dao.MDB.Table("views").Select("views.id,views.title,views.click,views.created_at,views.pic,views.typeid,views.content, tps.name as Typename")
	JoinDao := db.Joins("left join tps on tps.id = views.typeid").Where("views.status = 1")

	num := 10 //一页默认10条
	if page < 1 {
		page = 1
	}
	page -= 1 //过滤掉的页面要减 1
	order := "created_at desc"
	switch id {
	case "0":
		JoinDao.Limit(num).Offset(page * num).Order(order).Find(&vi)
	case "-1":
		JoinDao.Where("tuijian = ?", 1).Limit(num).Order(order).Find(&vi)
	case "-2":
		JoinDao.Where("swiper = ?", 1).Limit(num / 2).Order(order).Find(&vi)
	case "-3":
		JoinDao.Limit(num).Offset(page * num).Order("click desc").Find(&vi)
	case "-4":
		JoinDao.Where("tuijian = ?", 1).Limit(3).Order(order).Find(&vi)
	default:

		JoinDao.Where("typeid = ?", id).Limit(num).Offset(page * num).Order(order).Find(&vi)
		//fmt.Printf("123%+v", &vi)
	}
	return
}

//获取当前分类下面的10条文章
func Findlist2(id string) (vi []model.ViewJson) {
	db := dao.MDB.Table("views").Select("views.id,views.title,views.click,views.created_at,views.pic,views.typeid,views.content, tps.name as Typename")
	JoinDao := db.Joins("left join tps on tps.id = views.typeid").Where("views.status = 1")
	if id != "0" {
		JoinDao = JoinDao.Where("typeid = ?", id)
	}
	JoinDao.Limit(10).Order("created_at desc").Find(&vi)
	return
}

//获取搜索到的信息
func SearchView(views *model.ListInfo) (err error) {
	db := dao.MDB.Table("views").Select("views.id,views.title,views.click,views.created_at,views.pic,views.typeid,views.content, tps.name as Typename")
	JoinDao := db.Joins("left join tps on tps.id = views.typeid").Where("views.status = 1")
	//util.Ini(&views.Page)
	//fmt.Printf("pageinfo=%+v", *views)
	//开始处理一下page信息
	err = JoinDao.Where("title LIKE ?", "%"+views.Page.Keyword+"%").Limit(views.Page.Num).Order("created_at desc").Find(&views.Views).Count(&views.Page.Sum).Error
	return
}

//获取列表新的
func GetListV(pa *model.PageList) (views *model.ListInfo, err error) {
	//pa := views.Page

	if pa.Page == 0 {
		pa.Page = 1
	}
	if pa.Num == 0 {
		pa.Num = 10
	}
	views, err = util.GetListCache(pa)
	if err == nil {
		return
	}
	if views == nil {
		views = new(model.ListInfo)
	}
	//fmt.Printf("info%+v", views)
	//util.Ini(&pa)
	db := dao.MDB.Table("views").Where("status = 1 AND typeid = ?", pa.ID)
	if pa.Keyword != "" {
		db = db.Where("title LIKE ?", "%"+pa.Keyword+"%")
	}
	db = db.Limit(pa.Num).Offset((pa.Page - 1) * pa.Num).Order("id desc")
	err = db.Find(&views.Views).Count(&pa.Sum).Error
	views.Page = pa
	//fmt.Println("数据库找到的数据列表是:", views.Views)
	util.PagesinfoTo(views)  //处理一下页面分类信息
	util.SetListCache(views) //设置进缓存
	return
}

//获取最新,推荐,tdk,cid为分类id，用于后面判断当前分类是否选中
func Getinfo(id int) (baseinfo *model.BaseInfo, err error) {
	base, err := util.GetCache(util.BaseCache)
	if err != nil {
		baseinfo = new(model.BaseInfo)
		//fmt.Println("没有缓存")
		baseinfo.New = util.Imgsrc(GetViewlist("0", 1))      //最新
		baseinfo.Tuijian = util.Imgsrc(GetViewlist("-4", 1)) //推荐
		baseinfo.Tdk = config.GetTDK()
		err = dao.MDB.Where("status = ?", "1").Find(&baseinfo.Typeinfo).Error //获取全部分类信息
		util.SetCache(util.BaseCache, baseinfo)
	} else {
		baseinfo = base.(*model.BaseInfo)
	}
	for k, v := range baseinfo.Typeinfo {
		if id == v.ID {
			baseinfo.Clicktp = v
		}
		baseinfo.Typeinfo[k].IsTz = v.SkipPath != ""
		baseinfo.Typeinfo[k].IsTrue = id == v.ID
	}
	return
}
func F文档获取基础信息并更新(arc *model.View) (baseinfo *model.BaseInfo, err error) {
	baseinfo, err = Getinfo(arc.Typeid)
	if err != nil {
		return
	}
	arc.Tempdir = "view.html"
	for k := range baseinfo.Typeinfo {
		baseinfo.Typeinfo[k].IsTrue = baseinfo.Typeinfo[k].ID == arc.Typeid
		if baseinfo.Typeinfo[k].ID == arc.Typeid {
			arc.Typename = baseinfo.Typeinfo[k].Name
			if baseinfo.Typeinfo[k].ArcTempdir != "" {
				arc.Tempdir = baseinfo.Typeinfo[k].ArcTempdir
			}
		}
		if baseinfo.Typeinfo[k].SkipPath != "" {
			baseinfo.Typeinfo[k].IsTz = true
		}
	}
	return
}
