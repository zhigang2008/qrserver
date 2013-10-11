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
define("esri/tasks/RelationParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.RelationParameters",geometries1:null,geometries2:null,relation:null,relationParam:null,toJson:function(){var _9=_3.map(this.geometries1,function(_a){return _a.toJson();});var _b=_3.map(this.geometries2,function(_c){return _c.toJson();});var _d={};var _e=this.geometries1;if(_e&&_e.length>0){_d.geometries1=_4.toJson({geometryType:_7.getJsonType(_e[0]),geometries:_9});var _f=this.geometries1[0].spatialReference;_d.sr=_f.wkid?_f.wkid:_4.toJson(_f.toJson());}var _10=this.geometries2;if(_10&&_10.length>0){_d.geometries2=_4.toJson({geometryType:_7.getJsonType(_10[0]),geometries:_b});}if(this.relation){_d.relation=this.relation;}if(this.relationParam){_d.relationParam=_4.toJson(this.relationParam);}return _d;}});_2.mixin(_8,{SPATIAL_REL_CROSS:"esriGeometryRelationCross",SPATIAL_REL_DISJOINT:"esriGeometryRelationDisjoint",SPATIAL_REL_IN:"esriGeometryRelationIn",SPATIAL_REL_INTERIORINTERSECTION:"esriGeometryRelationInteriorIntersection",SPATIAL_REL_INTERSECTION:"esriGeometryRelationIntersection",SPATIAL_REL_COINCIDENCE:"esriGeometryRelationLineCoincidence",SPATIAL_REL_LINETOUCH:"esriGeometryRelationLineTouch",SPATIAL_REL_OVERLAP:"esriGeometryRelationOverlap",SPATIAL_REL_POINTTOUCH:"esriGeometryRelationPointTouch",SPATIAL_REL_TOUCH:"esriGeometryRelationTouch",SPATIAL_REL_WITHIN:"esriGeometryRelationWithin",SPATIAL_REL_RELATION:"esriGeometryRelationRelation"});if(_5("extend-esri")){_2.setObject("tasks.RelationParameters",_8,_6);}return _8;});