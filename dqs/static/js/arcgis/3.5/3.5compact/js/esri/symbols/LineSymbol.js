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
define("esri/symbols/LineSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/has","dojox/gfx/_base","esri/kernel","esri/symbols/Symbol"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.symbol.LineSymbol",constructor:function(_8){if(_2.isObject(_8)){this.width=_4.pt2px(this.width);}else{this.width=12;}},setWidth:function(_9){this.width=_9;return this;},toJson:function(){var _a=_4.px2pt(this.width);_a=isNaN(_a)?undefined:_a;return _2.mixin(this.inherited("toJson",arguments),{width:_a});}});if(_3("extend-esri")){_2.setObject("symbol.LineSymbol",_7,_5);}return _7;});