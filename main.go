package main

import (
	"log"
	"strings"

	"github.com/dollarkillerx/erguotou"
	"github.com/rs/xid"
)

func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.Ctx.SetContentType("text/html")
		ctx.Write(200, []byte(strings.ReplaceAll(html, "{xid}", xid.New().String())))
	})

	err := app.Run(erguotou.SetHost("0.0.0.0:8975"))
	if err != nil {
		log.Fatalln(err)
	}
}

var html = `
<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU" crossorigin="anonymous">

    <title>Xid Generate</title>

    <style>
        body {
            background-color: black;
            color: aliceblue;
        }
    </style>
</head>
<body>
<div class="jumbotron">
    <div class="container">
        <div class="text-center">
            <h1>Online Xid Generate</h1>
            <p class="info">Your XID:</p>
            <h2 class="uuid">
                <span id="generated-uuid">{xid}</span>
                <br>
                <button id="copy-button" class="btn btn-primary" title="Copy the UUID to the clipboard."
                        alt="Copy the UUID to the clipboard.">
                    <img src="data:image/svg+xml;utf8;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iaXNvLTg4NTktMSI/Pgo8IS0tIEdlbmVyYXRvcjogQWRvYmUgSWxsdXN0cmF0b3IgMTkuMS4wLCBTVkcgRXhwb3J0IFBsdWctSW4gLiBTVkcgVmVyc2lvbjogNi4wMCBCdWlsZCAwKSAgLS0+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgdmVyc2lvbj0iMS4xIiBpZD0iQ2FwYV8xIiB4PSIwcHgiIHk9IjBweCIgdmlld0JveD0iMCAwIDQ4OC4zIDQ4OC4zIiBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IDAgMCA0ODguMyA0ODguMzsiIHhtbDpzcGFjZT0icHJlc2VydmUiIHdpZHRoPSIxNnB4IiBoZWlnaHQ9IjE2cHgiPgo8Zz4KCTxnPgoJCTxwYXRoIGQ9Ik0zMTQuMjUsODUuNGgtMjI3Yy0yMS4zLDAtMzguNiwxNy4zLTM4LjYsMzguNnYzMjUuN2MwLDIxLjMsMTcuMywzOC42LDM4LjYsMzguNmgyMjdjMjEuMywwLDM4LjYtMTcuMywzOC42LTM4LjZWMTI0ICAgIEMzNTIuNzUsMTAyLjcsMzM1LjQ1LDg1LjQsMzE0LjI1LDg1LjR6IE0zMjUuNzUsNDQ5LjZjMCw2LjQtNS4yLDExLjYtMTEuNiwxMS42aC0yMjdjLTYuNCwwLTExLjYtNS4yLTExLjYtMTEuNlYxMjQgICAgYzAtNi40LDUuMi0xMS42LDExLjYtMTEuNmgyMjdjNi40LDAsMTEuNiw1LjIsMTEuNiwxMS42VjQ0OS42eiIgZmlsbD0iI0ZGRkZGRiIvPgoJCTxwYXRoIGQ9Ik00MDEuMDUsMGgtMjI3Yy0yMS4zLDAtMzguNiwxNy4zLTM4LjYsMzguNmMwLDcuNSw2LDEzLjUsMTMuNSwxMy41czEzLjUtNiwxMy41LTEzLjVjMC02LjQsNS4yLTExLjYsMTEuNi0xMS42aDIyNyAgICBjNi40LDAsMTEuNiw1LjIsMTEuNiwxMS42djMyNS43YzAsNi40LTUuMiwxMS42LTExLjYsMTEuNmMtNy41LDAtMTMuNSw2LTEzLjUsMTMuNXM2LDEzLjUsMTMuNSwxMy41YzIxLjMsMCwzOC42LTE3LjMsMzguNi0zOC42ICAgIFYzOC42QzQzOS42NSwxNy4zLDQyMi4zNSwwLDQwMS4wNSwweiIgZmlsbD0iI0ZGRkZGRiIvPgoJPC9nPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+CjxnPgo8L2c+Cjwvc3ZnPgo=">
                    <span id="copy-button-text" onclick="copy()">Copy</span>
                </button>
            </h2>
            <p class="refresh"><a onclick="reload()">Refresh</a> page to generate another.</p>
        </div>
    </div>
</div>

<script>
    function copy() {
        let aux = document.getElementById("generated-uuid").innerText;
        const input = document.createElement('input');
        document.body.appendChild(input);
        input.setAttribute('value', aux);
        input.select();
        if (document.execCommand('copy')) {
            document.execCommand('copy');
            console.log('????????????');
        }
        document.body.removeChild(input);
    }

    function reload() {
        location.reload();
    }
</script>

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
        crossorigin="anonymous"></script>

</body>
</html>

`
