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
define("esri/tasks/GenerateRendererTask",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/_base/Deferred","dojo/has","esri/kernel","esri/request","esri/deferredUtils","esri/renderers/jsonUtils","esri/tasks/Task","dojo/has!extend-esri?esri/tasks/GenerateRendererParameters","dojo/has!extend-esri?esri/tasks/ClassificationDefinition","dojo/has!extend-esri?esri/tasks/ClassBreaksDefinition","dojo/has!extend-esri?esri/tasks/UniqueValueDefinition","dojo/has!extend-esri?esri/tasks/ColorRamp","dojo/has!extend-esri?esri/tasks/AlgorithmicColorRamp","dojo/has!extend-esri?esri/tasks/MultipartColorRamp"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1(_a,{declaredClass:"esri.tasks.GenerateRendererTask",constructor:function(_c,_d){this.url=_c;this._url.path+="/generateRenderer";this._handler=_2.hitch(this,this._handler);this.source=_d&&_d.source;this.gdbVersion=_d&&_d.gdbVersion;this.registerConnectEvents("esri.tasks.GenerateRendererTask",{"complete":["renderer"]});},_handler:function(_e,io,_f,_10,dfd){try{var _11=_9.fromJson(_e);if(_e.type==="classBreaks"){_11.setMaxInclusive(true);}this._successHandler([_11],"onComplete",_f,dfd);}catch(err){this._errorHandler(err,_10,dfd);}},execute:function(_12,_13,_14){var _15=_2.mixin(_12.toJson(),{f:"json"}),_16=this._handler,_17=this._errorHandler;if(this.source){var _18={source:this.source.toJson()};_15.layer=_3.toJson(_18);}if(this.gdbVersion){_15.gdbVersion=this.gdbVersion;}var dfd=new _4(_8._dfdCanceller);dfd._pendingDfd=_7({url:this._url.path,content:_15,callbackParamName:"callback",load:function(r,i){_16(r,i,_13,_14,dfd);},error:function(r){_17(r,_14,dfd);}});return dfd;},onComplete:function(){}});if(_5("extend-esri")){_2.setObject("tasks.GenerateRendererTask",_b,_6);}return _b;});