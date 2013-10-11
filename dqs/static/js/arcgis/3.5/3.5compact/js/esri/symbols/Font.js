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
define("esri/symbols/Font",["dojo/_base/declare","dojo/_base/lang","dojo/sniff","dojox/gfx/_base","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5,_6){var _7={STYLE_NORMAL:"normal",STYLE_ITALIC:"italic",STYLE_OBLIQUE:"oblique",VARIANT_NORMAL:"normal",VARIANT_SMALLCAPS:"small-caps",WEIGHT_NORMAL:"normal",WEIGHT_BOLD:"bold",WEIGHT_BOLDER:"bolder",WEIGHT_LIGHTER:"lighter"};var _8=_1(null,{declaredClass:"esri.symbol.Font",constructor:function(_9,_a,_b,_c,_d){if(_9){if(_2.isObject(_9)){_2.mixin(this,_9);}else{this.size=_9;if(_a!==undefined){this.style=_a;}if(_b!==undefined){this.variant=_b;}if(_c!==undefined){this.weight=_c;}if(_d!==undefined){this.family=_d;}}}else{_2.mixin(this,_4.defaultFont);}if(_3("ie")<9&&this.size&&_2.isString(this.size)&&this.size.indexOf("em")>-1){this.size=_4.pt2px(parseFloat(this.size)*12)+"px";}},setSize:function(_e){this.size=_e;return this;},setStyle:function(_f){this.style=_f;return this;},setVariant:function(_10){this.variant=_10;return this;},setWeight:function(_11){this.weight=_11;return this;},setFamily:function(_12){this.family=_12;return this;},toJson:function(){return _6.fixJson({size:this.size,style:this.style,variant:this.variant,decoration:this.decoration,weight:this.weight,family:this.family});}});_2.mixin(_8,_7);if(_3("extend-esri")){_2.setObject("symbol.Font",_8,_5);}return _8;});