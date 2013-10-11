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
define("esri/layers/LayerTimeOptions",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.layers.LayerTimeOptions",constructor:function(_7){if(_7){_2.mixin(this,_7);}},toJson:function(){var _8={timeDataCumulative:this.timeDataCumulative,timeOffset:this.timeOffset,timeOffsetUnits:this.timeOffsetUnits,useTime:this.useTime};return _5.fixJson(_8);}});if(_3("extend-esri")){_2.setObject("layers.LayerTimeOptions",_6,_4);}return _6;});