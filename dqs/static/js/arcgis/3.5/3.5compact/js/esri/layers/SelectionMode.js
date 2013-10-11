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
define("esri/layers/SelectionMode",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/layers/RenderMode"],function(_1,_2,_3,_4,_5){var _6=_1([_5],{declaredClass:"esri.layers._SelectionMode",constructor:function(_7){this.featureLayer=_7;this._featureMap={};},propertyChangeHandler:function(_8){if(this._init&&_8===0){this._applyTimeFilter();}},resume:function(){this.propertyChangeHandler(0);}});if(_3("extend-esri")){_2.setObject("layers._SelectionMode",_6,_4);}return _6;});