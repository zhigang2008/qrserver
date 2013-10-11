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
define("esri/layers/TimeInfo",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/TimeExtent","esri/layers/TimeReference","esri/layers/LayerTimeOptions"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.layers.TimeInfo",constructor:function(_9){if(_9!==null){_2.mixin(this,_9);if(_9.exportOptions){this.exportOptions=new _7(_9.exportOptions);}this.timeExtent=new _5(_9.timeExtent);this.timeReference=new _6(_9.timeReference);}}});_2.mixin(_8,{UNIT_CENTURIES:"esriTimeUnitsCenturies",UNIT_DAYS:"esriTimeUnitsDays",UNIT_DECADES:"esriTimeUnitsDecades",UNIT_HOURS:"esriTimeUnitsHours",UNIT_MILLISECONDS:"esriTimeUnitsMilliseconds",UNIT_MINUTES:"esriTimeUnitsMinutes",UNIT_MONTHS:"esriTimeUnitsMonths",UNIT_SECONDS:"esriTimeUnitsSeconds",UNIT_UNKNOWN:"esriTimeUnitsUnknown",UNIT_WEEKS:"esriTimeUnitsWeeks",UNIT_YEARS:"esriTimeUnitsYears"});if(_3("extend-esri")){_2.setObject("layers.TimeInfo",_8,_4);}return _8;});