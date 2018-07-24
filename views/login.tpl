<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
	</head>
	<body>
		<div class="container" style="width:300px;padding-top:150px">
			<form class="form-signin">
		        <h2 class="form-signin-heading" >预警分析系统后台</h2>
		        <label for="inputAccount" class="sr-only">帐号</label>
		        <input id="inputAccount" class="form-control" placeholder="帐号" value="admin">
		        <label for="inputPassword" class="sr-only">密码</label>
		        <input type="password" id="inputPassword" class="form-control" placeholder="密码"  value="admin">
		        <div class="checkbox">
		          <label>
		            <input type="checkbox" value="remember-me">记住密码
		          </label>
		        </div>
		        <button class="btn btn-lg btn-primary btn-block" id="login">登陆</button>
	      	</form>
		</div>
		<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
		<script src="https://cdn.bootcss.com/jquery-cookie/1.4.1/jquery.cookie.js"></script>
		<script src="https://cdn.bootcss.com/Base64/1.0.1/base64.js"></script>
		<script>
			$('#login').on('click',function(){	
				var data={
					'name':$("#inputAccount").val(),
					'pwd':$("#inputPassword").val()
				}
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/api/v1/login",
					data:JSON.stringify(data),
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(res){
						if(res.code==200){
							//alert("登录成功")
							//设置缓存
							$.cookie('token', res.data.token, { expires: 0.5 ,path: '/'});
							window.location.href="/v1/warn_monitor"																			
						}else{
							alert("账户密码错误")
						}						
					}					
				});	
				return false;						
		  	});
		</script>
	</body>
</html>
