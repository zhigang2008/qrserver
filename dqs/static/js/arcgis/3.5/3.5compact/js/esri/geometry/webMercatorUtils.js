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
define("esri/geometry/webMercatorUtils",["dojo/_base/array","dojo/_base/lang","dojo/has","esri/kernel","esri/SpatialReference","esri/geometry/Point"],function(_1,_2,_3,_4,_5,_6){function _7(_8,_9,sr,_a){if(_8.type==="point"){var pt=_9(_8.x,_8.y,_a);return new _8.constructor(pt[0],pt[1],new _5(sr));}else{if(_8.type==="extent"){var _b=_9(_8.xmin,_8.ymin,_a),_c=_9(_8.xmax,_8.ymax,_a);return new _8.constructor(_b[0],_b[1],_c[0],_c[1],new _5(sr));}else{if(_8.type==="polyline"||_8.type==="polygon"){var _d=(_8.type==="polyline"),_e=_d?_8.paths:_8.rings,_f=[],_10;_1.forEach(_e,function(_11){_f.push(_10=[]);_1.forEach(_11,function(iPt){_10.push(_9(iPt[0],iPt[1],_a));});});if(_d){return new _8.constructor({paths:_f,spatialReference:sr});}else{return new _8.constructor({rings:_f,spatialReference:sr});}}else{if(_8.type==="multipoint"){var _12=[];_1.forEach(_8.points,function(iPt){_12.push(_9(iPt[0],iPt[1],_a));});return new _8.constructor({points:_12,spatialReference:sr});}}}}};var _13={lngLatToXY:_6.lngLatToXY,xyToLngLat:_6.xyToLngLat,geographicToWebMercator:function(_14){return _7(_14,_6.lngLatToXY,{wkid:102100});},webMercatorToGeographic:function(_15,_16){return _7(_15,_6.xyToLngLat,{wkid:4326},_16);}};if(_3("extend-esri")){_2.mixin(_2.getObject("geometry",true,_4),_13);}return _13;});