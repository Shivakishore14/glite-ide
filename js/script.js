window.onload = function(){
	var Base64={_keyStr:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",encode:function(e){var t="";var n,r,i,s,o,u,a;var f=0;e=Base64._utf8_encode(e);while(f<e.length){n=e.charCodeAt(f++);r=e.charCodeAt(f++);i=e.charCodeAt(f++);s=n>>2;o=(n&3)<<4|r>>4;u=(r&15)<<2|i>>6;a=i&63;if(isNaN(r)){u=a=64}else if(isNaN(i)){a=64}t=t+this._keyStr.charAt(s)+this._keyStr.charAt(o)+this._keyStr.charAt(u)+this._keyStr.charAt(a)}return t},decode:function(e){var t="";var n,r,i;var s,o,u,a;var f=0;e=e.replace(/[^A-Za-z0-9+/=]/g,"");while(f<e.length){s=this._keyStr.indexOf(e.charAt(f++));o=this._keyStr.indexOf(e.charAt(f++));u=this._keyStr.indexOf(e.charAt(f++));a=this._keyStr.indexOf(e.charAt(f++));n=s<<2|o>>4;r=(o&15)<<4|u>>2;i=(u&3)<<6|a;t=t+String.fromCharCode(n);if(u!=64){t=t+String.fromCharCode(r)}if(a!=64){t=t+String.fromCharCode(i)}}t=Base64._utf8_decode(t);return t},_utf8_encode:function(e){e=e.replace(/rn/g,"n");var t="";for(var n=0;n<e.length;n++){var r=e.charCodeAt(n);if(r<128){t+=String.fromCharCode(r)}else if(r>127&&r<2048){t+=String.fromCharCode(r>>6|192);t+=String.fromCharCode(r&63|128)}else{t+=String.fromCharCode(r>>12|224);t+=String.fromCharCode(r>>6&63|128);t+=String.fromCharCode(r&63|128)}}return t},_utf8_decode:function(e){var t="";var n=0;var r=c1=c2=0;while(n<e.length){r=e.charCodeAt(n);if(r<128){t+=String.fromCharCode(r);n++}else if(r>191&&r<224){c2=e.charCodeAt(n+1);t+=String.fromCharCode((r&31)<<6|c2&63);n+=2}else{c2=e.charCodeAt(n+1);c3=e.charCodeAt(n+2);t+=String.fromCharCode((r&15)<<12|(c2&63)<<6|c3&63);n+=3}}return t}}
	var th = $("#test").height();
	$("#test").hide();
	function init(){
		var h = $("#htmltitle").offset().top;
		var h1 = th-h-h/2;
		htmleditor.setSize("100%",h1);
		csseditor.setSize("100%",h1);
		jseditor.setSize("100%",h1);
		$("#result1").height(h1);
		$("#result2").height(th);
	}
	var withbox = false;
	var arr = [true,true,true,true];
	var editorarr = ["htmleditor","csseditor","jseditor","resouter"];
	var currentold = [3,0];
	var theme=["default","ambiance","the-matrix","neat"];
	var sizeclass = ["col-sm-","col-md-","col-lg-"];
	function resize(){
		var i,j;
		for(i=0;i<4;i++)
			for(j=0;j<3;j++){
				if(i != 3 || withbox){
					$("#"+editorarr[i]).removeClass(sizeclass[j]+Math.round((12/currentold[1])).toString());
					$("#"+editorarr[i]).addClass(sizeclass[j]+Math.round((12/currentold[0])).toString());
				}
			}		
	}
	function handle( i){
		arr[i] = !arr[i];
		if(i != 3 || withbox){
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
		}else{
			if(!arr[i])
				$("#res2outer").hide();
			else
				$("#res2outer").show();	
			init();		
		}
	}	
	$("#htmlbtn").on('click', function() {
        	handle(0);
     	});
	$("#cssbtn").on('click',function() {
                handle(1);
        });
	$("#jsbtn").on('click', function() {
                handle(2);
        });
	$("#outputbtn").on('click', function() {
                handle(3);
        });
	//alert("hi");
	$("#changeres2").on('click', function() {
		withbox = true;
		currentold[1] = currentold[0];
		currentold[0]++;
		$("#resouter").show();	
		$("#res2outer").hide()
		resize();
	});
	$("#changeres1").on('click', function() {
		withbox = false;
		currentold[1] = currentold[0];
		currentold[0]--;
		$("#resouter").hide();	
		$("#res2outer").show()
		resize();
	});
	var htmleditor = CodeMirror.fromTextArea(document.getElementById("htmlcode"), {
		lineNumbers: true,
		extraKeys: {"Ctrl-Space": "autocomplete"},
		matchBrackets: true,
		mode: "htmlmixed"
  	});
	var csseditor = CodeMirror.fromTextArea(document.getElementById("csscode"), {
    		lineNumbers: true,
		extraKeys: {"Ctrl-Space": "autocomplete"},
		matchBrackets: true,		
		mode: "css"
		});
	var jseditor = CodeMirror.fromTextArea(document.getElementById("jscode"), {
		lineNumbers: true,
		extraKeys: {"Ctrl-Space": "autocomplete"},
		matchBrackets: true,
		mode: {name: "javascript", globalVars: true}
	});
	$("#result").on('click',function(event){
		htmleditor.save();
		csseditor.save();
		jseditor.save();
		var html = document.getElementById("htmlcode").value;
		var css = document.getElementById("csscode").value;
		var js = document.getElementById("jscode").value;
		var result = html + "<style>" + css +"</style>"+"<script>"+js+"</script>";
		var data_url = "data:text/html;charset=utf-8;base64," + Base64.encode(result);
		if(withbox)
        		document.getElementById("result1").src = data_url;
		else
			document.getElementById("result2").src = data_url;
		var hash = "#result2"
		$('html, body').animate({
			scrollTop: $(hash).offset().top
		}, 900, function(){
			window.location.hash = hash;
 		});
	});
  	function setTheme(a){
		htmleditor.setOption("theme", a);
		csseditor.setOption("theme", a);
		jseditor.setOption("theme", a);	
	}	
	setTheme("default");
	init();	
}

