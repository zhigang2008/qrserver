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
define("esri/layers/RasterFunction",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.layers.RasterFunction",functionName:null,"arguments":null,"functionArguments":null,variableName:null,constructor:function(_7){if(!_2.isObject(_7)){return;}_2.mixin(this,_7);if(_7.rasterFunction){this.functionName=_7.rasterFunction;}if(_7.rasterFunctionArguments){this["functionArguments"]=_7.rasterFunctionArguments;}else{if(_7["arguments"]){this["functionArguments"]=_7["arguments"];}}},toJson:function(){var _8={rasterFunction:this.functionName,rasterFunctionArguments:this.functionArguments||this["arguments"],variableName:this.variableName};return _5.filter(_8,function(_9){if(_9!==null){return true;}});}});if(_3("extend-esri")){_2.setObject("layers.RasterFunction",_6,_4);}return _6;});