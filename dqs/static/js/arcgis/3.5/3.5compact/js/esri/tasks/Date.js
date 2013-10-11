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
define("esri/tasks/Date",["dojo/_base/declare","dojo/_base/lang","dojo/date/locale","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.tasks.Date",constructor:function(_7){if(_7){if(_7.format){this.format=_7.format;}this.date=_3.parse(_7.date,{selector:"date",datePattern:this.format});}},date:new Date(),format:"EEE MMM dd HH:mm:ss zzz yyyy",toJson:function(){return {date:_3.format(this.date,{selector:"date",datePattern:this.format}),format:this.format};}});if(_4("extend-esri")){_2.setObject("tasks.Date",_6,_5);}return _6;});