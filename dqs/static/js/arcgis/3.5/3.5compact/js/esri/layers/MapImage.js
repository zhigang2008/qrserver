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
define("esri/layers/MapImage",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/domUtils","esri/geometry/Extent"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.layers.MapImage",constructor:function(_8){_2.mixin(this,_8);this.extent=new _6(this.extent);},visible:true,getLayer:function(){return this._layer;},getNode:function(){return this._node;},show:function(){if(!this.visible){this.visible=true;var _9=this._node,_a=this._layer,_b;if(_9){_b=_a&&_a._div;if(_b){if(!_a.suspended){_a._setPos(_9,_b._left,_b._top);}(_a._active||_b).appendChild(_9);}_5.show(_9);}}},hide:function(){if(this.visible){this.visible=false;var _c=this._node;if(_c){_5.hide(_c);if(_c.parentNode){_c.parentNode.removeChild(_c);}}}}});if(_3("extend-esri")){_2.setObject("layers.MapImage",_7,_4);}return _7;});