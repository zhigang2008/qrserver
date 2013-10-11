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
define("esri/tasks/TrimExtendParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.tasks.TrimExtendParameters",polylines:null,trimExtendTo:null,extendHow:null,toJson:function(){var _8=_3.map(this.polylines,function(_9){return _9.toJson();});var _a={};_a.polylines=_4.toJson(_8);_a.trimExtendTo=_4.toJson(this.trimExtendTo.toJson());_a.sr=_4.toJson(this.polylines[0].spatialReference.toJson());_a.extendHow=this.extendHow||0;return _a;}});_2.mixin(_7,{DEFAULT_CURVE_EXTENSION:0,RELOCATE_ENDS:1,KEEP_END_ATTRIBUTES:2,NO_END_ATTRIBUTES:4,NO_EXTEND_AT_FROM:8,NO_EXTEND_AT_TO:16});if(_5("extend-esri")){_2.setObject("tasks.TrimExtendParameters",_7,_6);}return _7;});