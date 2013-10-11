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
define("esri/tasks/RasterData",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.tasks.RasterData",constructor:function(_6){if(_6){_2.mixin(this,_6);}},url:null,format:null,itemID:null,toJson:function(){var _7={};if(this.url){_7.url=this.url;}if(this.format){_7.format=this.format;}if(this.itemID){_7.itemID=this.itemID;}return _7;}});if(_3("extend-esri")){_2.setObject("tasks.RasterData",_5,_4);}return _5;});