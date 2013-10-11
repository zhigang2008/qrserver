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
define("esri/renderers/TimeRampAger",["dojo/_base/declare","dojo/_base/lang","dojo/_base/Color","dojo/has","esri/kernel","esri/symbols/jsonUtils","esri/renderers/SymbolAger"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(_7,{declaredClass:"esri.renderer.TimeRampAger",constructor:function(_9,_a,_b){this.colorRange=_9;this.sizeRange=_a;this.alphaRange=_b;},getAgedSymbol:function(_c,_d){var _e=_d.getLayer(),_f=_d.attributes;_c=_6.fromJson(_c.toJson());var _10=_e._map.timeExtent;var _11=_10.startTime,_12=_10.endTime;if(!_11||!_12){return _c;}_11=_11.getTime();_12=_12.getTime();var _13=new Date(_f[_e._startTimeField]);_13=_13.getTime();if(_13<_11){_13=_11;}var _14=(_12===_11)?1:(_13-_11)/(_12-_11);var _15=this.sizeRange,_16,_17;if(_15){var _18=_15[0],to=_15[1];_17=Math.abs(to-_18)*_14;this._setSymbolSize(_c,(_18<to)?(_18+_17):(_18-_17));}_15=this.colorRange;if(_15){var _19=_15[0],_1a=_15[1],_1b=Math.round;_16=new _3();var _1c=_19.r,toR=_1a.r;_17=Math.abs(toR-_1c)*_14;_16.r=_1b((_1c<toR)?(_1c+_17):(_1c-_17));var _1d=_19.g,toG=_1a.g;_17=Math.abs(toG-_1d)*_14;_16.g=_1b((_1d<toG)?(_1d+_17):(_1d-_17));var _1e=_19.b,toB=_1a.b;_17=Math.abs(toB-_1e)*_14;_16.b=_1b((_1e<toB)?(_1e+_17):(_1e-_17));var _1f=_19.a,toA=_1a.a;_17=Math.abs(toA-_1f)*_14;_16.a=(_1f<toA)?(_1f+_17):(_1f-_17);_c.setColor(_16);}_16=_c.color;_15=this.alphaRange;if(_15&&_16){var _20=_15[0],_21=_15[1];_17=Math.abs(_21-_20)*_14;_16.a=(_20<_21)?(_20+_17):(_20-_17);}return _c;}});if(_4("extend-esri")){_2.setObject("renderer.TimeRampAger",_8,_5);}return _8;});