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
define("esri/tasks/FeatureSet",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/lang","esri/graphic","esri/SpatialReference","esri/graphicsUtils","esri/geometry/jsonUtils","esri/symbols/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b){var _c=_1(null,{declaredClass:"esri.tasks.FeatureSet",constructor:function(_d){if(_d){_2.mixin(this,_d);var _e=this.features,sr=_d.spatialReference,_f=_a.getGeometryType(_d.geometryType);sr=(this.spatialReference=new _8(sr));this.geometryType=_d.geometryType;if(_d.fields){this.fields=_d.fields;}_3.forEach(_e,function(_10,i){var _11=_10.geometry&&_10.geometry.spatialReference;_e[i]=new _7((_f&&_10.geometry)?new _f(_10.geometry):null,_10.symbol&&_b.fromJson(_10.symbol),_10.attributes);if(_e[i].geometry&&!_11){_e[i].geometry.setSpatialReference(sr);}});}else{this.features=[];}},displayFieldName:null,geometryType:null,spatialReference:null,fieldAliases:null,toJson:function(_12){var _13={};if(this.displayFieldName){_13.displayFieldName=this.displayFieldName;}if(this.fields){_13.fields=this.fields;}if(this.spatialReference){_13.spatialReference=this.spatialReference.toJson();}else{if(this.features[0]&&this.features[0].geometry){_13.spatialReference=this.features[0].geometry.spatialReference.toJson();}}if(this.features[0]){if(this.features[0].geometry){_13.geometryType=_a.getJsonType(this.features[0].geometry);}_13.features=_9._encodeGraphics(this.features,_12);}_13.exceededTransferLimit=this.exceededTransferLimit;return _6.fixJson(_13);}});if(_4("extend-esri")){_2.setObject("tasks.FeatureSet",_c,_5);}return _c;});