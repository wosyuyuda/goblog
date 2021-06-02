

var gtime = Getdata();
function Getdata(sj_str) {
    var data = new Date().toLocaleDateString()
    var dd = Date.parse(data) / 1000
    var iday = Math.floor(parseInt(dd - sj_str) / 60 / 60 / 24);
    if (-1 == iday) {
        return '今天'
    } else if (0 == iday) {
        return '昨天'
    } else {
        var dd = [];
        dd = new Date(sj_str * 1000).toLocaleDateString().split('/');
        var y = dd[0];
        var m = dd[1];//获取当前月份的日期
        var d = dd[2];
        return y + "-" + m + "-" + d;
    }
}