//部分共用方法
function showSuccess(msg){
	$("#msgcontent").html(msg)
  $('#msg').removeClass("alert-info");
  $('#msg').removeClass("alert-warning");
  $('#msg').removeClass("alert-danger");
  $('#msg').addClass("alert-success");
  $('#msg').show()
}
function showInfo(msg){
	$("#msgcontent").html(msg)
  $('#msg').removeClass("alert-success");
  $('#msg').removeClass("alert-warning");
  $('#msg').removeClass("alert-danger");
  $('#msg').addClass("alert-info");
  $('#msg').show()
}
function showWarning(msg){
	$("#msgcontent").html(msg)
  $('#msg').removeClass("alert-success");
  $('#msg').removeClass("alert-info");
  $('#msg').removeClass("alert-danger");
  $('#msg').addClass("alert-warning");
  $('#msg').show()
}
function showError(msg){
	$("#msgcontent").html(msg)
  $('#msg').removeClass("alert-success");
  $('#msg').removeClass("alert-info");
  $('#msg').removeClass("alert-warning");
  $('#msg').addClass("alert-danger");
  $('#msg').show()
}