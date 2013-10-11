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
define("esri/symbols/Symbol",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5,_6){function _7(_8){return _8&&new _3([_8[0],_8[1],_8[2],_8[3]/255]);};function _9(_a){return _a&&[_a.r,_a.g,_a.b,Math.round(_a.a*255)];};var _b=_1(null,{declaredClass:"esri.symbol.Symbol",color:new _3([0,0,0,1]),type:null,_stroke:null,_fill:null,constructor:function(_c){if(_c&&_2.isObject(_c)){_2.mixin(this,_c);if(this.color&&_6.isDefined(this.color[0])){this.color=_7(this.color);}var _d=this.type;if(_d&&_d.indexOf("esri")===0){this.type={"esriSMS":"simplemarkersymbol","esriPMS":"picturemarkersymbol","esriSLS":"simplelinesymbol","esriCLS":"cartographiclinesymbol","esriSFS":"simplefillsymbol","esriPFS":"picturefillsymbol","esriTS":"textsymbol"}[_d];}}},setColor:function(_e){this.color=_e;return this;},toJson:function(){return {color:_9(this.color)};}});_b.toDojoColor=_7;_b.toJsonColor=_9;if(_4("extend-esri")){_2.setObject("symbol.Symbol",_b,_5);_5.symbol.toDojoColor=_7;_5.symbol.toJsonColor=_9;}return _b;});