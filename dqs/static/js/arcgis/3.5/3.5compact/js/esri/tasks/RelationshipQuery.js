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
define("esri/tasks/RelationshipQuery",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.tasks.RelationshipQuery",definitionExpression:"",relationshipId:null,returnGeometry:false,objectIds:null,outSpatialReference:null,outFields:null,toJson:function(){var _6={definitionExpression:this.definitionExpression,relationshipId:this.relationshipId,returnGeometry:this.returnGeometry,maxAllowableOffset:this.maxAllowableOffset,geometryPrecision:this.geometryPrecision},_7=this.objectIds,_8=this.outFields,_9=this.outSpatialReference;if(_7){_6.objectIds=_7.join(",");}if(_8){_6.outFields=_8.join(",");}if(_9){_6.outSR=_9.toJson();}_6._ts=this._ts;return _6;}});if(_3("extend-esri")){_2.setObject("tasks.RelationshipQuery",_5,_4);}return _5;});