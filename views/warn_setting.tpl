<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>预警设置</title>
<link rel="stylesheet" href="/static/css/layui.css">
<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
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
			<li class="layui-this">设置预警通知方式</li>
		    <li>学业预警</li>
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
					    	<label class="layui-form-label" style="width:auto;">预警通知方式</label>
					    	<div class="layui-input-inline">
					     	 	<select name="PlanClass" id="Tzfs" lay-filter="status_select">
								    <option value="短信" > 短信</option>
									<option value="微信" > 微信</option>
									<option value="邮件" > 邮件</option>
									<option value="站内信" > 站内信</option>
						      	</select>
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
					     	 	<input type="text" id="number" autocomplete="off" class="layui-input" value="学生学业预警">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label" >总成绩低于</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="XyyjZcj" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">分数</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label" style="width:auto">单科成绩低于</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="XyyjDkcj" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">分数</div>
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
					     	 	<input type="text" id="YktDbxf" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(元)</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">当日消费</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="YktDrxf" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(元)</div>
					  	</div>
					</div>	
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">单月消费</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="YktDyxf" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(元)</div>
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
					     	 	<input type="text" id="SwyjCs" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">上网总时长</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="SwyjZsc" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(小时)</div>
					  	</div>
					</div>	
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label" style="width:auto">单次上网时长</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="SwyjDcswsc" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(小时)</div>
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
					     	 	<input type="text" id="TsjySl" autocomplete="off" class="layui-input">
					   	 	</div>
					  	</div>
					</div>
					<div class="layui-form-item">
					  	<div class="layui-inline">
					    	<label class="layui-form-label">借阅时间</label>
					    	<div class="layui-input-inline">
					     	 	<input type="text" id="TsjySj" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(小时)</div>
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
					     	 	<input type="text" id="PksYxf" autocomplete="off" class="layui-input">
					   	 	</div>
							<div class="layui-form-mid layui-word-aux">单位(元)</div>
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
					     	 	<input type="text" id="TsgbmRs" autocomplete="off" class="layui-input">
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
					     	 	<input type="text" id="Gkcs" autocomplete="off" class="layui-input">
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
    ©2018 奚米. All Rights Reserved
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
		<<<range .m>>>
			$("#Tzfs").val(<<<.Tzfs>>>)
			$("#XyyjZcj").val(<<<.XyyjZcj>>>)
			$("#XyyjDkcj").val(<<<.XyyjDkcj>>>)
			$("#YktDbxf").val(<<<.YktDbxf>>>)
			$("#YktDrxf").val(<<<.YktDrxf>>>)
			$("#YktDyxf").val(<<<.YktDyxf>>>)
			$("#SwyjCs").val(<<<.SwyjCs>>>)
			$("#SwyjZsc").val(<<<.SwyjZsc>>>)
			$("#SwyjDcswsc").val(<<<.SwyjDcswsc>>>)
			$("#TsjySl").val(<<<.TsjySl>>>)
			$("#TsjySj").val(<<<.TsjySj>>>)
			$("#PksYxf").val(<<<.PksYxf>>>)
			$("#TsgbmRs").val(<<<.TsgbmRs>>>)
			$("#Gkcs").val(<<<.Gkcs>>>)
		<<<end>>>		
	});	
	var style=""
	var list = ["设置预警通知方式","学业预警","一卡通消费预警","上网预警","图书借阅预警","贫困生预警","图书馆爆满预警","挂科预警"]
	//tab监听
	layui.use('element', function(){
	  var element = layui.element;  
	  //一些事件监听
	  element.on('tab(demo)', function(data){
		console.log(data.index)
	    style=list[data.index]
	  });
	});	
	//点击新增按钮
	$('#set').on('click',function(){
		if(style==""){
			style="设置预警通知方式"
		}
		var data={
			'style':style,
			'Tzfs':$("#Tzfs").val(),
			'XyyjZcj':parseFloat($("#XyyjZcj").val()),
			'XyyjDkcj':parseFloat($("#XyyjDkcj").val()),
			'YktDbxf':parseFloat($("#YktDbxf").val()),
			'YktDrxf':parseFloat($("#YktDrxf").val()),
			'YktDyxf':parseFloat($("#YktDyxf").val()),
			'SwyjCs':parseInt($("#SwyjCs").val()),
			'SwyjZsc':parseFloat($("#SwyjZsc").val()),
			'SwyjDcswsc':parseFloat($("#SwyjDcswsc").val()),
			'TsjySl':parseInt($("#TsjySl").val()),
			'TsjySj':parseFloat($("#TsjySj").val()),
			'PksYxf':parseFloat($("#PksYxf").val()),
			'TsgbmRs':parseInt($("#TsgbmRs").val()),
			'Gkcs':parseInt($("#Gkcs").val())
		};
			$.ajax({
				type:"POST",
				contentType:"application/json;charset=utf-8",
				url:"/v1/warn/updata",
				data:JSON.stringify(data),
				async:false,
				error:function(request){
					alert("post error")						
				},
				success:function(res){
					if(res.code==200){
						alert("设置成功")
						window.location.reload();						
					}else{
						layer.msg("设置失败")
					}						
				}
			});
			return false;
		});		  	
  	});
		
</script>

</body>
</html>