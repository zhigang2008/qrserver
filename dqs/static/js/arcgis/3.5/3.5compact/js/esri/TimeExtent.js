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
define("esri/TimeExtent",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.TimeExtent",constructor:function(_6){if(arguments.length>1){this._create(arguments[0],arguments[1]);}else{if(_6){if(_2.isArray(_6)){var _7=_6[0],_8=_6[1];this.startTime=(_7===null||_7==="null")?null:new Date(_7);this.endTime=(_8===null||_8==="null")?null:new Date(_8);}else{if(_6 instanceof Date){this._create(_6,null);}}}}},offset:function(_9,_a){var _b=new _5();var _c=this.startTime,_d=this.endTime;if(_c){_b.startTime=this._getOffsettedDate(_c,_9,_a);}if(_d){_b.endTime=this._getOffsettedDate(_d,_9,_a);}return _b;},intersection:function(_e){return this._intersection(this,_e);},toJson:function(){var _f=[];var _10=this.startTime;_f.push(_10?_10.getTime():"null");var end=this.endTime;_f.push(end?end.getTime():"null");return _f;},_create:function(_11,end){this.startTime=_11?new Date(_11.getTime()):null;this.endTime=end?new Date(end.getTime()):null;},_refData:{"esriTimeUnitsMilliseconds":{getter:"getUTCMilliseconds",setter:"setUTCMilliseconds",multiplier:1},"esriTimeUnitsSeconds":{getter:"getUTCSeconds",setter:"setUTCSeconds",multiplier:1},"esriTimeUnitsMinutes":{getter:"getUTCMinutes",setter:"setUTCMinutes",multiplier:1},"esriTimeUnitsHours":{getter:"getUTCHours",setter:"setUTCHours",multiplier:1},"esriTimeUnitsDays":{getter:"getUTCDate",setter:"setUTCDate",multiplier:1},"esriTimeUnitsWeeks":{getter:"getUTCDate",setter:"setUTCDate",multiplier:7},"esriTimeUnitsMonths":{getter:"getUTCMonth",setter:"setUTCMonth",multiplier:1},"esriTimeUnitsYears":{getter:"getUTCFullYear",setter:"setUTCFullYear",multiplier:1},"esriTimeUnitsDecades":{getter:"getUTCFullYear",setter:"setUTCFullYear",multiplier:10},"esriTimeUnitsCenturies":{getter:"getUTCFullYear",setter:"setUTCFullYear",multiplier:100}},_intersection:function(_12,_13){if(_12&&_13){var _14=_12.startTime,_15=_12.endTime;var _16=_13.startTime,_17=_13.endTime;_14=_14?_14.getTime():-Infinity;_16=_16?_16.getTime():-Infinity;_15=_15?_15.getTime():Infinity;_17=_17?_17.getTime():Infinity;var _18,end;if(_16>=_14&&_16<=_15){_18=_16;}else{if(_14>=_16&&_14<=_17){_18=_14;}}if(_15>=_16&&_15<=_17){end=_15;}else{if(_17>=_14&&_17<=_15){end=_17;}}if(!isNaN(_18)&&!isNaN(end)){var _19=new _5();_19.startTime=(_18===-Infinity)?null:new Date(_18);_19.endTime=(end===Infinity)?null:new Date(end);return _19;}else{return null;}}else{return null;}},_getOffsettedDate:function(_1a,_1b,_1c){var _1d=this._refData;var _1e=new Date(_1a.getTime());if(_1b&&_1c){_1d=_1d[_1c];_1e[_1d.setter](_1e[_1d.getter]()+(_1b*_1d.multiplier));}return _1e;}});if(_3("extend-esri")){_4.TimeExtent=_5;}return _5;});