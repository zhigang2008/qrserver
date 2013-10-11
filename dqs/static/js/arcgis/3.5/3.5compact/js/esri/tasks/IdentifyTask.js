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
define("esri/tasks/IdentifyTask",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/request","esri/geometry/normalizeUtils","esri/tasks/Task","esri/tasks/IdentifyResult"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a=_1(_8,{declaredClass:"esri.tasks.IdentifyTask",constructor:function(_b,_c){this._url.path+="/identify";this._handler=_2.hitch(this,this._handler);this.gdbVersion=_c&&_c.gdbVersion;this.registerConnectEvents("esri.tasks.IdentifyTask",{"complete":["results"]});},__msigns:[{n:"execute",c:3,a:[{i:0,p:["geometry"]}],e:2}],_handler:function(_d,io,_e,_f,dfd){try{var _10=[];_3.forEach(_d.results,function(_11,i){_10[i]=new _9(_11);});this._successHandler([_10],"onComplete",_e,dfd);}catch(err){this._errorHandler(err,_f,dfd);}},execute:function(_12,_13,_14,_15){var _16=_15.assembly,_17=this._encode(_2.mixin({},this._url.query,{f:"json"},_12.toJson(_16&&_16[0]))),_18=this._handler,_19=this._errorHandler;if(this.gdbVersion){_17.gdbVersion=this.gdbVersion;}return _6({url:this._url.path,content:_17,callbackParamName:"callback",load:function(r,i){_18(r,i,_13,_14,_15.dfd);},error:function(r){_19(r,_14,_15.dfd);}});},onComplete:function(){}});_7._createWrappers(_a);if(_4("extend-esri")){_2.setObject("tasks.IdentifyTask",_a,_5);}return _a;});