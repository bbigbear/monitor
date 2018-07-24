<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>预警设置</title>
<link rel="stylesheet" href="/static/css/layui.css">
<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/jquery-cookie/1.4.1/jquery.cookie.js"></script>
<script src="https://cdn.bootcss.com/Base64/1.0.1/base64.js"></script>
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  <div class="layui-header">
    <div class="layui-logo" style="width:auto;padding-left:10px;font-size:20px;">预警分析系统</div>
    <ul class="layui-nav layui-layout-right">
      <li class="layui-nav-item">
        <a href="javascript:;">
          <img src="../static/img/admin.jpg" class="layui-nav-img">
          学校管理员
        </a>
        <dl class="layui-nav-child">
          <dd><a href="">基本资料</a></dd>
          <dd><a href="">安全设置</a></dd>
        </dl>
      </li>
      <li class="layui-nav-item"><a href="/login">退出</a></li>
    </ul>
  </div>
  
  <div class="layui-side layui-bg-black">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
      <ul class="layui-nav layui-nav-tree"  lay-filter="test">
        <li class="layui-nav-item"><a href="/v1/warn_monitor">预警监控</a></li>
        <li class="layui-nav-item"><a href="/v1/warn_histroy">预警历史</a></li>
        <li class="layui-nav-item"><a href="/v1/warn_setting">预警配置</a></li>
      </ul>
    </div>
  </div>
  <div class="layui-body">
    <!-- 内容主体区域 -->
    <div style="padding: 15px;">
		<div class="layui-tab layui-tab-card" lay-filter="demo">
		  <ul class="layui-tab-title">
		    <li class="layui-this">学业预警</li>
			<li>一卡通消费预警</li>
		    <li>上网预警</li>
		    <li>图书借阅预警</li>
			<li>贫困生预警</li>
			<li>图书馆爆满预警</li>
			<li>挂科预警</li>
		  </ul>
		  <div class="layui-tab-content" style="height:auto;">
		    <div class="layui-tab-item layui-show">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="学生学业预警">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label" >总成绩低于</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label" style="width:auto">单科成绩低于</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>								  
				</form>
			</div>
			<div class="layui-tab-item">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="一卡通消费预警">
					   	 	</div>
					  	</div>
					</div>					
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">单笔消费</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">当日消费</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>	
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">单月消费</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>	  
				</form>
			</div>
		    <div class="layui-tab-item">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="学生上网预警">
					   	 	</div>
					  	</div>
					</div>					
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">上网次数</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">上网总时长</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>	
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label" style="width:auto">单次上网时长</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>	  
				</form>
			</div>
		    <div class="layui-tab-item">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="图书借阅预警">
					   	 	</div>
					  	</div>
					</div>					
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">借阅数量</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">借阅时间</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>  
				</form>
			</div>
		    <div class="layui-tab-item">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="贫困生预警">
					   	 	</div>
					  	</div>
					</div>					
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">月消费金额</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>  
				</form>
			</div>
		    <div class="layui-tab-item">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="图书馆爆满预警">
					   	 	</div>
					  	</div>
					</div>					
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">人数</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>  
				</form>
			</div>
		    <div class="layui-tab-item">
				<form class="layui-form layui-form-pane1" action="">
				  	<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">预警类型</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="学生挂科预警">
					   	 	</div>
					  	</div>
					</div>					
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">挂科次数</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>  
				</form>
			</div>
		  </div>
		</div>
		<button class="layui-btn" id="set">设置</button>	
	</div>
  </div> 
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    ©2018 智慧校园. All Rights Reserved
  </div>
</div>

<script src="/static/layui.js"></script>
<!--<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>-->
<script>
	//JavaScript代码区域
	layui.use(['element','layer','jquery','table'], function(){
	  var element = layui.element
		,form=layui.form
		,layer=layui.layer
		,$=layui.jquery
		,table=layui.table;
	  //layer.msg("你好");
	//自动加载
	$(function(){			
	});	
	var style=""
	//tab监听
	layui.use('element', function(){
	  var element = layui.element;  
	  //一些事件监听
	  element.on('tab(demo)', function(data){
	    console.log(data.elem.context.textContent);
	  });
	});	
	//点击新增按钮
	$('#set').on('click',function(){
		
		var data={
			'name':name,
			'price':parseFloat(price),
			'dishType':$("#dishType").val(),
			'dishPicPath':path_src,
			'detail':detail,
			};
			$.ajax({
				type:"POST",
				contentType:"application/json;charset=utf-8",
				url:"/v1/restaurant_dish/add_action",
				data:JSON.stringify(data),
				async:false,
				error:function(request){
					alert("post error")						
				},
				success:function(res){
					if(res.code==200){
						alert("新增成功")
						window.location.reload();						
					}else{
						alert("新增失败")
					}						
				}
			});
		});		  	
  	});
		
</script>

</body>
</html>