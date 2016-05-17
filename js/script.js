window.onload = function(){
var arr = [true,true,true,true];
var editorarr = ["htmleditor","csseditor","jseditor","resouter"];
var currentold = [4,0];
var sizeclass = ["col-sm-","col-md-","col-lg-"];
function resize(){
	var i,j;
	for(i=0;i<4;i++)
		for(j=0;j<3;j++){
			$("#"+editorarr[i]).removeClass(sizeclass[j]+Math.round((12/currentold[1])).toString());
			$("#"+editorarr[i]).addClass(sizeclass[j]+Math.round((12/currentold[0])).toString());
		}
			
}
function handle( i){
	arr[i] = !arr[i];
	if(!arr[i]) {
		$("#"+editorarr[i]).hide();
		currentold[1] = currentold[0];
		currentold[0]--;
	} else {
		$("#"+editorarr[i]).show();
		currentold[1] = currentold[0];
                currentold[0]++;
	}
	resize();
}
$("#htmlbtn").on('click', function() {
                handle(0);
        }
);
$("#cssbtn").on('click',function() {
                handle(1);
        }
);
$("#jsbtn").on('click', function() {
                handle(2);
        }
);
$("#outputbtn").on('click', function() {
                handle(3);
        }
);
}
$("#result").on('click',function(event){
	var html = $("#htmlcode").val();
	var css = $("#csscode").val();
	var js = $("#jscode").val();
	var result = html + "<style>" + css +"</style>"+"<script>"+js+"</script>";
	$("#result1").html(result);
});
