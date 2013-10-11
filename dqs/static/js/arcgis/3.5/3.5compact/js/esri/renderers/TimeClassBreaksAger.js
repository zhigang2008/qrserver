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
define("esri/renderers/TimeClassBreaksAger",["dojo/_base/declare","dojo/_base/array","dojo/_base/lang","dojo/has","dojo/date","esri/kernel","esri/lang","esri/symbols/jsonUtils","esri/renderers/SymbolAger"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a={UNIT_DAYS:"day",UNIT_HOURS:"hour",UNIT_MILLISECONDS:"millisecond",UNIT_MINUTES:"minute",UNIT_MONTHS:"month",UNIT_SECONDS:"second",UNIT_WEEKS:"week",UNIT_YEARS:"year"};var _b=_1(_9,{declaredClass:"esri.renderer.TimeClassBreaksAger",constructor:function(_c,_d){this.infos=_c;this.timeUnits=_d||"day";_c.sort(function(a,b){if(a.minAge<b.minAge){return -1;}if(a.minAge>b.minAge){return 1;}return 0;});},getAgedSymbol:function(_e,_f){var _10=_f.getLayer(),_11=_f.attributes,_12=_7.isDefined;_e=_8.fromJson(_e.toJson());var _13=_10._map.timeExtent;var _14=_13.endTime;if(!_14){return _e;}var _15=new Date(_11[_10._startTimeField]);var _16=_5.difference(_15,_14,this.timeUnits);_2.some(this.infos,function(_17){if(_16>=_17.minAge&&_16<=_17.maxAge){var _18=_17.color,_19=_17.size,_1a=_17.alpha;if(_18){_e.setColor(_18);}if(_12(_19)){this._setSymbolSize(_e,_19);}if(_12(_1a)&&_e.color){_e.color.a=_1a;}return true;}},this);return _e;}});_3.mixin(_b,_a);if(_4("extend-esri")){_3.setObject("renderer.TimeClassBreaksAger",_b,_6);}return _b;});