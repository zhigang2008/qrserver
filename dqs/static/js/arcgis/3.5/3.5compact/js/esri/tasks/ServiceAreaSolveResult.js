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
define("esri/tasks/ServiceAreaSolveResult",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/graphic","esri/SpatialReference","esri/tasks/NAMessage"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.tasks.ServiceAreaSolveResult",constructor:function(_a){if(_a.saPolygons){this.serviceAreaPolygons=this._graphicsFromJson(_a.saPolygons);}if(_a.saPolylines){this.serviceAreaPolylines=this._graphicsFromJson(_a.saPolylines);}if(_a.facilities){this.facilities=this._graphicsFromJson(_a.facilities);}if(_a.barriers){this.pointBarriers=this._graphicsFromJson(_a.barriers);}if(_a.polylineBarriers){this.polylineBarriers=this._graphicsFromJson(_a.polylineBarriers);}if(_a.polygonBarriers){this.polygonBarriers=this._graphicsFromJson(_a.polygonBarriers);}if(_a.messages){this.messages=_3.map(_a.messages,function(_b,i){return new _8(_b);});}},serviceAreaPolygons:null,serviceAreaPolylines:null,facilities:null,pointBarriers:null,polylineBarriers:null,polygonBarriers:null,messages:null,_graphicsFromJson:function(_c){var sr=new _7(_c.spatialReference);var _d=_c.features;return _3.map(_d,function(_e,i){var _f=new _6(_e);_f.geometry.setSpatialReference(sr);return _f;});}});if(_4("extend-esri")){_2.setObject("tasks.ServiceAreaSolveResult",_9,_5);}return _9;});