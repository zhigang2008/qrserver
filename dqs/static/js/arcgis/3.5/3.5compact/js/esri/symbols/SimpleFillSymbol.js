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
define("esri/symbols/SimpleFillSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","dojox/gfx/_base","esri/kernel","esri/lang","esri/symbols/FillSymbol","esri/symbols/SimpleLineSymbol","require"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b={STYLE_SOLID:"solid",STYLE_NULL:"none",STYLE_HORIZONTAL:"horizontal",STYLE_VERTICAL:"vertical",STYLE_FORWARD_DIAGONAL:"forwarddiagonal",STYLE_BACKWARD_DIAGONAL:"backwarddiagonal",STYLE_CROSS:"cross",STYLE_DIAGONAL_CROSS:"diagonalcross",STYLE_FORWARDDIAGONAL:"forwarddiagonal",STYLE_BACKWARDDIAGONAL:"backwarddiagonal",STYLE_DIAGONALCROSS:"diagonalcross"};var _c={style:_b.STYLE_SOLID,color:[0,0,0,0.25]};var _d=_1(_8,{declaredClass:"esri.symbol.SimpleFillSymbol",type:"simplefillsymbol",_styles:{solid:"esriSFSSolid",none:"esriSFSNull",horizontal:"esriSFSHorizontal",vertical:"esriSFSVertical",forwarddiagonal:"esriSFSForwardDiagonal",backwarddiagonal:"esriSFSBackwardDiagonal",cross:"esriSFSCross",diagonalcross:"esriSFSDiagonalCross"},constructor:function(_e,_f,_10){if(_e){if(_2.isString(_e)){this.style=_e;if(_f!==undefined){this.outline=_f;}if(_10!==undefined){this.color=_10;}}else{this.style=_7.valueOf(this._styles,_e.style);}}else{_2.mixin(this,_c);this.outline=new _9(this.outline);this.color=new _3(this.color);}var _11=this.style;if(_11!=="solid"&&_11!=="none"){this._src=_a.toUrl("esri")+"/images/symbol/sfs/"+_11+".png";}},setStyle:function(_12){this.style=_12;return this;},getStroke:function(){return this.outline&&this.outline.getStroke();},getFill:function(){var _13=this.style;if(_13===_b.STYLE_NULL){return null;}else{if(_13===_b.STYLE_SOLID){return this.color;}else{return _2.mixin(_2.mixin({},_5.defaultPattern),{src:this._src,width:10,height:10});}}},getShapeDescriptors:function(){return {defaultShape:{type:"path",path:"M -10,-10 L 10,0 L 10,10 L -10,10 L -10,-10 E"},fill:this.getFill(),stroke:this.getStroke()};},toJson:function(){return _7.fixJson(_2.mixin(this.inherited("toJson",arguments),{type:"esriSFS",style:this._styles[this.style]}));}});_2.mixin(_d,_b);_d.defaultProps=_c;if(_4("extend-esri")){_2.setObject("symbol.SimpleFillSymbol",_d,_6);_6.symbol.defaultSimpleFillSymbol=_c;}return _d;});