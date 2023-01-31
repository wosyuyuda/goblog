
var body = $('body');
var Oparent = $('<div>');
var row = $('<span>');

// style
Oparent.css({
	'position': 'fixed',
	'left': '50%',
	'transform': 'translate(-50%)',
	'-ms-transform': 'translate(-50%)',
	'-webkit-transform': 'translate(-50%)',
	'-moz-transform': 'translate(-50%)',
	'-o-transform': 'translate(-50%)',
	'z-index': '100',
});
row.css({
	'font-size': '14px',
	'color': '#ffffff',
	'padding': '5px 10px',
	'background': 'rgba(0, 0, 0, .5)',
	'-webkit-border-radius': '20px',
	'-moz-border-radius': '20px',
	'-o-border-radius': '20px',
	'border-radius': '20px', 
});


function showToast(text,pos, duration) {
	row.text('');	//清空内容

	// 判断是否传入值
	if(!text) {	
		return;
	}
	row.append(text);
	Oparent.append(row).fadeIn();	//插入内容，并实现渐显效果
	body.append(Oparent);
	duration = duration ? duration : 2000;

	// 提示框显示位置
	if(pos === 'top') {
		Oparent.css('top', '10%');
	}else if(pos === 'center') {
		Oparent.css('top', '50%');
	}else if(pos === 'bottom') {
		Oparent.css('bottom', '10%');
	}else {	//默认居中显示
		Oparent.css('top', '50%');
	}


	var t = setTimeout(function() {	//隐藏toast
		Oparent.fadeOut();	//添加渐隐效果
	}, duration)
}