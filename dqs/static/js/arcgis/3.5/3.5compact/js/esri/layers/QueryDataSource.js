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
define("esri/layers/QueryDataSource",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/DataSource","esri/SpatialReference"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(_6,{declaredClass:"esri.layers.QueryDataSource",constructor:function(_9){if(_9){if(_9.oidFields&&_2.isString(_9.oidFields)){this.oidFields=_9.oidFields.split(",");}if(_9.spatialReference){this.spatialReference=new _7(_9.spatialReference);}}},toJson:function(){var _a={type:"queryTable",workspaceId:this.workspaceId,query:this.query,oidFields:this.oidFields&&this.oidFields.join(),spatialReference:this.spatialReference&&this.spatialReference.toJson()};if(this.geometryType){var _b;if(this.geometryType.toLowerCase()==="point"){_b="esriGeometryPoint";}else{if(this.geometryType.toLowerCase()==="multipoint"){_b="esriGeometryMultipoint";}else{if(this.geometryType.toLowerCase()==="polyline"){_b="esriGeometryPolyline";}else{if(this.geometryType.toLowerCase()==="polygon"){_b="esriGeometryPolygon";}else{_b=this.geometryType;}}}}_a.geometryType=_b;}return _5.fixJson(_a);}});if(_3("extend-esri")){_2.setObject("layers.QueryDataSource",_8,_4);}return _8;});