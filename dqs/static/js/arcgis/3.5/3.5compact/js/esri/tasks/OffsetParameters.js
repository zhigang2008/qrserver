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
define("esri/tasks/OffsetParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.OffsetParameters",geometries:null,bevelRatio:null,offsetDistance:null,offsetHow:null,offsetUnit:null,toJson:function(){var _9=_3.map(this.geometries,function(_a){return _a.toJson();});var _b={};if(this.geometries&&this.geometries.length>0){_b.geometries=_4.toJson({geometryType:_7.getJsonType(this.geometries[0]),geometries:_9});_b.sr=_4.toJson(this.geometries[0].spatialReference.toJson());}if(this.bevelRatio){_b.bevelRatio=this.bevelRatio;}if(this.offsetDistance){_b.offsetDistance=this.offsetDistance;}if(this.offsetHow){_b.offsetHow=this.offsetHow;}if(this.offsetUnit){_b.offsetUnit=this.offsetUnit;}return _b;}});_2.mixin(_8,{OFFSET_BEVELLED:"esriGeometryOffsetBevelled",OFFSET_MITERED:"esriGeometryOffsetMitered",OFFSET_ROUNDED:"esriGeometryOffsetRounded"});if(_5("extend-esri")){_2.setObject("tasks.OffsetParameters",_8,_6);}return _8;});