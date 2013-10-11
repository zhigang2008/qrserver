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
define("esri/layers/RangeDomain",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/Domain"],function(_1,_2,_3,_4,_5,_6){var _7=_1([_6],{declaredClass:"esri.layers.RangeDomain",constructor:function(_8){if(_8&&_2.isObject(_8)){this.minValue=_8.range[0];this.maxValue=_8.range[1];}},toJson:function(){var _9=this.inherited(arguments);_9.range=[this.minValue,this.maxValue];return _5.fixJson(_9);}});if(_3("extend-esri")){_2.setObject("layers.RangeDomain",_7,_4);}return _7;});