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
define("esri/symbols/MarkerSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/has","dojox/gfx/_base","esri/kernel","esri/symbols/Symbol"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.symbol.MarkerSymbol",angle:0,xoffset:0,yoffset:0,size:12,constructor:function(_8){if(_8&&_2.isObject(_8)){this.size=_4.pt2px(this.size);this.xoffset=_4.pt2px(this.xoffset);this.yoffset=_4.pt2px(this.yoffset);}},setAngle:function(_9){this.angle=_9;return this;},setSize:function(_a){this.size=_a;return this;},setOffset:function(x,y){this.xoffset=x;this.yoffset=y;return this;},toJson:function(){var _b=_4.px2pt(this.size);_b=isNaN(_b)?undefined:_b;var _c=_4.px2pt(this.xoffset);_c=isNaN(_c)?undefined:_c;var _d=_4.px2pt(this.yoffset);_d=isNaN(_d)?undefined:_d;return _2.mixin(this.inherited("toJson",arguments),{size:_b,angle:this.angle,xoffset:_c,yoffset:_d});}});if(_3("extend-esri")){_2.setObject("symbol.MarkerSymbol",_7,_5);}return _7;});