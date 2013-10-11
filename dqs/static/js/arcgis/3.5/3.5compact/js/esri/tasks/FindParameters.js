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
define("esri/tasks/FindParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/layerUtils"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.FindParameters",searchText:null,contains:true,searchFields:null,outSpatialReference:null,layerIds:null,returnGeometry:false,layerDefinitions:null,dynamicLayerInfos:null,toJson:function(){var _9={searchText:this.searchText,contains:this.contains,returnGeometry:this.returnGeometry,maxAllowableOffset:this.maxAllowableOffset},_a=this.layerIds,_b=this.searchFields,_c=this.outSpatialReference;if(_a){_9.layers=_a.join(",");}if(_b){_9.searchFields=_b.join(",");}if(_c){_9.sr=_c.wkid||_4.toJson(_c.toJson());}_9.layerDefs=_7._serializeLayerDefinitions(this.layerDefinitions);if(this.dynamicLayerInfos&&this.dynamicLayerInfos.length>0){var _d,_e=[];_3.forEach(this.dynamicLayerInfos,function(_f){if(!_f.subLayerIds){var _10=_f.id;if(this.layerIds&&_3.indexOf(this.layerIds,_10)!==-1){var _11={id:_10};_11.source=_f.source&&_f.source.toJson();var _12;if(this.layerDefinitions&&this.layerDefinitions[_10]){_12=this.layerDefinitions[_10];}if(_12){_11.definitionExpression=_12;}_e.push(_11);}}},this);_d=_4.toJson(_e);if(_d==="[]"){_d="[{}]";}_9.dynamicLayers=_d;}return _9;}});if(_5("extend-esri")){_2.setObject("tasks.FindParameters",_8,_6);}return _8;});