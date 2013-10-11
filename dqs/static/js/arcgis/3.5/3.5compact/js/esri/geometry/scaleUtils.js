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
define("esri/geometry/scaleUtils",["dojo/_base/lang","dojo/has","esri/kernel","esri/config","esri/lang","esri/WKIDUnitConversion"],function(_1,_2,_3,_4,_5,_6){var _7=39.37,_8=20015077/180,_9=_4.defaults,_a=_6;function _b(_c,_d,_e){return (_c&&_d)?((_c.getWidth()/_d)*(_e||_8)*_7*_9.screenDPI):0;};function _f(_10,_11,_12,_13,_14){var _15;if(_14){_15=_12;}else{_15=_a.values[_a[_12]];}return _10.expand(((_13*_11)/((_15||_8)*_7*_9.screenDPI))/_10.getWidth());};var _16={getScale:function(map,_17,_18){var _19,_1a,wkt;if(arguments.length>1&&(_5.isDefined(_17)&&!_17.declaredClass)){_19=map;_1a=_17;_17=null;}else{_19=map.extent;_1a=map.width;var sr=map.spatialReference;if(sr){_18=sr.wkid;wkt=sr.wkt;}}var _1b;if(_18){_1b=_a.values[_a[_18]];}else{if(wkt&&(wkt.search(/^PROJCS/i)!==-1)){var _1c=/UNIT\[([^\]]+)\]\]$/i.exec(wkt);if(_1c&&_1c[1]){_1b=parseFloat(_1c[1].split(",")[1]);}}}return _b(_17||_19,_1a,_1b);},getExtentForScale:function(map,_1d,_1e){var _1f,wkt,sr=map.spatialReference;if(sr){_1f=sr.wkid;wkt=sr.wkt;}var _20;if(_1f){_20=_a.values[_a[_1f]];}else{if(wkt&&(wkt.search(/^PROJCS/i)!==-1)){var _21=/UNIT\[([^\]]+)\]\]$/i.exec(wkt);if(_21&&_21[1]){_20=parseFloat(_21[1].split(",")[1]);}}}return _f(_1e||map.extent,map.width,_20,_1d,true);}};if(_2("extend-esri")){_1.mixin(_1.getObject("geometry",true,_3),_16);_3.geometry._getScale=_b;_3.geometry._getExtentForScale=_f;}return _16;});