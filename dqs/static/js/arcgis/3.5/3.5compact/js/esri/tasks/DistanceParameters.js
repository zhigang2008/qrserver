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
define("esri/tasks/DistanceParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.tasks.DistanceParameters",geometry1:null,geometry2:null,distanceUnit:null,geodesic:null,toJson:function(){var _8={};var _9=this.geometry1;if(_9){_8.geometry1=_3.toJson({geometryType:_6.getJsonType(_9),geometry:_9});}var _a=this.geometry2;if(_a){_8.geometry2=_3.toJson({geometryType:_6.getJsonType(_a),geometry:_a});}_8.sr=_3.toJson(this.geometry1.spatialReference.toJson());if(this.distanceUnit){_8.distanceUnit=this.distanceUnit;}if(this.geodesic){_8.geodesic=this.geodesic;}return _8;}});if(_4("extend-esri")){_2.setObject("tasks.DistanceParameters",_7,_5);}return _7;});