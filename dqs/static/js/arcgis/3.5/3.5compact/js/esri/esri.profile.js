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
var profile=(function(){var _1=/^esri\/tests\//,_2=/\.js$/i,_3=/^esri\/arcgisonline\//i,_4=/^esri\/mobile\//i,_5=/^esri\/discovery\//i,_6=/^esri\/dijit\/analysis\//i,_7=function(_8,_9){var _a={"esri/package.json":1,"esri/esri.profile.js":1,"esri/esri.js":1};return (_9 in _a);},_b={"esri/arcgisonline":1,"esri/base":1,"esri/gallery":1,"esri/mobile":1,"esri/arcgismanager":1},_c={"arcgis/Portal":1,"dijit/Directions":1,"dijit/Geocoder":1,"dijit/Gauge":1,"dijit/Scalebar":1,"main":1,"MapNavigationManager":1,"MouseEvents":1,"PointerEvents":1,"TouchEvents":1};return {resourceTags:{test:function(_d,_e){return _1.test(_e)||(_e.search(/\.17$/)!==-1);},copyOnly:function(_f,mid){return _7(_f,mid);},amd:function(_10,mid){return _2.test(_10)&&(/^esri\/arcgisonline\/sharing\/dijit\/FeatureLayerQueryResult/i.test(mid)||!((mid in _b)||_3.test(mid)||_4.test(mid)||_5.test(mid)));}}};}());