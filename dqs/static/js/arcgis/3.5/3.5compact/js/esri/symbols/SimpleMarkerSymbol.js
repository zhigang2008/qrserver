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
define("esri/symbols/SimpleMarkerSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","dojox/gfx/_base","esri/kernel","esri/lang","esri/symbols/MarkerSymbol","esri/symbols/SimpleLineSymbol"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a={STYLE_CIRCLE:"circle",STYLE_SQUARE:"square",STYLE_CROSS:"cross",STYLE_X:"x",STYLE_DIAMOND:"diamond",STYLE_PATH:"path",STYLE_TARGET:"target"};var _b={style:_a.STYLE_CIRCLE,color:[255,255,255,0.25],size:12,angle:0,xoffset:0,yoffset:0};var _c=_1(_8,{declaredClass:"esri.symbol.SimpleMarkerSymbol",type:"simplemarkersymbol",_styles:{circle:"esriSMSCircle",square:"esriSMSSquare",cross:"esriSMSCross",x:"esriSMSX",diamond:"esriSMSDiamond",path:"esriSMSPath"},constructor:function(_d,_e,_f,_10){if(_d){if(_2.isString(_d)){this.style=_d;if(_e){this.size=_e;}if(_f){this.outline=_f;}if(_10){this.color=_10;}}else{this.style=_7.valueOf(this._styles,this.style);if(_d.outline){this.outline=new _9(_d.outline);}}}else{_2.mixin(this,_b);this.size=_5.pt2px(this.size);this.outline=new _9(this.outline);this.color=new _3(this.color);}if(!this.style){this.style=_a.STYLE_CIRCLE;}},setStyle:function(_11){this.style=_11;return this;},setPath:function(_12){this.path=_12;this.setStyle(_a.STYLE_PATH);return this;},setOutline:function(_13){this.outline=_13;return this;},getStroke:function(){return this.outline&&this.outline.getStroke();},getFill:function(){return this.color;},_setDim:function(_14,_15,_16){this._targetWidth=_14;this._targetHeight=_15;this._spikeSize=_16;},getShapeDescriptors:function(){var _17,_18,_19;var _1a=this.style,_1b=this.size||_5.pt2px(_b.size),cx=0,cy=0,_1c=_1b/2,_1d=cx-_1c,_1e=cx+_1c,top=cy-_1c,_1f=cy+_1c;switch(_1a){case _a.STYLE_CIRCLE:_17={type:"circle",cx:cx,cy:cy,r:_1c};_18=this.getFill();_19=this.getStroke();if(_19){_19.style=_19.style||"Solid";}break;case _a.STYLE_CROSS:_17={type:"path",path:"M "+_1d+",0 L "+_1e+",0 M 0,"+top+" L 0,"+_1f+" E"};_18=null;_19=this.getStroke();break;case _a.STYLE_DIAMOND:_17={type:"path",path:"M "+_1d+",0 L 0,"+top+" L "+_1e+",0 L 0,"+_1f+" L "+_1d+",0 E"};_18=this.getFill();_19=this.getStroke();break;case _a.STYLE_SQUARE:_17={type:"path",path:"M "+_1d+","+_1f+" L "+_1d+","+top+" L "+_1e+","+top+" L "+_1e+","+_1f+" L "+_1d+","+_1f+" E"};_18=this.getFill();_19=this.getStroke();break;case _a.STYLE_X:_17={type:"path",path:"M "+_1d+","+_1f+" L "+_1e+","+top+" M "+_1d+","+top+" L "+_1e+","+_1f+" E"};_18=null;_19=this.getStroke();break;case _a.STYLE_PATH:_17={type:"path",path:this.path||""};_18=this.getFill();_19=this.getStroke();break;}return {defaultShape:_17,fill:_18,stroke:_19};},toJson:function(){var _20=_2.mixin(this.inherited("toJson",arguments),{type:"esriSMS",style:this._styles[this.style]}),_21=this.outline;if(_21){_20.outline=_21.toJson();}_20.path=this.path;return _7.fixJson(_20);}});_2.mixin(_c,_a);_c.defaultProps=_b;if(_4("extend-esri")){_2.setObject("symbol.SimpleMarkerSymbol",_c,_6);_6.symbol.defaultSimpleMarkerSymbol=_b;}return _c;});