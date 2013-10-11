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
define("esri/tasks/ClosestFacilityTask",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/request","esri/geometry/normalizeUtils","esri/tasks/Task","esri/tasks/ClosestFacilitySolveResult"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(_7,{declaredClass:"esri.tasks.ClosestFacilityTask",constructor:function(_a){this._url.path+="/solveClosestFacility";this._handler=_2.hitch(this,this._handler);this.registerConnectEvents("esri.tasks.ClosestFacilityTask",{"solve-complete":["result"]});},__msigns:[{n:"solve",c:3,a:[{i:0,p:["incidents.features","facilities.features","pointBarriers.features","polylineBarriers.features","polygonBarriers.features"]}],e:2}],_handler:function(_b,io,_c,_d,_e){try{var _f=new _8(_b);this._successHandler([_f],"onSolveComplete",_c,_e);}catch(err){this._errorHandler(err,_d,_e);}},solve:function(_10,_11,_12,_13){var _14=_13.assembly,_15=this._encode(_2.mixin({},this._url.query,{f:"json"},_10.toJson(_14&&_14[0]))),_16=this._handler,_17=this._errorHandler;return _5({url:this._url.path,content:_15,callbackParamName:"callback",load:function(r,i){_16(r,i,_11,_12,_13.dfd);},error:function(r){_17(r,_12,_13.dfd);}});},onSolveComplete:function(){}});_6._createWrappers(_9);if(_3("extend-esri")){_2.setObject("tasks.ClosestFacilityTask",_9,_4);}return _9;});