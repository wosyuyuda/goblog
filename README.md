# gomybolg

# 初始化
最新版博客网址:https://www.iq2.cn  
 1.先在配置文件config/set.json 输入初始的数据库,网站名称,这个后期都是可以修改的  
 2.go mod tidy初始化一下包  
 3.运行main.go//路由结构请看main.go的注释说明  
 4.进入初始化 /install/ 创建用户,第一篇文章,第一个分类  
 5.默认账号密码,longfei 123456,这个在insall里面,后期有空也放到set.json里面去   
 6.测试地址:http://test.ibiaoqin.cn/  管理地址 http://test.ibiaoqin.cn/admin  
 7.上面的测试地址加了一个nginx负载均衡,两个服务器的数据库未做同步,所以有时候刷新数据会不一样.  
 8.如果是初学者,建议查看dao,config目录,一个封装了数据库,一个封装了设置目录,比较杂的一些功能都是放在controller目录,,server跟util目录功能有些混装,等空了再重构一下,现在先放着  

# 喜欢的帮忙点个赞哪
业余默默的写也不容易,对你有帮助的话不要吝啬自己的start,如果有问题的话可以给我留言哈,还有不少待优化的地方,待后面空了补齐

# 介绍

1. 主要还是mvc结构，部分会使用前后端分离，看哪个便利哪个来  
2. golang的gin框架+gorm.v2+hui后端管理，前端界面考虑用mip的模板，更利于优化  
3. 博客的基本功能，后台登陆功能已经好了，再加上增删改查的功能，再加一个分类功能就好啦  

# 软件架构
1. golang的gin框架+gorm.v2+hui后端管理，前端界面考虑用mip的模板
2. 项目的依赖请自行查看go.mod包,主要依赖gorm,gin,验证码

# 目录结构
│  go.mod 包的声明文件,初始化需要用到  
│  go.sum  
│  main.go启动文件,主包,包含路由文件  
│      
├─config 设置  
│      db.json 数据库的配置文件  
│      set.json  
│      
├─controller 控制器文件  
│      admin.go 管理员的控制器  
│      config.go 设置相关  
│      index.go 前台的控制器  
│      types.go 分类相关  
│      user.go 用户的一些功能相关  
│      
├─install 初即化的一些文件,  
│      install.go初始化用户,新增加一个分类,新加文章与用户  
│      
├─middleware 中间件  
│      islogin.go判断是否登陆的是中间件  
│      
├─model 连接数据库,数据库的结构声明文件  
│      model.go  
│      
├─server 服务文件  
│      cap.go 验证码  
│      config.go 设置文件  
│      db.go 数据库的服务  
│      
├─static  静态文件目录,包含h ui的一些静态文件与ueditor的静态文件,再加上后台的一些文件  
│  │  
│  ├─css  
│  │              
│  ├─lib  
│  │                  
│  ├─static  
│  │  ├─h-ui  
│  │                  
│  ├─ueditor   
│  │               
│  └─view 后台新增编辑分类与文章,新加了一个设置页面  
│          admin_addtype.html  分类管理  
│          admin_addview.html  文章增改   
│          admin_set.html  设置  
│          
├─util 常用的一些工具页面  
│      body.go  内容的一些工具  
│      md5.go  md5  
│      page.go  分类的一些设置  
│      session.go  
│      ueditor.go  编辑器  
│      
└─view 这里是前端的一些模板文件,admin是后台的,其它的为前台的  
        admin_addview.html  
        admin_index.html  
        admin_list.html  
        index.html  
        list.html   
        login.html  
        login1.html  
        view.html  



# 安装教程
1. 更改数据库/config/set.json里面的，俩个地方都改一下  
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
```

	newList := FindListNew("0", 1)      //最新  
	tuijian := FindListNew("-4", 1)     //推荐  

    func FindListNew(id string, page int) (vi []d.View) {  
    .......省略部分  
	switch id {  
	case "0":  
		db.Limit(num).Offset(page * num).Order(order).Preload("Tps").Find(&vi)  
    .......省略部分  
	default:  
		db.Where("typeid = ?", id).Limit(num).Offset(page * num).Order(order).Preload  ("Tps").Find(&vi)  
	}  
	return  
}  
```

    如果是php函数可以 这样写(参考thinkphp6),使用的时候是可以不传值,也可以只传一个或者俩个,真是贼方便   

```
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
```