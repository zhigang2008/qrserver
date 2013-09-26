package httpserver

import (
	"github.com/astaxie/beego"
	"html/template"
	"io/ioutil"
	"net/http"
)

var errtpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<title>{{.Title}}</title>
		<style type="text/css">
			* {
				margin:0;
				padding:0;
			}

			body {
				background-color:#EFEFEF;
				font: .9em "Lucida Sans Unicode", "Lucida Grande", sans-serif;
			}

			#wrapper{
				width:600px;
				margin:40px auto 0;
				text-align:center;
				-moz-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
				-webkit-box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
				box-shadow: 5px 5px 10px rgba(0,0,0,0.3);
			}

			#wrapper h1{
				color:#FFF;
				text-align:center;
				margin-bottom:20px;
			}

			#wrapper a{
				display:block;
				font-size:.9em;
				padding-top:20px;
				color:#FFF;
				text-decoration:none;
				text-align:center;
			}

			#container {
				width:600px;
				padding-bottom:15px;
				background-color:#FFFFFF;
			}

			.navtop{
				height:40px;
				background-color:#24B2EB;
				padding:13px;
			}

			.content {
				padding:10px 10px 25px;
				background: #FFFFFF;
				margin:;
				color:#333;
			}

			a.button{
				color:white;
				padding:15px 20px;
				text-shadow:1px 1px 0 #00A5FF;
				font-weight:bold;
				text-align:center;
				border:1px solid #24B2EB;
				margin:0px 200px;
				clear:both;
				background-color: #24B2EB;
				border-radius:100px;
				-moz-border-radius:100px;
				-webkit-border-radius:100px;
			}

			a.button:hover{
				text-decoration:none;
				background-color: #24B2EB;
			}

		</style>
	</head>
	<body>
		<div id="wrapper">
			<div id="container">
				<div class="navtop">
					<h1>{{.Title}}</h1>
				</div>
				<div id="content">
					{{.Content}}
					<a href="/" title="Home" class="button">返回首页</a><br />
				</div>
			</div>
		</div>
	</body>
</html>
`

//404页面
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t := template.New("steven404template")

	body, err := ioutil.ReadFile(beego.ViewsPath + "/404.html")
	if err != nil {
		t, _ = t.Parse(errtpl)
	} else {
		t, _ = t.Parse(string(body))
	}

	data := make(map[string]interface{})
	data["Title"] = "Page Not Found"
	data["Content"] = "该页面不存在"
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(http.StatusNotFound)
	t.Execute(rw, data)
}

//401页面未授权
func page_unauth(rw http.ResponseWriter, r *http.Request) {
	t := template.New("steven401template")

	body, err := ioutil.ReadFile(beego.ViewsPath + "/401.html")
	if err != nil {
		t, _ = t.Parse(errtpl)
	} else {
		t, _ = t.Parse(string(body))
	}
	data := make(map[string]interface{})
	data["Title"] = "Unauthorized"
	data["Content"] = "您没有访问该页面的权限"
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(http.StatusUnauthorized)
	t.Execute(rw, data)
}
