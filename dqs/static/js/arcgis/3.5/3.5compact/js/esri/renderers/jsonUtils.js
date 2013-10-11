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
define("esri/renderers/jsonUtils",["dojo/_base/lang","dojo/has","esri/kernel","esri/renderers/SimpleRenderer","esri/renderers/UniqueValueRenderer","esri/renderers/ClassBreaksRenderer"],function(_1,_2,_3,_4,_5,_6){var _7={fromJson:function(_8){var _9=_8.type||"",_a;switch(_9){case "simple":_a=new _4(_8);break;case "uniqueValue":_a=new _5(_8);break;case "classBreaks":_a=new _6(_8);break;}return _a;}};if(_2("extend-esri")){_1.mixin(_1.getObject("renderer",true,_3),_7);}return _7;});