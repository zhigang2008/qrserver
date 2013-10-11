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
define("esri/tasks/UniqueValueDefinition",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/tasks/ClassificationDefinition"],function(_1,_2,_3,_4,_5){var _6=_1(_5,{declaredClass:"esri.tasks.UniqueValueDefinition",type:"uniqueValueDef",attributeField:null,attributeField2:null,attributeField3:null,fieldDelimiter:null,toJson:function(){var _7=this.inherited(arguments);this.uniqueValueFields=[];if(this.attributeField){this.uniqueValueFields.push(this.attributeField);}if(this.attributeField2){this.uniqueValueFields.push(this.attributeField2);}if(this.attributeField3){this.uniqueValueFields.push(this.attributeField3);}_2.mixin(_7,{type:this.type,uniqueValueFields:this.uniqueValueFields});if(this.fieldDelimiter){_2.mixin(_7,{fieldDelimiter:this.fieldDelimiter});}return _7;}});if(_3("extend-esri")){_2.setObject("tasks.UniqueValueDefinition",_6,_4);}return _6;});