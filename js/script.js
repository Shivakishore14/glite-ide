window.onload = function() {
//	$("#inputblock").hide();
//	var ww = $("#prog").width();	
//	$("#prog").width(ww-ww/20);	
}
$("#uploadfile , #uploadtext").on('click',function(event) {
	$("#myfile").trigger("click");
});
$("#mybtn").click(function(){
	var a = $("#myfile").val();
	$("#status").html("Uploading...");
	$("#myfile").upload("files/hello.py",function(success){
		$("#status").html(a+" uploaded");
	},$("#prog"));
});
$("#result").on('click',function(event){
	var html = $("#htmlcode").val();
	var css = $("#csscode").val();
	var js = $("#jscode").val();
	var result = html + "<style>" + css +"</style>"+"<script>"+js+"</script>";
	$("#result1").html(result);
});
