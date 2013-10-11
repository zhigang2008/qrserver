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
define("esri/dijit/BookmarkItem",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.dijit.BookmarkItem",constructor:function(_6){this.name=_6.name;this.extent=_6.extent;},toJson:function(){var _7={};var _8=this.extent.toJson();_7.extent={spatialReference:_8.spatialReference,xmax:_8.xmax,xmin:_8.xmin,ymax:_8.ymax,ymin:_8.ymin};_7.name=this.name;return _7;}});if(_3("extend-esri")){_2.setObject("dijit.BookmarkItem",_5,_4);}return _5;});