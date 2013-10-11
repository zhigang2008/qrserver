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
define("esri/graphicsUtils",["dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/geometry/Extent"],function(_1,_2,_3,_4,_5){var _6={graphicsExtent:function(_7){var g=_7[0].geometry,_8=g.getExtent(),_9,i,il=_7.length;if(_8===null){_8=new _5(g.x,g.y,g.x,g.y,g.spatialReference);}for(i=1;i<il;i++){_9=(g=_7[i].geometry).getExtent();if(_9===null){_9=new _5(g.x,g.y,g.x,g.y,g.spatialReference);}_8=_8.union(_9);}if(_8.getWidth()<=0&&_8.getHeight()<=0){return null;}return _8;},getGeometries:function(_a){return _2.map(_a,function(_b){return _b.geometry;});},_encodeGraphics:function(_c,_d){var _e=[],_f,enc,_10;_2.forEach(_c,function(g,i){_f=g.toJson();enc={};if(_f.geometry){_10=_d&&_d[i];enc.geometry=_10&&_10.toJson()||_f.geometry;}if(_f.attributes){enc.attributes=_f.attributes;}_e[i]=enc;});return _e;}};if(_3("extend-esri")){_1.mixin(_4,_6);}return _6;});