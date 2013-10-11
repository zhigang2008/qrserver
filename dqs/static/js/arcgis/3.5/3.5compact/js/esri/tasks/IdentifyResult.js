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
define("esri/tasks/IdentifyResult",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/graphic","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.tasks.IdentifyResult",constructor:function(_8){_2.mixin(this,_8);this.feature=new _5(_8.geometry?_6.fromJson(_8.geometry):null,null,_8.attributes);delete this.geometry;delete this.attributes;}});if(_3("extend-esri")){_2.setObject("tasks.IdentifyResult",_7,_4);}return _7;});