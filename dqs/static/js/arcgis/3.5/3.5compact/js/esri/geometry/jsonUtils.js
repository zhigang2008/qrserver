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
define("esri/geometry/jsonUtils",["dojo/_base/lang","dojo/has","esri/kernel","esri/geometry/Point","esri/geometry/Polyline","esri/geometry/Polygon","esri/geometry/Multipoint","esri/geometry/Extent"],function(_1,_2,_3,_4,_5,_6,_7,_8){function _9(_a){if(_a.x!==undefined&&_a.y!==undefined){return new _4(_a);}else{if(_a.paths!==undefined){return new _5(_a);}else{if(_a.rings!==undefined){return new _6(_a);}else{if(_a.points!==undefined){return new _7(_a);}else{if(_a.xmin!==undefined&&_a.ymin!==undefined&&_a.xmax!==undefined&&_a.ymax!==undefined){return new _8(_a);}}}}}};function _b(_c){if(_c instanceof _4){return "esriGeometryPoint";}else{if(_c instanceof _5){return "esriGeometryPolyline";}else{if(_c instanceof _6){return "esriGeometryPolygon";}else{if(_c instanceof _8){return "esriGeometryEnvelope";}else{if(_c instanceof _7){return "esriGeometryMultipoint";}}}}}return null;};function _d(_e){if(_e==="esriGeometryPoint"){return _4;}else{if(_e==="esriGeometryPolyline"){return _5;}else{if(_e==="esriGeometryPolygon"){return _6;}else{if(_e==="esriGeometryEnvelope"){return _8;}else{if(_e==="esriGeometryMultipoint"){return _7;}}}}}return null;};var _f={fromJson:_9,getJsonType:_b,getGeometryType:_d};if(_2("extend-esri")){_1.mixin(_1.getObject("geometry",true,_3),_f);}return _f;});