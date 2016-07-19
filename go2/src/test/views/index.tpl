<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }


    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h1 class="logo"></h1>
      <div>
        if：{{if .IsGood}}very Good
        {{else}}Bad!
        {{end}}<br/>
        对象变量：{{.p.Name}}-{{.p.Age}}<br/>
        with用法：{{with .p}}
        {{.Name}}-{{.Age}}
        {{end}}<br/>
        遍历数组：{{range .ns}}
          {{.}}
        {{end}}<br/>
        设置变量：{{$v11 := .v1}}
        {{$v11}}<br/>
        html转义:{{str2html .htmlVar}}<br/>
         html转义:{{.htmlVar| htmlquote}}<br/>
         {{template "tt"}}
      </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>

    </div>
  </footer>
  <div class="backdrop"></div>

</body>
</html>
{{define "tt"}}
  <div>hello</div>
{{end}}
