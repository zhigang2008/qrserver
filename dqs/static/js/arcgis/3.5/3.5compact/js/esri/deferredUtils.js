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
define("esri/deferredUtils",["dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3){var _4={_dfdCanceller:function(_5){_5.canceled=true;var _6=_5._pendingDfd;if(_5.fired===-1&&_6&&_6.fired===-1){_6.cancel();}_5._pendingDfd=null;},_fixDfd:function(_7){var _8=_7.then;_7.then=function(_9,b,c){if(_9){var _a=_9;_9=function(_b){if(_b&&_b._argsArray){return _a.apply(null,_b);}return _a(_b);};}return _8.call(this,_9,b,c);};return _7;},_resDfd:function(_c,_d,_e){var _f=_d.length;if(_f===1){if(_e){_c.errback(_d[0]);}else{_c.callback(_d[0]);}}else{if(_f>1){_d._argsArray=true;_c.callback(_d);}else{_c.callback();}}}};if(_2("extend-esri")){_1.mixin(_3,_4);}return _4;});