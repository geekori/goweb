  <form action="/service/testlogin1" method="post">
{{ .User }}：<input type="text" name="username"><p/>
{{ .PW }}：<input type="password" name="password"><p/>
       <input type="submit" value="登录"><p/>
   </form>