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
define("esri/tasks/BufferParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/geometry/Polygon","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.tasks.BufferParameters",geometries:null,outSpatialReference:null,bufferSpatialReference:null,distances:null,unit:null,unionResults:false,geodesic:false,toJson:function(){var _a={unit:this.unit,unionResults:this.unionResults,geodesic:this.geodesic},dt=this.distances,_b=this.outSpatialReference,_c=this.bufferSpatialReference;var _d=_3.map(this.geometries,function(_e){_e=(_e.type==="extent")?this._extentToPolygon(_e):_e;return _e.toJson();},this);var _f=this.geometries;if(_f&&_f.length>0){var _10=_f[0].type==="extent"?"esriGeometryPolygon":_8.getJsonType(_f[0]);_a.geometries=_4.toJson({geometryType:_10,geometries:_d});_a.inSR=_f[0].spatialReference.wkid?_f[0].spatialReference.wkid:_4.toJson(_f[0].spatialReference.toJson());}if(dt){_a.distances=dt.join(",");}if(_b){_a.outSR=_b.wkid?_b.wkid:_4.toJson(_b.toJson());}if(_c){_a.bufferSR=_c.wkid?_c.wkid:_4.toJson(_c.toJson());}return _a;},_extentToPolygon:function(_11){var _12=_11.xmin,_13=_11.ymin,_14=_11.xmax,_15=_11.ymax;return new _7({"rings":[[[_12,_13],[_12,_15],[_14,_15],[_14,_13],[_12,_13]]],"spatialReference":_11.spatialReference.toJson()});}});if(_5("extend-esri")){_2.setObject("tasks.BufferParameters",_9,_6);}return _9;});