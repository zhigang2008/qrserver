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
define("esri/symbols/jsonUtils",["dojo/_base/lang","dojo/has","esri/kernel","esri/symbols/SimpleMarkerSymbol","esri/symbols/PictureMarkerSymbol","esri/symbols/SimpleLineSymbol","esri/symbols/CartographicLineSymbol","esri/symbols/SimpleFillSymbol","esri/symbols/PictureFillSymbol","esri/symbols/TextSymbol"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b={fromJson:function(_c){var _d=_c.type,_e=null;switch(_d.substring(0,"esriXX".length)){case "esriSM":_e=new _4(_c);break;case "esriPM":_e=new _5(_c);break;case "esriTS":_e=new _a(_c);break;case "esriSL":if(_c.cap!==undefined){_e=new _7(_c);}else{_e=new _6(_c);}break;case "esriSF":_e=new _8(_c);break;case "esriPF":_e=new _9(_c);break;}return _e;},getShapeDescriptors:function(_f){return (_f&&_f.getShapeDescriptors)?_f.getShapeDescriptors():{defaultShape:null,fill:null,stroke:null};}};if(_2("extend-esri")){_1.mixin(_1.getObject("symbol",true,_3),_b);}return _b;});