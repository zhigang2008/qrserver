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
define("esri/symbols/TextSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","dojox/gfx/_base","esri/kernel","esri/lang","esri/symbols/Symbol","esri/symbols/Font"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a={ALIGN_START:"start",ALIGN_MIDDLE:"middle",ALIGN_END:"end",DECORATION_NONE:"none",DECORATION_UNDERLINE:"underline",DECORATION_OVERLINE:"overline",DECORATION_LINETHROUGH:"line-through"};var _b={color:[0,0,0,1],font:_5.defaultFont,angle:0,xoffset:0,yoffset:0};_2.mixin(_b,_5.defaultText,{type:"textsymbol",align:"middle"});var _c=_1(_8,{declaredClass:"esri.symbol.TextSymbol",angle:0,xoffset:0,yoffset:0,constructor:function(_d,_e,_f){_2.mixin(this,_b);this.font=new _9(this.font);this.color=new _3(this.color);if(_d){if(_2.isObject(_d)){_2.mixin(this,_d);if(this.color&&_7.isDefined(this.color[0])){this.color=_8.toDojoColor(this.color);}this.type="textsymbol";this.font=new _9(this.font);this.xoffset=_5.pt2px(this.xoffset);this.yoffset=_5.pt2px(this.yoffset);}else{this.text=_d;if(_e){this.font=_e;}if(_f){this.color=_f;}}}},setFont:function(_10){this.font=_10;return this;},setAngle:function(_11){this.angle=_11;return this;},setOffset:function(x,y){this.xoffset=x;this.yoffset=y;return this;},setAlign:function(_12){this.align=_12;return this;},setDecoration:function(_13){this.decoration=_13;return this;},setRotated:function(_14){this.rotated=_14;return this;},setKerning:function(_15){this.kerning=_15;return this;},setText:function(_16){this.text=_16;return this;},getStroke:function(){return null;},getFill:function(){return this.color;},toJson:function(){var _17=_5.px2pt(this.xoffset);_17=isNaN(_17)?undefined:_17;var _18=_5.px2pt(this.yoffset);_18=isNaN(_18)?undefined:_18;return _7.fixJson(_2.mixin(this.inherited("toJson",arguments),{type:"esriTS",backgroundColor:this.backgroundColor,borderLineColor:this.borderLineColor,verticalAlignment:this.verticalAlignment,horizontalAlignment:this.horizontalAlignment,rightToLeft:this.rightToLeft,width:this.width,angle:this.angle,xoffset:_17,yoffset:_18,text:this.text,align:this.align,decoration:this.decoration,rotated:this.rotated,kerning:this.kerning,font:this.font.toJson()}));}});_2.mixin(_c,_a);_c.defaultProps=_b;if(_4("extend-esri")){_2.setObject("symbol.TextSymbol",_c,_6);_6.symbol.defaultTextSymbol=_b;}return _c;});