// JavaScript Document
/*
	Tabs : 选项卡函数
	兼容：ie6,ie7,ie8,ie9,ff
	
	参数说明：
	pubCardClass : 选项卡公共样式
	cardHoverClass : 鼠标选中选项卡后，选项卡固定的样式
	cardOutClass : 未选中的选项卡样式
	pubContClass ： 选项卡内容公共样式
	
	说明：
	选项卡ID与选项卡内容ID需自行设定，对应关系
	选项卡：#cardID, 选项卡内容：cardID
	
	应用实例：
	载入jq库文件：<script src="jquery-1.5.2.min.js" type="text/javascript"></script>
	jq库下载地址：http://docs.jquery.com/Downloading_jQuery
	
	<div id="MainBody">
		<div id="Panel">
			<div style="height:100%;overflow:hidden">
				<ul>
					<li id="#card1" class="tt current">1</li>
					<li id="#card2" class="tt out">2</li>
					<li id="#card3" class="tt out">3</li>
				</ul>
			</div>
			
			<div id="card1" class="tt1">
				card1
			</div>
			
			<div id="card2" class="tt1" style="display:none">
				card2
			</div>
			
			<div id="card3" class="tt1" style="display:none">
				card3
			</div>
		</div>
	</div>
	<script>Tabs(".tt",".current",".out",".tt1");</script>
*/
function Tabs(pubCardClass,cardHoverClass,cardOutClass,pubContClass){
	$(pubCardClass).on("mouseover",function(){
		var targetClass = $(this).attr("class");
		var targetId = $(this).attr("id")
		var classValue = pubCardClass.slice(1) + " " +  cardHoverClass.slice(1);
		
		if(targetClass != classValue)
		{
			$(pubCardClass).removeClass(cardHoverClass.slice(1));
			$(pubCardClass).removeClass(cardOutClass.slice(1));
			$(pubCardClass).addClass(cardOutClass.slice(1));
			$(this).removeClass(cardOutClass.slice(1));
			$(this).addClass(cardHoverClass.slice(1));
			$(pubContClass).hide();
			$(targetId).show();
			var num = $(this).attr('num');
			$('.fanganbg').removeClass('f_active')
			$('.fanganbg'+num).addClass('f_active');
		}
	})		
}



//设为收藏
/*
使用方法：
	<a href="javascript:;" onclick="AddFavorite('网页地址','网站标题')">加入收藏夹|加入书签藏</a>
	网页地址：document.location.href
	网页标题：document.title
*/
function AddFavorite(sURL, sTitle) { 
try { 
//IE 
window.external.addFavorite(sURL, sTitle); 
} catch (e) { 
try { 
//Firefox 
window.sidebar.addPanel(sTitle, sURL, ""); 
} catch (e) { 
alert("您的浏览器不支持自动加入收藏，请手动设置！", "提示信息"); 
} 
} 
}