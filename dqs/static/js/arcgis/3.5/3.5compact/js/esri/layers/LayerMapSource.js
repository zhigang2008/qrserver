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
define("esri/layers/LayerMapSource",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/LayerSource"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.layers.LayerMapSource",type:"mapLayer",toJson:function(){var _8={type:"mapLayer",mapLayerId:this.mapLayerId,gdbVersion:this.gdbVersion};return _5.fixJson(_8);}});if(_3("extend-esri")){_2.setObject("layers.LayerMapSource",_7,_4);}return _7;});