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
define("esri/symbols/CartographicLineSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","dojox/gfx/_base","esri/kernel","esri/lang","esri/symbols/SimpleLineSymbol"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9={STYLE_SOLID:"solid",STYLE_DASH:"dash",STYLE_DOT:"dot",STYLE_DASHDOT:"dashdot",STYLE_DASHDOTDOT:"longdashdotdot",STYLE_NULL:"none",STYLE_INSIDE_FRAME:"insideframe",STYLE_SHORTDASH:"shortdash",STYLE_SHORTDOT:"shortdot",STYLE_SHORTDASHDOT:"shortdashdot",STYLE_SHORTDASHDOTDOT:"shortdashdotdot",STYLE_LONGDASH:"longdash",STYLE_LONGDASHDOT:"longdashdot",CAP_BUTT:"butt",CAP_ROUND:"round",CAP_SQUARE:"square",JOIN_MITER:"miter",JOIN_ROUND:"round",JOIN_BEVEL:"bevel"};var _a={color:[0,0,0,1],style:_9.STYLE_SOLID,width:1,cap:_9.CAP_BUTT,join:_9.JOIN_MITER,miterLimit:10};var _b=_1(_8,{declaredClass:"esri.symbol.CartographicLineSymbol",type:"cartographiclinesymbol",_caps:{butt:"esriLCSButt",round:"esriLCSRound",square:"esriLCSSquare"},_joins:{miter:"esriLJSMiter",round:"esriLJSRound",bevel:"esriLJSBevel"},constructor:function(_c,_d,_e,_f,_10,_11){if(_c){if(_2.isString(_c)){this.style=_c;if(_d){this.color=_d;}if(_e!==undefined){this.width=_e;}if(_f){this.cap=_f;}if(_10){this.join=_10;}if(_11!==undefined){this.miterLimit=_11;}}else{this.cap=_7.valueOf(this._caps,_c.cap);this.join=_7.valueOf(this._joins,_c.join);this.width=_5.pt2px(_c.width);this.miterLimit=_5.pt2px(_c.miterLimit);}}else{_2.mixin(this,_a);this.color=new _3(this.color);this.width=_5.pt2px(this.width);this.miterLimit=_5.pt2px(this.miterLimit);}},setCap:function(cap){this.cap=cap;return this;},setJoin:function(_12){this.join=_12;return this;},setMiterLimit:function(_13){this.miterLimit=_13;return this;},getStroke:function(){return _2.mixin(this.inherited("getStroke",arguments),{cap:this.cap,join:(this.join===_9.JOIN_MITER?this.miterLimit:this.join)});},getFill:function(){return null;},getShapeDescriptors:function(){return {defaultShape:{type:"path",path:"M -15,0 L 15,0 E"},fill:null,stroke:this.getStroke()};},toJson:function(){var _14=_5.px2pt(this.miterLimit);_14=isNaN(_14)?undefined:_14;return _7.fixJson(_2.mixin(this.inherited("toJson",arguments),{type:"esriCLS",cap:this._caps[this.cap],join:this._joins[this.join],miterLimit:_14}));}});_2.mixin(_b,_9);_b.defaultProps=_a;if(_4("extend-esri")){_2.setObject("symbol.CartographicLineSymbol",_b,_6);_6.symbol.defaultCartographicLineSymbol=_a;}return _b;});