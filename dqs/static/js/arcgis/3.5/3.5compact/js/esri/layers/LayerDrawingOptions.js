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
define("esri/layers/LayerDrawingOptions",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/renderers/jsonUtils"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.layers.LayerDrawingOptions",constructor:function(_8){if(_8){_2.mixin(this,_8);if(_8.renderer){this.renderer=_6.fromJson(_8.renderer);}}},toJson:function(){var _9={renderer:this.renderer&&this.renderer.toJson(),transparency:this.transparency,scaleSymbols:this.scaleSymbols,showLabels:this.showLabels};return _5.fixJson(_9);}});if(_3("extend-esri")){_2.setObject("layers.LayerDrawingOptions",_7,_4);}return _7;});