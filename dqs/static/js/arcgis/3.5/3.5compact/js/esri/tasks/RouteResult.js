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
define("esri/tasks/RouteResult",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/graphic","esri/tasks/DirectionsFeatureSet"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.RouteResult",constructor:function(_9){var sr=_9.spatialReference,_a=_9.route;if(_9.directions){var _b=[],_c=[];_3.forEach(_9.directions.features,function(f,i){_c[i]=f.compressedGeometry;_b[i]=f.strings;});_9.directions.strings=_b;this.directions=new _7(_9.directions,_c);}this.routeName=_9.routeName;if(_a){if(_a.geometry){_a.geometry.spatialReference=sr;}this.route=new _6(_a);}if(_9.stops){var ss=(this.stops=[]);_3.forEach(_9.stops,function(_d,i){if(_d.geometry){_d.geometry.spatialReference=sr;}ss[_d.attributes.Sequence-1]=new _6(_d);});}},routeName:null,directions:null,route:null,stops:null});if(_4("extend-esri")){_2.setObject("tasks.RouteResult",_8,_5);}return _8;});