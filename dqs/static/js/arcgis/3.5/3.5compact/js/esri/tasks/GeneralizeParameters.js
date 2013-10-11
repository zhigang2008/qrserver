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
define("esri/tasks/GeneralizeParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.GeneralizeParameters",geometries:null,deviationUnit:null,maxDeviation:null,toJson:function(){var _9=_3.map(this.geometries,function(_a){return _a.toJson();});var _b={};if(this.geometries&&this.geometries.length>0){_b.geometries=_4.toJson({geometryType:_7.getJsonType(this.geometries[0]),geometries:_9});_b.sr=_4.toJson(this.geometries[0].spatialReference.toJson());}if(this.deviationUnit){_b.deviationUnit=this.deviationUnit;}if(this.maxDeviation){_b.maxDeviation=this.maxDeviation;}return _b;}});if(_5("extend-esri")){_2.setObject("tasks.GeneralizeParameters",_8,_6);}return _8;});