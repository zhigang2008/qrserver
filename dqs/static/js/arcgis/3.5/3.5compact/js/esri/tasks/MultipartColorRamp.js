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
define("esri/tasks/MultipartColorRamp",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/symbols/Symbol","esri/tasks/ColorRamp"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.tasks.MultipartColorRamp",type:"multipart",constructor:function(){this.colorRamps=[];},addColorRamp:function(_8){this.colorRamps.push(_8);},toJson:function(){var _9=_3.map(this.colorRamps,function(_a){return _a.toJson();});var _b={type:"multipart",colorRamps:_9};return _b;}});if(_4("extend-esri")){_2.setObject("tasks.MultipartColorRamp",_7,_5);}return _7;});