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
define("esri/layers/PurgeOptions",["dojo/_base/declare","dojo/_base/lang","dojo/Stateful","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5){var _6=_1([_3],{declaredClass:"esri.layers.PurgeOptions",constructor:function(_7,_8){this.parent=_7;var p;for(p in _8){this[p]=_8[p];}},_displayCountSetter:function(_9){this.displayCount=_9;this.parent.refresh();}});if(_4("extend-esri")){_2.setObject("layers.PurgeOptions",_6,_5);}return _6;});