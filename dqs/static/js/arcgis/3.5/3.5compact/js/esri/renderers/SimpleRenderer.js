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
define("esri/renderers/SimpleRenderer",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/symbols/jsonUtils","esri/renderers/Renderer"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(_7,{declaredClass:"esri.renderer.SimpleRenderer",constructor:function(_9){if(_9&&!_9.declaredClass){var _a=_9;_9=_a.symbol;if(_9){this.symbol=_6.fromJson(_9);}this.label=_a.label;this.description=_a.description;}else{this.symbol=_9;}},getSymbol:function(_b){return this.symbol;},toJson:function(){return _5.fixJson({type:"simple",label:this.label,description:this.description,symbol:this.symbol&&this.symbol.toJson()});}});if(_3("extend-esri")){_2.setObject("renderer.SimpleRenderer",_8,_4);}return _8;});