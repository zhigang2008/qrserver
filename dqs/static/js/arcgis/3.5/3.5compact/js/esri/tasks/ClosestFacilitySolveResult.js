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
define("esri/tasks/ClosestFacilitySolveResult",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/graphic","esri/SpatialReference","esri/tasks/DirectionsFeatureSet","esri/tasks/NAMessage"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a=_1(null,{declaredClass:"esri.tasks.ClosestFacilitySolveResult",constructor:function(_b){if(_b.directions){this.directions=[];_3.forEach(_b.directions,function(_c,_d){var _e=[],_f=[];_3.forEach(_c.features,function(f,i){_f[i]=f.compressedGeometry;_e[i]=f.strings;});_c.strings=_e;this.directions[_d]=new _8(_c,_f);},this);}if(_b.routes){this.routes=this._graphicsFromJson(_b.routes);}if(_b.facilities){this.facilities=this._graphicsFromJson(_b.facilities);}if(_b.incidents){this.incidents=this._graphicsFromJson(_b.incidents);}if(_b.barriers){this.pointBarriers=this._graphicsFromJson(_b.barriers);}if(_b.polylineBarriers){this.polylineBarriers=this._graphicsFromJson(_b.polylineBarriers);}if(_b.polygonBarriers){this.polygonBarriers=this._graphicsFromJson(_b.polygonBarriers);}if(_b.messages){this.messages=_3.map(_b.messages,function(_10,i){return new _9(_10);});}},routes:null,facilities:null,incidents:null,pointBarriers:null,polylineBarriers:null,polygonBarriers:null,directions:null,messages:null,_graphicsFromJson:function(_11){var sr=new _7(_11.spatialReference);var _12=_11.features;return _3.map(_12,function(_13,i){var _14=new _6(_13);_14.geometry.setSpatialReference(sr);return _14;});}});if(_4("extend-esri")){_2.setObject("tasks.ClosestFacilitySolveResult",_a,_5);}return _a;});