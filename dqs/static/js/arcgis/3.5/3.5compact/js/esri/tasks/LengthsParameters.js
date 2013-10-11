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
define("esri/tasks/LengthsParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.tasks.LengthsParameters",polylines:null,lengthUnit:null,geodesic:null,calculationType:null,toJson:function(){var _8=_3.map(this.polylines,function(_9){return _9.toJson();});var _a={};_a.polylines=_4.toJson(_8);var _b=this.polylines[0].spatialReference;_a.sr=_b.wkid?_b.wkid:_4.toJson(_b.toJson());if(this.lengthUnit){_a.lengthUnit=this.lengthUnit;}if(this.geodesic){_a.geodesic=this.geodesic;}if(this.calculationType){_a.calculationType=this.calculationType;}return _a;}});if(_5("extend-esri")){_2.setObject("tasks.LengthsParameters",_7,_6);}return _7;});