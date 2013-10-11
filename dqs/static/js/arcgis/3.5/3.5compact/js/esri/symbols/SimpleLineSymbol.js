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
define("esri/symbols/SimpleLineSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","dojox/gfx/_base","esri/kernel","esri/lang","esri/symbols/LineSymbol"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9={STYLE_SOLID:"solid",STYLE_DASH:"dash",STYLE_DOT:"dot",STYLE_DASHDOT:"dashdot",STYLE_DASHDOTDOT:"longdashdotdot",STYLE_NULL:"none",STYLE_SHORTDASH:"shortdash",STYLE_SHORTDOT:"shortdot",STYLE_SHORTDASHDOT:"shortdashdot",STYLE_SHORTDASHDOTDOT:"shortdashdotdot",STYLE_LONGDASH:"longdash",STYLE_LONGDASHDOT:"longdashdot"};var _a={color:[0,0,0,1],style:_9.STYLE_SOLID,width:1};var _b=_1(_8,{declaredClass:"esri.symbol.SimpleLineSymbol",type:"simplelinesymbol",_styles:{solid:"esriSLSSolid",dash:"esriSLSDash",dot:"esriSLSDot",dashdot:"esriSLSDashDot",longdashdotdot:"esriSLSDashDotDot",none:"esriSLSNull",insideframe:"esriSLSInsideFrame",shortdash:"esriSLSShortDash",shortdot:"esriSLSShortDot",shortdashdot:"esriSLSShortDashDot",shortdashdotdot:"esriSLSShortDashDotDot",longdash:"esriSLSLongDash",longdashdot:"esriSLSLongDashDot"},constructor:function(_c,_d,_e){if(_c){if(_2.isString(_c)){this.style=_c;if(_d){this.color=_d;}if(_e){this.width=_e;}}else{this.style=_7.valueOf(this._styles,_c.style)||_9.STYLE_SOLID;}}else{_2.mixin(this,_a);this.color=new _3(this.color);this.width=_5.pt2px(this.width);}},setStyle:function(_f){this.style=_f;return this;},getStroke:function(){return (this.style===_9.STYLE_NULL||this.width===0)?null:{color:this.color,style:this.style,width:this.width};},getFill:function(){return null;},getShapeDescriptors:function(){return {defaultShape:{type:"path",path:"M -15,0 L 15,0 E"},fill:null,stroke:this.getStroke()};},toJson:function(){return _7.fixJson(_2.mixin(this.inherited("toJson",arguments),{type:"esriSLS",style:this._styles[this.style]}));}});_2.mixin(_b,_9);_b.defaultProps=_a;if(_4("extend-esri")){_2.setObject("symbol.SimpleLineSymbol",_b,_6);_6.symbol.defaultSimpleLineSymbol=_a;}return _b;});