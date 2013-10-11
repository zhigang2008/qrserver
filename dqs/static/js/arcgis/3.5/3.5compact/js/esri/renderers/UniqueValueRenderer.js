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
define("esri/renderers/UniqueValueRenderer",["dojo/_base/declare","dojo/_base/array","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/symbols/jsonUtils","esri/renderers/Renderer"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(_8,{declaredClass:"esri.renderer.UniqueValueRenderer",constructor:function(_a,_b,_c,_d,_e){this.values=[];this._symbols={};this.infos=[];if(_a&&!_a.declaredClass){var _f=_a;_a=_f.defaultSymbol;if(_a){this.defaultSymbol=_7.fromJson(_a);}this.attributeField=_f.field1;this.attributeField2=_f.field2;this.attributeField3=_f.field3;this.fieldDelimiter=_f.fieldDelimiter;this.defaultLabel=_f.defaultLabel;_2.forEach(_f.uniqueValueInfos,this._addValueInfo,this);}else{this.defaultSymbol=_a;this.attributeField=_b;this.attributeField2=_c;this.attributeField3=_d;this.fieldDelimiter=_e;}this._multiple=!!this.attributeField2;},addValue:function(_10,_11){var _12=_3.isObject(_10)?_10:{value:_10,symbol:_11};this._addValueInfo(_12);},removeValue:function(_13){var i=_2.indexOf(this.values,_13);if(i===-1){return;}this.values.splice(i,1);delete this._symbols[_13];this.infos.splice(i,1);},getSymbol:function(_14){var _15=this.attributeField,_16=_14.attributes,_17,_18,_19;if(this._multiple){_17=this.attributeField2;_18=this.attributeField3;_19=[];if(_15){_19.push(_16[_15]);}if(_17){_19.push(_16[_17]);}if(_18){_19.push(_16[_18]);}return this._symbols[_19.join(this.fieldDelimiter||"")]||this.defaultSymbol;}else{_15=_3.isFunction(_15)?_15(_14):_16[_15];return this._symbols[_15]||this.defaultSymbol;}},_addValueInfo:function(_1a){var _1b=_1a.value;this.values.push(_1b);this.infos.push(_1a);var _1c=_1a.symbol;if(_1c){if(!_1c.declaredClass){_1a.symbol=_7.fromJson(_1c);}}this._symbols[_1b]=_1a.symbol;},toJson:function(){var _1d=_6.fixJson;return _1d({type:"uniqueValue",field1:this.attributeField,field2:this.attributeField2,field3:this.attributeField3,fieldDelimiter:this.fieldDelimiter,defaultSymbol:this.defaultSymbol&&this.defaultSymbol.toJson(),defaultLabel:this.defaultLabel,uniqueValueInfos:_2.map(this.infos||[],function(_1e){_1e=_3.mixin({},_1e);_1e.symbol=_1e.symbol&&_1e.symbol.toJson();_1e.value=_1e.value+"";return _1d(_1e);})});}});if(_4("extend-esri")){_3.setObject("renderer.UniqueValueRenderer",_9,_5);}return _9;});