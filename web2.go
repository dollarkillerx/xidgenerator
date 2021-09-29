package internal

import "github.com/gin-gonic/gin"

func (s *Server) home(ctx *gin.Context) {
	ctx.Header("content-type", "text/html charset=utf-8")
	ctx.Writer.Write([]byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>ERROR RECOVERY</title>\n    <!-- CSS only -->\n    <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"\n          integrity=\"sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU\" crossorigin=\"anonymous\">\n</head>\n<body>\n<div class=\"container\">\n    <div class=\"jumbotron\">\n        <h1>ERROR RECOVERY</h1>\n        <p>Smoothie 错误恢复</p>\n        <p>\n            <a class=\"btn btn-primary btn-lg\" href=\"/v1\" role=\"button\">根据ID文件恢复</a>\n\n            <a class=\"btn btn-primary btn-lg\" href=\"/v2\" role=\"button\">根据服务版本恢复</a>\n        </p>\n    </div>\n</div>\n\n\n<!-- JavaScript Bundle with Popper -->\n<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js\"\n        integrity=\"sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ\"\n        crossorigin=\"anonymous\"></script>\n</body>\n</html>\n"))
}

func (s *Server) v1(ctx *gin.Context) {
	ctx.Header("content-type", "text/html charset=utf-8")
	ctx.Writer.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>ERROR RECOVERY</title>

    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css"
          integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap-theme.min.css"
          integrity="sha384-6pzBo3FDv/PJ8r2KRkGHifhEocL+1X2rVCTTkUfGk7/0pbek5mMa1upzvWbrUbOZ" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <!--    <script src="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"-->
    <!--            integrity="sha384-aJ21OjlMXNL5UyIl/XNwTMqvzeRMZH2w8c5cRVpzpU8Y5bApTppSuUkhZXN0VxHd"-->
    <!--            crossorigin="anonymous"></script>-->

    <script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.min.js"
            integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>

    <style>
        /*动态加载圈-loading-start*/
        body .loading-box-shadow-omg {
            width: -webkit-fill-available;
            height: -webkit-fill-available;
            background-color: #211f1f5c;
            position: absolute;
            top: 0;
            z-index: 9999999;
        }

        .hidden {
            display: none !important;
        }

        body .loading-box-shadow-omg .loading-box {
            background-color: white;
            border-radius: 5px;
            position: absolute;
            top: 20%;
            left: 40%;
            width: 20%;
            height: 25%;
        }

        body .loading-box-shadow-omg .loading-box .loading-circle {
            width: 80px;
            height: 80px;
            background-color: transparent;
            position: absolute;
            left: 35%;
            top: 10%;
            animation: init-circle 1s linear infinite;
        }

        body .loading-box-shadow-omg .loading-box .loading-content {
            position: absolute;
            bottom: 5%;
            font-weight: bold;
            color: rebeccapurple;
            width: 100%;
            text-align: center;
        }

        body .loading-box-shadow-omg .loading-box .loading-circle > div {
            background-color: #292961;
            border-radius: 20px;
            position: absolute;
        }

        @keyframes init-circle {
            from {
                transform: rotate(360deg);
            }
            to {
                transform: rotate(0deg);
            }
        }
    </style>
</head>
<body>
<div class="container">
    <div class="jumbotron">
        <h1>Smoothie ERROR RECOVERY</h1>
        <p> 根据ID文件恢复</p>
        <p>

            <!--        <form method="post" action="/recovery_v1" enctype="multipart/form-data">-->
        <form id="vx">
            <div class="form-group">
                <label for="serverName">当前服务名称</label>
                <input type="text" class="form-control" id="serverName" placeholder="Server Name " name="serverName">
            </div>
            <div class="form-group">
                <label for="version">当前服务版本</label>
                <input type="text" class="form-control" id="version" placeholder="version" name="version">
            </div>
            <div class="form-group">
                <label for="upperLayerVersion">上级服务版本</label>
                <input type="text" class="form-control" id="upperLayerVersion" placeholder="Upper Layer Version"
                       name="upperLayerVersion">
            </div>
            <div class="form-group">
                <label for="upperLayerCollection">上级服务Mongo Collection</label>
                <input type="text" class="form-control" id="upperLayerCollection"
                       placeholder="Upper Layer Collection" name="upperLayerCollection">
            </div>
            <div class="form-group">
                <label for="upperLayerTopic">上级服务 Kafka Topic</label>
                <input type="text" class="form-control" id="upperLayerTopic" placeholder="kafka topic"
                       name="upperLayerTopic">
            </div>
            <div class="form-group">
                <label for="startTime">查询起始时间</label>
                <input type="text" class="form-control" id="startTime" placeholder="Start Time" value="2021-06-27"
                       name="startTime">
            </div>
            <div class="form-group">
                <label for="fromJson">ID JSON FILE</label>
                <input type="file" id="fromJson" name="fromJson">
            </div>

            <br>
            <button type="button" class="btn btn-default" id="sub">Submit</button>
        </form>
        </p>
    </div>
</div>

<script>
    $("#sub").on("click", function () {
        let formData = new FormData(document.querySelector("#vx"))
        initLoading({
            type: "start",
            content: "loading...",
        });
        $.ajax({
            url: "/recovery_v1",
            type: "POST",
            data: formData,
            processData: false,  // 不处理数据
            contentType: false,   // 不设置内容类型
            success: function (data) {
                initLoading({
                    type: "stop",
                });
                var str = JSON.stringify(data);
                console.log(data)
                alert("成功: " + str)
            },
            error: function (data) {
                initLoading({
                    type: "stop",
                });
                var str = JSON.stringify(data);
                console.log("提交失败！！！:", str)
                alert("提交失败！！！:" + str)
            }
        });
    });

    function initLoading(property) {
        $("body .loading-box-shadow-omg").remove();//清除加载图层
        if (property.type == "stop") {
            return;
        }
        var nodeHtml = '<div class="loading-box-shadow-omg"><div class="loading-box">';
        nodeHtml += '<div class="loading-circle"></div><div class="loading-content"></div></div></div>';
        $("body").append(nodeHtml);//加载布局元素
        var html = '<div style="top: 5%;left: 53%;width: 5px;height: 5px;"></div>';
        html += '<div style="top: 11%;left: 30%;width: 7px; height: 7px;"></div>';
        html += '<div style="top: 26%;left: 12%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 48%;left: 7%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 70%;left: 15%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 85%;left: 33%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 89%;left: 54%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 80%;left: 75%;width: 9px;height: 9px;"></div>';
        $("body .loading-circle").html(html);
        var content = "正在加载中...";//提示内容
        if (property.content) {
            content = property.content;
        }
        $("body .loading-content").text(content);
        var shadowColor = "#211f1f5c";//阴影颜色
        if (property.shadowColor) {
            shadowColor = property.shadowColor;
            $('body .loading-box-shadow-omg').css("background-color", shadowColor);
        }
        var loadingBoxColor = "white";//加载框背景色
        if (property.loadingBoxColor) {
            loadingBoxColor = property.loadingBoxColor;
            $('body .loading-box-shadow-omg .loading-box').css("background-color", loadingBoxColor);
        }
        var loadingPointColor = "#292961";//动态点颜色
        if (property.loadingPointColor) {
            loadingPointColor = property.loadingPointColor;
            $('body .loading-box-shadow-omg .loading-circle>div').css("background-color", loadingPointColor);
        }
        var loadingContentColor = "rebeccapurple";//提示内容颜色
        if (property.loadingContentColor) {
            loadingContentColor = property.loadingContentColor;
            $('body .loading-box-shadow-omg .loading-content').css("color", loadingContentColor);
        }
    };
    // initLoading({
    //     type: "start",
    //     content: "loading...",
    //     // shadowColor:"grey",
    //     // loadingBoxColor:"yellow",
    //     // loadingPointColor:"green",
    //     // loadingContentColor:"red"
    // });
</script>

</body>
</html>
`))
}

func (s *Server) v2(ctx *gin.Context) {
	ctx.Header("content-type", "text/html charset=utf-8")
	ctx.Writer.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>ERROR RECOVERY</title>

    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css"
          integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap-theme.min.css"
          integrity="sha384-6pzBo3FDv/PJ8r2KRkGHifhEocL+1X2rVCTTkUfGk7/0pbek5mMa1upzvWbrUbOZ" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <!--    <script src="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"-->
    <!--            integrity="sha384-aJ21OjlMXNL5UyIl/XNwTMqvzeRMZH2w8c5cRVpzpU8Y5bApTppSuUkhZXN0VxHd"-->
    <!--            crossorigin="anonymous"></script>-->

    <script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.min.js"
            integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>

    <style>
        /*动态加载圈-loading-start*/
        body .loading-box-shadow-omg {
            width: -webkit-fill-available;
            height: -webkit-fill-available;
            background-color: #211f1f5c;
            position: absolute;
            top: 0;
            z-index: 9999999;
        }

        .hidden {
            display: none !important;
        }

        body .loading-box-shadow-omg .loading-box {
            background-color: white;
            border-radius: 5px;
            position: absolute;
            top: 20%;
            left: 40%;
            width: 20%;
            height: 25%;
        }

        body .loading-box-shadow-omg .loading-box .loading-circle {
            width: 80px;
            height: 80px;
            background-color: transparent;
            position: absolute;
            left: 35%;
            top: 10%;
            animation: init-circle 1s linear infinite;
        }

        body .loading-box-shadow-omg .loading-box .loading-content {
            position: absolute;
            bottom: 5%;
            font-weight: bold;
            color: rebeccapurple;
            width: 100%;
            text-align: center;
        }

        body .loading-box-shadow-omg .loading-box .loading-circle > div {
            background-color: #292961;
            border-radius: 20px;
            position: absolute;
        }

        @keyframes init-circle {
            from {
                transform: rotate(360deg);
            }
            to {
                transform: rotate(0deg);
            }
        }
    </style>
</head>
<body>
<div class="container">
    <div class="jumbotron">
        <h1>Smoothie ERROR RECOVERY</h1>
        <p> 根据服务版本恢复</p>
        <p>

            <!--        <form method="post" action="/recovery_v1" enctype="multipart/form-data">-->
        <form id="vx">
            <div class="form-group">
                <label for="serverName">当前服务名称</label>
                <input type="text" class="form-control" id="serverName" placeholder="Server Name " name="serverName">
            </div>
            <div class="form-group">
                <label for="version">当前服务版本</label>
                <input type="text" class="form-control" id="version" placeholder="version" name="version">
            </div>
            <div class="form-group">
                <label for="upperLayerVersion">上级服务版本</label>
                <input type="text" class="form-control" id="upperLayerVersion" placeholder="Upper Layer Version"
                       name="upperLayerVersion">
            </div>
            <div class="form-group">
                <label for="upperLayerCollection">上级服务Mongo Collection</label>
                <input type="text" class="form-control" id="upperLayerCollection"
                       placeholder="Upper Layer Collection" name="upperLayerCollection">
            </div>
            <div class="form-group">
                <label for="upperLayerTopic">上级服务 Kafka Topic</label>
                <input type="text" class="form-control" id="upperLayerTopic" placeholder="kafka topic"
                       name="upperLayerTopic">
            </div>
            <div class="form-group">
                <label for="startTime">查询起始时间</label>
                <input type="text" class="form-control" id="startTime" placeholder="Start Time" value="2021-06-27"
                       name="startTime">
            </div>

            <br>
            <button type="button" class="btn btn-default" id="sub">Submit</button>
        </form>
        </p>
    </div>
</div>

<script>
    $("#sub").on("click", function () {
        let formData = new FormData(document.querySelector("#vx"))
        initLoading({
            type: "start",
            content: "loading...",
        });
        $.ajax({
            url: "/recovery_v2",
            type: "POST",
            data: formData,
            processData: false,  // 不处理数据
            contentType: false,   // 不设置内容类型
            success: function (data) {
                initLoading({
                    type: "stop",
                });
                var str = JSON.stringify(data);
                console.log(data)
                alert("成功: " + str)
            },
            error: function (data) {
                initLoading({
                    type: "stop",
                });
                var str = JSON.stringify(data);
                console.log("提交失败！！！:", str)
                alert("提交失败！！！:" + str)
            }
        });
    });

    function initLoading(property) {
        $("body .loading-box-shadow-omg").remove();//清除加载图层
        if (property.type == "stop") {
            return;
        }
        var nodeHtml = '<div class="loading-box-shadow-omg"><div class="loading-box">';
        nodeHtml += '<div class="loading-circle"></div><div class="loading-content"></div></div></div>';
        $("body").append(nodeHtml);//加载布局元素
        var html = '<div style="top: 5%;left: 53%;width: 5px;height: 5px;"></div>';
        html += '<div style="top: 11%;left: 30%;width: 7px; height: 7px;"></div>';
        html += '<div style="top: 26%;left: 12%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 48%;left: 7%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 70%;left: 15%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 85%;left: 33%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 89%;left: 54%;width: 9px;height: 9px;"></div>';
        html += '<div style="top: 80%;left: 75%;width: 9px;height: 9px;"></div>';
        $("body .loading-circle").html(html);
        var content = "正在加载中...";//提示内容
        if (property.content) {
            content = property.content;
        }
        $("body .loading-content").text(content);
        var shadowColor = "#211f1f5c";//阴影颜色
        if (property.shadowColor) {
            shadowColor = property.shadowColor;
            $('body .loading-box-shadow-omg').css("background-color", shadowColor);
        }
        var loadingBoxColor = "white";//加载框背景色
        if (property.loadingBoxColor) {
            loadingBoxColor = property.loadingBoxColor;
            $('body .loading-box-shadow-omg .loading-box').css("background-color", loadingBoxColor);
        }
        var loadingPointColor = "#292961";//动态点颜色
        if (property.loadingPointColor) {
            loadingPointColor = property.loadingPointColor;
            $('body .loading-box-shadow-omg .loading-circle>div').css("background-color", loadingPointColor);
        }
        var loadingContentColor = "rebeccapurple";//提示内容颜色
        if (property.loadingContentColor) {
            loadingContentColor = property.loadingContentColor;
            $('body .loading-box-shadow-omg .loading-content').css("color", loadingContentColor);
        }
    };
    // initLoading({
    //     type: "start",
    //     content: "loading...",
    //     // shadowColor:"grey",
    //     // loadingBoxColor:"yellow",
    //     // loadingPointColor:"green",
    //     // loadingContentColor:"red"
    // });
</script>

</body>
</html>
`))
}

