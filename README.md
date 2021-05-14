# gomybolg

# 介绍

1. 主要还是mvc结构，部分会使用前后端分离，看哪个便利哪个来  
2. golang的gin框架+gorm.v2+hui后端管理，前端界面考虑用mip的模板，更利于优化  
3. 博客的基本功能，后台登陆功能已经好了，再加上增删改查的功能，再加一个分类功能就好啦  

# 软件架构
1. golang的gin框架+gorm.v2+hui后端管理，前端界面考虑用mip的模板
2. 项目的依赖请自行查看go.mod包,主要依赖gorm,gin,验证码

# 目录结构
1. controller 控制器
    1.1 --admin.go 管理后台文件  
    1.2 --index.go 前端页面的文件  
    1.3 --user.go 啥用没有，后期留着用来管理用户或者管理员  
    1.4 --list.go 列表功能，本来应该放在admin.go里面的，不过划分出来结构会更清晰一些  
2. main.go启动文件，包含路由  
3. model 数据库目录，添加数据库处理类的功能,新的已经把数据的结构全放到这个文件里了  
4. server目录，后期添加一下处理服务的功能，现在没啥东西就先空着  
5. static 静态资源目录，放一些css,js图片等资源  
    5.1 css-lib-temp目录为hui前端框架的一些依赖     
    5.2 view目录为添加编辑的一些静态html  
    5.3 upload为上传的图片文件
6. view目录用来放模板的目录  
7. util目录用来放一些工具,比如处理列表的图片,转换数据格式,处理分页的一些数据,获取设置删除session,百度编辑器接收图片上传的一些工具  
8. install 目录是初始化的目录,后面用来初始一些数据  
9. middleware 中间件的目录,用来放一些中间件,比如验证登陆权限的  



# 安装教程
1. 更改数据库/model/model.go里面的，俩个地方都改一下  
2. 这个拉到本地，运行main.go即可  

# 使用说明

static里面的静态资源是h ui的模板，想研究的自己去研究一下吧  

# 参与贡献
1. h ui  http://h-ui.net/  
2. gin+gorm  
3. 部分使用layui，及jq ajax  
4. 后端考虑部分使用vue(/static/view),功能部分  
5. 前端模板考虑使用百度mip，更利于百度优化  



# 注意事项
本来打算添加编辑分类跟文章功能页面直接用mvc结构的,实在是vue.js的语法跟gin的模板语法冲突,木有办法,所以有俩个静态文件放在/static/view目录下面,分别是添加编辑文章与分类的功能  


# 使用总结
### 如果选效率的话,我用Thinkphp,如果看成长性的话就golang了

1. 这gin golang真不太适合用mvc结构,还是前后端分开比较好,特别是涉及到功能页面,不太利于前端功能的扩展,特别是像我这种功能页面比较喜欢用vue去写的人,vue写功能是真的方便  
2. 还有这个数据库用着也麻烦,得单独去找包,弄配置文件还得写(现在都还没有写用配置)  
如果功能页面不用vue的话,也特么难用,模板功能不支持三元运算,不然我不用专门写一个处理数据的util.body.Imgsrc里面循环把每个数据里面判断有没有图片,总不能我自己写个支持三元运算的模板语法吧  
3. golang的优点的话,熟悉了之后使用起来很方便,所有的函数有严格的数据格式要求,感觉比我之前用PHP学得快,只要输出这个函数,就知道他要传哪些数据进去,这些数据是什么格式,返回来的数据是什么,一看就懂,这里得点个赞  
4. 再说一个golang的缺点,不支持默认参数!!!这影响了我很多的写方法,比如我想写一个获取分类的,如果是用PHP写,我会写成不传值获取全部,传值则获取这个值的分类,如果第二个值不传则默认不获取二级参数,如果传了就获取,而golang不行..必须得传,所以有些获取功能看着很怪,参考下面  
golang的获取  
`	newList := FindListNew("0", 1)      //最新
	tuijian := FindListNew("-4", 1)     //推荐

    func FindListNew(id string, page int) (vi []d.View) {
    .......省略部分
	switch id {
	case "0":
		db.Limit(num).Offset(page * num).Order(order).Preload("Tps").Find(&vi)
    .......省略部分
	default:
		db.Where("typeid = ?", id).Limit(num).Offset(page * num).Order(order).Preload("Tps").Find(&vi)
	}
	return
}
    `

    如果是php函数可以 这样写(参考thinkphp6),使用的时候是可以不传值,也可以只传一个或者俩个,真是贼方便   

`
$newlist = $this->getlist();//获取全部
public function getlist($id=0,$page=1,$limt=10)
{
    $model=new ViewModel;
    //笑死,php if里面声明的变量外面也可以用,golang不行.
    if($id==0){
        $data=$model->limit($limt)->page($page)->select();
    }else{
        $data=$model->where("id",$id)->limit($limt)->page($page)->select();
    }
    return $data;
}
`