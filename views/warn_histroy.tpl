<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>监控历史</title>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="apple-mobile-web-app-status-bar-style" content="black"> 
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="format-detection" content="telephone=no">

  <link rel="stylesheet" href="/static/css/layui.css">

<style>
body{padding: 10px;}
</style>
</head>
<body>
<form class="layui-form layui-form-pane1" action="">
  <div class="layui-form-item">
  <div class="layui-inline">
    <label class="layui-form-label">学期</label>
    <div class="layui-input-inline" style="width: 150px;">
      <select name="Major" id="Major" lay-filter="status_select">
		    <option value="选择学期" > 选择学期</option>
			<option value="计算机软件" > 计算机软件</option>
      </select>
    </div>
  </div>
  <div class="layui-inline">
    <label class="layui-form-label">课程名称</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
  </div>
  </div>
  <div class="layui-form-item">
  <div class="layui-inline">
    <label class="layui-form-label">教案状态</label>
    <div class="layui-input-inline" style="width: 150px;">
      <select name="PlanClass" id="PlanClass" lay-filter="status_select">
		    <option value="已提交" > 已提交</option>
			<option value="2" > 2</option>
			<option value="3" > 3</option>
      </select>
    </div>
  </div>
  <div class="layui-inline">
    <label class="layui-form-label">教师名称</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
  </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-inline layui-layout-right" style="padding:10px;">
    	<button class="layui-btn" id="query">查询</button>
		<button class="layui-btn" id="clear">清除条件</button>
  	</div>
  </div>
</form>

<br><br>

	<table id="list" lay-filter="announcement" style="width:auto;"></table>
	<script type="text/html" id="barDemo">
		<a class="layui-btn layui-btn-normal layui-btn-xs" lay-event="edit">查看</a>
	</script>
<script src="/static/layui.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->

<script>
layui.use(['form','laydate','upload','jquery','layedit','element','table','laytpl'], function(){
  var form = layui.form
  ,laydate=layui.laydate
  ,upload = layui.upload
  , $ = layui.jquery
  ,layedit=layui.layedit
  ,element=layui.element
  ,table=layui.table
  ,laytpl = layui.laytpl;
	//自动加载
	$(function(){
				
	});	
	//table 渲染
	  table.render({
	    elem: '#list'
	    ,height: 315
	    ,url: '/v1/warn/getwarndata?status=未处理'//数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[   
	      {field:'Sname', title:'姓名', width:120}
		  ,{field:'WarnName',  title:'预警类型', width:120}
	      ,{field:'WarnTime',  title:'预警时间', width:120}
		  ,{field:'WarnInfo',  title:'预警描述', width:120}
	    ]]
	  });
	//监听工具条
		table.on('tool(announcement)', function(obj){ //注：tool是工具条事件名，test是table原始容器的属性 lay-filter="对应的值"
		    var data = obj.data //获得当前行数据
		    ,layEvent = obj.event; //获得 lay-event 对应的值
		    if(layEvent === 'edit'){
		      //layer.msg('查看操作');		
			  layer.open({
			  type: 2,
			  title: '查看课程',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/jxjh/look?id='+data.Id], //iframe的url，no代表不显示滚动条
			  cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
			  	return false; 
			  },
		});
	    }
	  });
  
	$('#add').on('click',function(){
		layer.open({
			  type: 2,
			  title: '新建计划',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/jxjh/add'], //iframe的url，no代表不显示滚动条
		});
		return false;
	});
	
	$('#query').on('click',function(){
		//alert("点击查询")
		var major=$("#Major").val()
		var grade=$("#PlanGrade").val()
		var classs=$("#PlanClass").val()
		var planId=$("#PlanId").val()
		table.reload('listReload',{
			where:{
				major:major,
				grade:grade,
				class:classs,
				planId:planId,
			}
		})
		return false;
	})
});
</script>

</body>
</html>