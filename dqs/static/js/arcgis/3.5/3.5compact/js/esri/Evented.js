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
define("esri/Evented",["dojo/_base/declare","dojo/_base/lang","dojo/aspect","dojo/Evented","dojo/on","dojo/has","esri/kernel"],function(_1,_2,_3,_4,on,_5,_6){var _7=_1([_4],{declaredClass:"esri.Evented",registerConnectEvents:function(_8,_9){if(!_9){return;}var _a=this.constructor.prototype,_b={},_c=this.declaredClass==_8,_d,_e,_f,k;if(!_a.hasOwnProperty("_onMap")||!this._onMap._final){for(_d in this){if(/^on/.test(_d)){_e=this._hyphenLower(_d).toLowerCase();if(!_a._onMap||!_a._onMap[_e]||_9[_e]){_b[_e]={"method":_d};if(_9[_e]){_b[_e].argKeys=_9[_e];}}}}for(k in _9){if(!_b[k]){_f=this._onCamelCase(k);_b[k]={"method":_f,"argKeys":_9[k]};}}_b._final=_c;if(!this._onMap){_a._onMap=_b;}else{var _10=_2.mixin({},_a._onMap);_a._onMap=_2.mixin(_10,_b);}}},on:function(_11,_12){var _13=this._onMap,_14=(typeof _11=="string")&&_11.toLowerCase(),_15=this._onCamelCase(_14),_16=_13&&_13[_14],_17=(_16&&_16.method)||(this[_15]&&_2.isFunction(this[_15])&&_15),_18;if(_17){if(_16&&_2.isArray(_16.argKeys)){_18=this._onArr2Obj(_12,_13[_14].argKeys);return _3.after(this,_17,_18,true);}else{return _3.after(this,_17,_12,true);}}return this.inherited(arguments);},emit:function(_19,_1a){var ret,_1b,_1c,_1d,_1e=_19.toLowerCase(),_1f=this._onCamelCase(_19),_20=this._onMap;_1c=(_20&&_20[_1e]&&_20[_1e].method)||(_2.isFunction(this[_1f])&&_1f);_1d=_1c&&this[_1c];if(_1c&&_20&&_20[_1e]){this._onObj2Arr(function(){_1b=Array.prototype.slice.call(arguments);},_20[_1e].argKeys)(_1a);}_1a=_1a||{};if(!_1a.target){_1a.target=this;}if(_1d){ret=_1d.apply(this,_1b||[_1a]);}this.inherited(arguments);return ret;},_onObj2Arr:function(_21,_22){if(!_22){return _21;}else{var _23=this;return function(evt){var i,_24=[],_25=_22.length;for(i=0;i<_25;i++){_24[i]=evt[_22[i]];}_21.apply(_23,_24);};}},_onArr2Obj:function(_26,_27){if(!_27){return _26;}else{var _28=this;return function(){var i,evt={},_29=arguments.length;for(i=0;i<_29;i++){evt[_27[i]]=arguments[i];}if(!evt.target){evt.target=_28;}_26(evt);};}},_hyphenLower:function(_2a){return _2a.replace(/^on/,"").replace(/[A-Z](?=[a-z])/g,function(m,off){return (off?"-":"")+m.toLowerCase();});},_onCamelCase:function(_2b){return "on"+_2b.substr(0,1).toUpperCase()+_2b.substr(1).replace(/\-([a-z])/g,function(m,s){return s.toUpperCase();});}});if(_5("extend-esri")){_6.Evented=_7;}return _7;});