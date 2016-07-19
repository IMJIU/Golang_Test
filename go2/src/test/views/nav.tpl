{{define "nav"}}
<nav class="navbar navbar-default navbar-fixed-top">
  <div class="container">
    <a href="/" class="navbar-brand">我的博客</a>
    <ul class="nav navbar-nav">
      <li class="active"><a href="/">首页</a></li>
      <li class=""><a href="/category">分类</a></li>
      <li class=""><a href="">文章</a></li>
       <li class=""><a href="/login">登陆</a></li>
        <li class=""><a href="/logout">退出</a></li>
    </ul>
  </div>
</nav>
<div class="pull-right navbar-right">
    <ul class="nav navbar-nav navbar-right">
       {{if .IsLogin}}
       <li><a href="/login?exit=true">退出</a></li>
       {{else}}
       <li><a href="/login">管理员登陆</a></li>
        {{end}}
    </ul>
</div>
{{end}}
