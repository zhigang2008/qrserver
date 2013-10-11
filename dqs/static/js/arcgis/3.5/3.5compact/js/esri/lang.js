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
define("esri/lang",["dojo/_base/array","dojo/_base/json","dojo/_base/kernel","dojo/_base/lang","dojo/date","dojo/has","dojo/number","dojo/date/locale","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){function _a(_b,_c,cb){return [_4.isString(_b)?_b.split(""):_b,_c||_3.global,_4.isString(cb)?new Function("item","index","array",cb):cb];};function _d(_e){return (_e!==undefined)&&(_e!==null);};var _f="${*}",_10=["NumberFormat","DateString","DateFormat"];function _11(_12){return _d(_12)?_12:"";};function _13(key,_14,_15){var _16=_15.match(/([^\(]+)(\([^\)]+\))?/i),_17=_4.trim(_16[1]),_18=_14[key],_19,_1a=_2.fromJson((_16[2]?_4.trim(_16[2]):"()").replace(/^\(/,"({").replace(/\)$/,"})")),_1b=_1a.utcOffset;if(_1.indexOf(_10,_17)===-1){var ref=_4.getObject(_17);if(_4.isFunction(ref)){_18=ref(_18,key,_14);}}else{if(typeof _18==="number"||(typeof _18==="string"&&_18&&!isNaN(Number(_18)))){_18=Number(_18);switch(_17){case "NumberFormat":return _7.format(_18,_1a);break;case "DateString":_19=new Date(_18);if(_1a.local||_1a.systemLocale){if(_1a.systemLocale){return _19.toLocaleDateString()+(_1a.hideTime?"":(" "+_19.toLocaleTimeString()));}else{return _19.toDateString()+(_1a.hideTime?"":(" "+_19.toTimeString()));}}else{_19=_19.toUTCString();if(_1a.hideTime){_19=_19.replace(/\s+\d\d\:\d\d\:\d\d\s+(utc|gmt)/i,"");}return _19;}break;case "DateFormat":_19=new Date(_18);if(_d(_1b)){_19=_5.add(_19,"minute",(_19.getTimezoneOffset()-_1b));}return _8.format(_19,_1a);break;}}}return _11(_18);};function _1c(obj,_1d){var _1e;if(_1d){for(_1e in obj){if(obj.hasOwnProperty(_1e)){if(obj[_1e]===undefined||obj[_1e]===null){delete obj[_1e];}else{if(obj[_1e] instanceof Object){_1c(obj[_1e],true);}}}}}else{for(_1e in obj){if(obj.hasOwnProperty(_1e)){if(obj[_1e]===undefined){delete obj[_1e];}}}}return obj;};var _1f={valueOf:function(_20,_21){var i;for(i in _20){if(_20[i]==_21){return i;}}return null;},substitute:function(_22,_23,_24){var _25,_26,_27;if(_d(_24)){if(_4.isObject(_24)){_25=_24.first;_26=_24.dateFormat;_27=_24.numberFormat;}else{_25=_24;}}if(!_23||_23===_f){var s=[],val,i;for(i in _22){val=_22[i];if(_26&&_1.indexOf(_26.properties||"",i)!==-1){val=_13(i,_22,_26.formatter||"DateString");}else{if(_27&&_1.indexOf(_27.properties||"",i)!==-1){val=_13(i,_22,_27.formatter||"NumberFormat");}}s.push(i+" = "+_11(val)+"<br/>");if(_25){break;}}return s.join("");}else{return _4.replace(_23,_4.hitch({obj:_22},function(_28,key){var _29=key.split(":");if(_29.length>1){key=_29[0];_29.shift();return _13(key,this.obj,_29.join(":"));}else{if(_26&&_1.indexOf(_26.properties||"",key)!==-1){return _13(key,this.obj,_26.formatter||"DateString");}if(_27&&_1.indexOf(_27.properties||"",key)!==-1){return _13(key,this.obj,_27.formatter||"NumberFormat");}}return _11(this.obj[key]);}),/\$\{([^\}]+)\}/g);}},filter:function(arr,_2a,_2b){var _2c=_a(arr,_2b,_2a),_2d={},i;arr=_2c[0];for(i in arr){if(_2c[2].call(_2c[i],arr[i],i,arr)){_2d[i]=arr[i];}}return _2d;},isDefined:_d,fixJson:_1c};if(_6("extend-esri")){_4.mixin(_9,_1f);_9._isDefined=_d;_9._getParts=_a;_9._sanitize=_1c;}return _1f;});