$(function () {
    // 加载 BDRP 类型
    // $("#bdrp").empty();
    // todo 加上 bdrp 选项
    loadBNS("localhost", $("#bdrp"));

    // 当 bdrp 发生选择变化的时候，请求最新的 instance
    $("#bdrp").change(function(){
        $("#instances").empty();
        loadInstancesByService($("#bdrp").val(), $("#instances"))
    });

    // 汇总请求
    $('#btn_do').click(function () {
        var bdrp = $("#bdrp").val();
        var address = $("#instances").val();
        var keyname = $("#key").val();
        var option = $("#option").val();

        var uri = "/redis/" + option + "?key=" + keyname +"&address=" + address;

        // 清空 preview 区域旧内容
        // ("#preview").clear();
        $.ajax({
            //请求方式
            type : "GET",
            //请求的媒体类型
            contentType: "application/json;charset=UTF-8",
            //请求地址
            url : uri,
            //请求成功
            success : function(result) {
                var tag = $("<pre><br>" + formatJson(result) + "<br></pre>");
                $("#preview").append(tag);
            },
            //请求失败，包含具体的错误信息
            error : function(e){
                console.log(e.status);
                console.log(e.responseText);
            }
        });
    });

    // 清空预览区内容
    $("#reset").click(function () {
        $("#preview").empty();
    });
});

var formatJson = function (json) {
    var outStr = '',     //转换后的json字符串
        padIdx = 0,         //换行后是否增减PADDING的标识
        space = '    ';   //4个空格符
    if (typeof json !== 'string') {
        json = JSON.stringify(json);
    }
    json = json.replace(/([\{\}\[\]])/g, '\r\n$1\r\n')
        .replace(/(\,)/g, '$1\r\n')
        .replace(/(\r\n\r\n)/g, '\r\n');
    (json.split('\r\n')).forEach(function (node, index) {
        var indent = 0,
            padding = '';
        if (node.match(/[\{\[]/)){
            indent = 1;
        }else if (node.match(/[\}\]]/)){
            padIdx = padIdx !== 0 ? --padIdx : padIdx;
        }else{
            indent = 0;
        }
        for (var i = 0; i < padIdx; i++){
            padding += space;
        }
        outStr += padding + node + '\r\n';
        padIdx += indent;
    });
    return outStr;
};

var loadBNS = function(bns, container) {
    var bnsList = [];
    var uri = "/redis/bnslist?bns=" + bns;
    $.ajax({
        //请求方式
        type : "GET",
        //请求的媒体类型
        contentType: "application/json;charset=UTF-8",
        //请求地址
        url : uri,
        // async: false,
        //请求成功
        success : function(result) {
            for (var index=0; index<result.data.length; index++) {
                var tag = $('<option value="'+result.data[index]+'">'+result.data[index]+'</option>');
                container.append(tag);
            }
        },
        //请求失败，包含具体的错误信息
        error : function(e){
            console.log(e.status);
            console.log(e.responseText);
        }
    });
    // return bnsList;
};

var loadInstancesByService = function (bns, container) {
    var uri = "/redis/getinstances?bns=" + bns;
    console.log(uri);
    $.ajax({
        //请求方式
        type : "GET",
        //请求的媒体类型
        contentType: "application/json;charset=UTF-8",
        //请求地址
        url : uri,
        // async: false,
        //请求成功
        success : function(result) {
            console.log(result);
            for (var index=0; index<result.data.length; index++) {
                console.log(result.data[index]);
                var value = result.data[index].Host + ":" + result.data[index].Port;
                var tag = $('<option value="'+value+'">'+value+'</option>');
                container.append(tag);
            }
        },
        //请求失败，包含具体的错误信息
        error : function(e){
            console.log(e.status);
            console.log(e.responseText);
        }
    });
};