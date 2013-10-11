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
define("esri/layers/DynamicLayerInfo",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/LayerInfo","esri/layers/LayerMapSource","esri/layers/LayerDataSource"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(_6,{declaredClass:"esri.layers.DynamicLayerInfo",defaultVisibility:true,parentLayerId:-1,maxScale:0,minScale:0,constructor:function(_a){if(_a){var _b;if(!_a.source){_b=new _7();_b.mapLayerId=this.id;}else{if(_a.source.type==="mapLayer"){_b=new _7(_a.source);}else{_b=new _8(_a.source);}}this.source=_b;}},toJson:function(){var _c=this.inherited(arguments);_c.source=this.source&&this.source.toJson();return _5.fixJson(_c);}});if(_3("extend-esri")){_2.setObject("layers.DynamicLayerInfo",_9,_4);}return _9;});