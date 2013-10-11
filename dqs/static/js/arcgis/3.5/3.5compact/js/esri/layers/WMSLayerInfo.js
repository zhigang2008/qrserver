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
define("esri/layers/WMSLayerInfo",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.layers.WMSLayerInfo",name:null,title:null,description:null,extent:null,legendURL:null,subLayers:[],allExtents:[],spatialReferences:[],constructor:function(_7){if(_7){this.name=_7.name;this.title=_7.title;this.description=_7.description;this.extent=_7.extent;this.legendURL=_7.legendURL;this.subLayers=_7.subLayers?_7.subLayers:[];this.allExtents=_7.allExtents?_7.allExtents:[];this.spatialReferences=_7.spatialReferences?_7.spatialReferences:[];}},clone:function(){var _8={name:this.name,title:this.title,description:this.description,legendURL:this.legendURL},_9;if(this.extent){_8.extent=this.extent.getExtent();}_8.subLayers=[];_3.forEach(this.subLayers,function(_a){_8.subLayers.push(_a.clone());});_8.allExtents=[];for(_9 in this.allExtents){_9=parseInt(_9,10);if(!isNaN(_9)){_8.allExtents[_9]=this.allExtents[_9].getExtent();}}_8.spatialReferences=[];_3.forEach(this.spatialReferences,function(_b){_8.spatialReferences.push(_b);});return _8;}});if(_4("extend-esri")){_2.setObject("layers.WMSLayerInfo",_6,_5);}return _6;});