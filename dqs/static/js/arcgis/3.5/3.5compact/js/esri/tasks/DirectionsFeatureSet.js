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
define("esri/tasks/DirectionsFeatureSet",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/geometry/Extent","esri/geometry/Polyline","esri/tasks/FeatureSet"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(_8,{declaredClass:"esri.tasks.DirectionsFeatureSet",constructor:function(_a,_b){this.routeId=_a.routeId;this.routeName=_a.routeName;_2.mixin(this,_a.summary);this.extent=new _6(this.envelope);var _c=this._fromCompressedGeometry,_d=this.features,sr=this.extent.spatialReference,_e=[];_3.forEach(_b,function(cg,i){_d[i].setGeometry(_e[i]=_c(cg,sr));});this.strings=_a.strings;this.mergedGeometry=this._mergePolylinesToSinglePath(_e,sr);this.geometryType="esriGeometryPolyline";delete this.envelope;},_fromCompressedGeometry:function(_f,sr){var _10=0,_11=0,_12=[],x,y,_13=_f.replace(/(\+)|(\-)/g," $&").split(" "),j,jl=_13.length,_14=parseInt(_13[1],32);for(j=2;j<jl;j+=2){_10=(x=(parseInt(_13[j],32)+_10));_11=(y=(parseInt(_13[j+1],32)+_11));_12.push([x/_14,y/_14]);}var po=new _7({paths:[_12]});po.setSpatialReference(sr);return po;},_mergePolylinesToSinglePath:function(_15,sr){var _16=[];_3.forEach(_15,function(_17){_3.forEach(_17.paths,function(_18){_16=_16.concat(_18);});});var _19=[],_1a=[0,0];_3.forEach(_16,function(_1b){if(_1b[0]!==_1a[0]||_1b[1]!==_1a[1]){_19.push(_1b);_1a=_1b;}});return new _7({paths:[_19]}).setSpatialReference(sr);}});if(_4("extend-esri")){_2.setObject("tasks.DirectionsFeatureSet",_9,_5);}return _9;});