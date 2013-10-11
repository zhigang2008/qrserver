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
define("esri/tasks/ClassificationDefinition",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.tasks.ClassificationDefinition",type:null,baseSymbol:null,colorRamp:null,toJson:function(){var _6={};if(this.baseSymbol){_2.mixin(_6,{baseSymbol:this.baseSymbol.toJson()});}if(this.colorRamp){_2.mixin(_6,{colorRamp:this.colorRamp.toJson()});}return _6;}});if(_3("extend-esri")){_2.setObject("tasks.ClassificationDefinition",_5,_4);}return _5;});