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
define("esri/geometry/Geometry",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/SpatialReference"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.geometry.Geometry",spatialReference:null,type:null,setSpatialReference:function(sr){this.spatialReference=sr;return this;},verifySR:function(){if(!this.spatialReference){this.setSpatialReference(new _5(4326));}},getExtent:function(){return null;}});if(_3("extend-esri")){_2.setObject("geometry.Geometry",_6,_4);}return _6;});