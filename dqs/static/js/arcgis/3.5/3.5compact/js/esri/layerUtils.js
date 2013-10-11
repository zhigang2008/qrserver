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
define("esri/layerUtils",["dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5){var _6={_serializeLayerDefinitions:function(_7){var _8=[],_9=false,re=/[:;]/;if(_7){_2.forEach(_7,function(_a,i){if(_a){_8.push([i,_a]);if(!_9&&re.test(_a)){_9=true;}}});if(_8.length>0){var _b;if(_9){_b={};_2.forEach(_8,function(_c){_b[_c[0]]=_c[1];});_b=_3.toJson(_b);}else{_b=[];_2.forEach(_8,function(_d){_b.push(_d[0]+":"+_d[1]);});_b=_b.join(";");}return _b;}}return null;},_serializeTimeOptions:function(_e,_f){if(!_e){return;}var _10=[];_2.forEach(_e,function(_11,i){if(_11){var _12=_11.toJson();if(_f&&_2.indexOf(_f,i)!==-1){_12.useTime=false;}_10.push("\""+i+"\":"+_3.toJson(_12));}});if(_10.length){return "{"+_10.join(",")+"}";}},_getDefaultVisibleLayers:function(_13){var _14=[],i;if(!_13){return _14;}for(i=0;i<_13.length;i++){if(_13[i].parentLayerId>=0&&_2.indexOf(_14,_13[i].parentLayerId)===-1&&_2.some(_13,function(_15){return _15.id===_13[i].parentLayerId;})){continue;}if(_13[i].defaultVisibility){_14.push(_13[i].id);}}return _14;},_getLayersForScale:function(_16,_17){var _18=[];if(_16>0&&_17){var i;for(i=0;i<_17.length;i++){if(_17[i].parentLayerId>=0&&_2.indexOf(_18,_17[i].parentLayerId)===-1&&_2.some(_17,function(_19){return _19.id===_17[i].parentLayerId;})){continue;}if(_17[i].id>=0){var _1a=true,_1b=_17[i].maxScale,_1c=_17[i].minScale;if(_1b>0||_1c>0){if(_1b>0&&_1c>0){_1a=_1b<=_16&&_16<=_1c;}else{if(_1b>0){_1a=_1b<=_16;}else{if(_1c>0){_1a=_16<=_1c;}}}}if(_1a){_18.push(_17[i].id);}}}}return _18;}};if(_4("extend-esri")){_1.mixin(_5,_6);}return _6;});