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
define("esri/symbols/FillSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/symbols/Symbol","esri/symbols/SimpleLineSymbol"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_5,{declaredClass:"esri.symbol.FillSymbol",constructor:function(_8){if(_8&&_2.isObject(_8)&&_8.outline){this.outline=new _6(_8.outline);}},setOutline:function(_9){this.outline=_9;return this;},toJson:function(){var _a=this.inherited("toJson",arguments);if(this.outline){_a.outline=this.outline.toJson();}return _a;}});if(_3("extend-esri")){_2.setObject("symbol.FillSymbol",_7,_4);}return _7;});