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
define("esri/tasks/DataLayer",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/geometry/jsonUtils","esri/tasks/SpatialRelationship"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.DataLayer",name:null,where:null,geometry:null,spatialRelationship:null,toJson:function(){var _9={type:"layer",layerName:this.name,where:this.where,spatialRel:this.spatialRelationship};var g=this.geometry;if(g){_9.geometryType=_6.getJsonType(g);_9.geometry=g.toJson();}return _5.filter(_9,function(_a){if(_a!==null){return true;}});}});_2.mixin(_8,_7);if(_3("extend-esri")){_2.setObject("tasks.DataLayer",_8,_4);}return _8;});