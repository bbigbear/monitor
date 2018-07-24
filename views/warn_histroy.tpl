<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>预警历史</title>
<link rel="stylesheet" href="/static/css/layui.css">
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
		<form class="layui-form layui-form-pane1" action="">
		<div class="layui-form-item">
			<div class="layui-inline">
			    <label class="layui-form-label">预警类型</label>
			    <div class="layui-input-inline" style="width: 150px;">
			      <select name="PlanClass" id="style" lay-filter="status_select">
					    <option value="全部" > 全部</option>
						<option value="学生上网预警" > 学生上网预警</option>
						<option value="学生学业预警" > 学生学业预警</option>
						<option value="图书借阅预警" > 图书借阅预警</option>
						<option value="学生沉迷预警" > 学生沉迷预警</option>
						<option value="一卡通消费预警" > 一卡通消费预警</option>
						<option value="学生紧急预警" > 学生紧急预警</option>
						<option value="贫困生预警" > 贫困生预警</option>
						<option value="图书馆爆满预警" > 图书馆爆满预警</option>
						<option value="学生挂科预警" > 学生挂科预警</option>
						<option value="学生失联预警" > 学生失联预警</option>
			      </select>
			    </div>
		  	</div>
			<!--<div class="layui-inline">
			    <label class="layui-form-label">开始时间</label>
			    <div class="layui-input-inline" style="width: 150px;">
			      	<input type="text" id="PlanId" autocomplete="off" class="layui-input">
			    </div>
		  	</div>
			<div class="layui-inline">
			    <label class="layui-form-label">结束时间</label>
			    <div class="layui-input-inline" style="width: 150px;">
			    	<input type="text" id="PlanId" autocomplete="off" class="layui-input">
			    </div>
		  	</div>-->
			<div class="layui-inline">
			    <label class="layui-form-label">预警状态</label>
			    <div class="layui-input-inline" style="width: 150px;">
			      <select name="PlanClass" id="status" lay-filter="status_select">
					    <option value="全部" > 全部</option>
						<option value="已处理" >已处理</option>
						<option value="待处理" >待处理</option>
			      </select>
			    </div>
		  	</div>
			<button class="layui-btn" id="query">查询</button>
		</div>		
		</form>
		<hr class="layui-bg-green">
		<i class="layui-icon layui-icon-delete" style="font-size: 20px; color: #FF5722" id="del"></i>
		<table id="dishList" lay-filter="room"></table>
		<script type="text/html" id="barDemo">
			{{#  if(d.Status =="待处理"){ }}
				<a class="layui-btn layui-btn-normal layui-btn-xs" lay-event="handle">已处理</a>
			{{# } }}
		</script>
		<hr class="layui-bg-green">		
	</div>
  </div>
  
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    ©2018 奚米. All Rights Reserved
  </div>
</div>	
<script src="/static/layui.js"></script>
<!--<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>-->
<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/jquery-cookie/1.4.1/jquery.cookie.js"></script>
<script src="https://cdn.bootcss.com/Base64/1.0.1/base64.js"></script>
<script>
	//JavaScript代码区域
	layui.use(['element','layer','jquery','table'], function(){
	  var element = layui.element
		,form=layui.form
		,layer=layui.layer
		,$=layui.jquery
		,table=layui.table;
	  //layer.msg("你好");	
	  //table 渲染
	  table.render({
	    elem: '#dishList'
	    ,height: 400
	    ,url: '/v1/warn/getwarndata?token='+$.cookie('token')//数据接口
	    ,page: true //开启分页
		,id: 'listReload'
		,size: 'lg'
	    ,cols: [[ //表头
		  {type:'checkbox'}
		  ,{field:'Sname', title:'姓名', width:100}
		  ,{field:'WarnName',  title:'预警类型', width:120}
	      ,{field:'WarnTime',  title:'预警时间', width:200}
		  ,{field:'WarnInfo',  title:'预警描述', width:250}
		  ,{field:'Remark',  title:'备注', width:250, event:'setSign',style:'cursor: pointer;'}
		  ,{field:'Status',  title:'预警状态', width:120}
		  ,{fixed: 'right', title:'操作',width:150, align:'center', toolbar: '#barDemo'}
	    ]]
	  });
		//监听工具条
		table.on('tool(room)', function(obj){ //注：tool是工具条事件名，test是table原始容器的属性 lay-filter="对应的值"
		    var data = obj.data //获得当前行数据
		    ,layEvent = obj.event; //获得 lay-event 对应的值
		    if(layEvent === 'handle'){	
				layer.confirm('确定为已处理？', function(index){
			        var jsData={'id':data.Id,'status':"已处理"}
					$.post('/v1/warn/change', jsData, function (out) {
		                if (out.code == 200) {
		                    layer.alert('已处理', {icon: 1},function(index){
		                        layer.close(index);
		                        location.reload();
		                    });
		                } else {
		                    layer.msg(out.message)
		                }
		            }, "json");
			        layer.close(index);
		      	});
	    	}else if(layEvent === 'setSign'){
		      layer.prompt({
		        formType: 2
		        ,title: '修改备注内容'
		        ,value: data.sign
		      }, function(value, index){
		        layer.close(index);
		        
		        //这里一般是发送修改的Ajax请求
		        var jsData={'id':data.Id,'remark':value}
					$.post('/v1/warn/change_remark', jsData, function (out) {
		                if (out.code == 200) {
		                    layer.alert('已处理', {icon: 1},function(index){
		                        layer.close(index);
		                        location.reload();
		                    });
		                } else {
		                    layer.msg(out.message)
		                }
		            }, "json");
			    layer.close(index);
		        //同步更新表格和缓存对应的值
		        //obj.update({
		          //sign: value
		        //});
		      });
		    }
	 	});	
		//查询
	 	$('#query').on('click',function(){ 
           	var style = $("#style").val()
			var status = $("#status").val()
			if (style=="全部"){
				style=""
			}
			if (status=="全部"){
				status=""
			}
			table.reload('listReload', {
                where: {
                    style: style,
					status: status,
                }
            });
			return false;
		});	
	//批量删除
	$('#del').on('click',function(){
		var str="";
		var checkStatus=table.checkStatus('listReload')
		,data=checkStatus.data;
		if(data.length==0){
			layer.msg("请选择要删除信息");
		}else{
			for(var i=0;i<data.length;i++){
				str+=data[i].Id+",";
			}
			layer.confirm('是否删除这'+data.length+'条数据?',{icon:3,title:'提示'},function(index){
				//window.location.href="/v1/delmultidata?id="+str+"";
				$.ajax({
					type:"POST",
					url:"/v1/warn/del",
					data:{
						id:str,
						token:$.cookie('token')
					},
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(res){
						if(res.code==200){
							layer.msg("删除成功");
							//重载表格
							table.reload('listReload', {							  
							});												
						}else{
							layer.msg("删除失败");
						}						
					}					
				});				
				layer.close(index);
			});
		}
		return false;
	});
	
  });
		
</script>

</body>
</html>