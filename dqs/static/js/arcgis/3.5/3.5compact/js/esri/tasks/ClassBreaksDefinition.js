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
define("esri/tasks/ClassBreaksDefinition",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/tasks/ClassificationDefinition"],function(_1,_2,_3,_4,_5){var _6=_1(_5,{declaredClass:"esri.tasks.ClassBreaksDefinition",type:"classBreaksDef",classificationField:null,classificationMethod:null,breakCount:null,standardDeviationInterval:null,normalizationType:null,normalizationField:null,toJson:function(){var _7=this.inherited(arguments);var _8;switch(this.classificationMethod.toLowerCase()){case "natural-breaks":_8="esriClassifyNaturalBreaks";break;case "equal-interval":_8="esriClassifyEqualInterval";break;case "quantile":_8="esriClassifyQuantile";break;case "standard-deviation":_8="esriClassifyStandardDeviation";break;case "geometrical-interval":_8="esriClassifyGeometricalInterval";break;default:_8=this.classificationMethod;}_2.mixin(_7,{type:this.type,classificationField:this.classificationField,classificationMethod:_8,breakCount:this.breakCount});if(this.normalizationType){var _9;switch(this.normalizationType.toLowerCase()){case "field":_9="esriNormalizeByField";break;case "log":_9="esriNormalizeByLog";break;case "percent-of-total":_9="esriNormalizeByPercentOfTotal";break;default:_9=this.normalizationType;}_2.mixin(_7,{normalizationType:_9});}if(this.normalizationField){_2.mixin(_7,{normalizationField:this.normalizationField});}if(this.standardDeviationInterval){_2.mixin(_7,{standardDeviationInterval:this.standardDeviationInterval});}return _7;}});if(_3("extend-esri")){_2.setObject("tasks.ClassBreaksDefinition",_6,_4);}return _6;});