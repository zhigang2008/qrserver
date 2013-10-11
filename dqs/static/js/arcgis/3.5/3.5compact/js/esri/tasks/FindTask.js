/*
 COPYRIGHT 2009 ESRI

 TRADE SECRETS: ESRI PROPRIETARY AND CONFIDENTIAL
 Unpublished material - all rights reserved under the
 Copyright Laws of the United States and applicable international
 laws, treaties, and conventions.

 For additional information, contact:
 Environmental Systems Research Institute, Inc.
 Attn: Contracts and Legal Services Department
 380 New York Street
 Redlands, California, 92373
 USA

 email: contracts@esri.com
 */
//>>built
define("esri/tasks/FindTask",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/Deferred","dojo/has","esri/kernel","esri/request","esri/deferredUtils","esri/tasks/Task","esri/tasks/FindResult"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1(_9,{declaredClass:"esri.tasks.FindTask",constructor:function(_c,_d){this._url.path+="/find";this._handler=_2.hitch(this,this._handler);this.gdbVersion=_d&&_d.gdbVersion;this.registerConnectEvents("esri.tasks.FindTask",{"complete":["results"]});},_handler:function(_e,io,_f,_10,dfd){try{var _11=[];_3.forEach(_e.results,function(_12,i){_11[i]=new _a(_12);});this._successHandler([_11],"onComplete",_f,dfd);}catch(err){this._errorHandler(err,_10,dfd);}},execute:function(_13,_14,_15){var _16=this._encode(_2.mixin({},this._url.query,{f:"json"},_13.toJson())),_17=this._handler,_18=this._errorHandler;if(this.gdbVersion){_16.gdbVersion=this.gdbVersion;}var dfd=new _4(_8._dfdCanceller);dfd._pendingDfd=_7({url:this._url.path,content:_16,callbackParamName:"callback",load:function(r,i){_17(r,i,_14,_15,dfd);},error:function(r){_18(r,_15,dfd);}});return dfd;},onComplete:function(){}});if(_5("extend-esri")){_2.setObject("tasks.FindTask",_b,_6);}return _b;});