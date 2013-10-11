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
define("esri/kernel",["dojo/_base/kernel","dojo/_base/config","dojo/has"],function(_1,_2,_3){var _4=window.location,_5=_4.pathname,_6={version:3.5,_appBaseUrl:_4.protocol+"//"+_4.host+_5.substring(0,_5.lastIndexOf(_5.split("/")[_5.split("/").length-1]))};if(!_2.noGlobals){window.esri=_6;}if(!_1.isAsync){_3.add("extend-esri",1);}var _7=_6.dijit=(_6.dijit||{});_7._arcgisUrl=_4.protocol+"//www.arcgis.com/sharing/rest";return _6;});