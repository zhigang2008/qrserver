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
define("esri/tasks/ProjectParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/lang","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.tasks.ProjectParameters",geometries:null,outSR:null,transformation:null,transformForward:null,toJson:function(){var _a=_3.map(this.geometries,function(_b){return _b.toJson();});var _c={};_c.outSR=this.outSR.wkid||_4.toJson(this.outSR.toJson());_c.inSR=this.geometries[0].spatialReference.wkid||_4.toJson(this.geometries[0].spatialReference.toJson());_c.geometries=_4.toJson({geometryType:_8.getJsonType(this.geometries[0]),geometries:_a});if(this.transformation){_c.transformation=this.transformation.wkid||_4.toJson(this.transformation);}if(_7.isDefined(this.transformForward)){_c.transformForward=this.transformForward;}return _c;}});if(_5("extend-esri")){_2.setObject("tasks.ProjectParameters",_9,_6);}return _9;});