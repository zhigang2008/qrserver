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
define("esri/tasks/AreasAndLengthsParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.tasks.AreasAndLengthsParameters",polygons:null,lengthUnit:null,areaUnit:null,calculationType:null,toJson:function(){var _8=_3.map(this.polygons,function(_9){return _9.toJson();});var _a={};_a.polygons=_4.toJson(_8);var _b=this.polygons[0].spatialReference;_a.sr=_b.wkid?_b.wkid:_4.toJson(_b.toJson());if(this.lengthUnit){_a.lengthUnit=this.lengthUnit;}if(this.areaUnit){if(_2.isString(this.areaUnit)){_a.areaUnit=_4.toJson({"areaUnit":this.areaUnit});}else{_a.areaUnit=this.areaUnit;}}if(this.calculationType){_a.calculationType=this.calculationType;}return _a;}});if(_5("extend-esri")){_2.setObject("tasks.AreasAndLengthsParameters",_7,_6);}return _7;});