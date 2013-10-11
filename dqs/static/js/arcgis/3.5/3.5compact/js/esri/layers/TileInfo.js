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
define("esri/layers/TileInfo",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/lang","esri/SpatialReference","esri/geometry/Point","esri/layers/LOD"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a=_1(null,{declaredClass:"esri.layers.TileInfo",constructor:function(_b){_2.mixin(this,_b);this.width=this.cols;this.height=this.rows;var sr=this.spatialReference,_c=this.origin;if(sr){sr=(this.spatialReference=new _7(sr.toJson?sr.toJson():sr));}if(_c){this.origin=new _8(_c.toJson?_c.toJson():_c);if(!_c.spatialReference&&sr){this.origin.setSpatialReference(new _7(sr.toJson()));}}this.lods=_3.map(this.lods,function(_d){return new _9(_d);});},toJson:function(){return _6.fixJson({rows:this.rows,cols:this.cols,dpi:this.dpi,format:this.format,compressionQuality:this.compressionQuality,origin:this.origin&&this.origin.toJson(),spatialReference:this.spatialReference&&this.spatialReference.toJson(),lods:this.lods&&_3.map(this.lods,function(_e){return _e.toJson();})});}});if(_4("extend-esri")){_2.setObject("layers.TileInfo",_a,_5);}return _a;});