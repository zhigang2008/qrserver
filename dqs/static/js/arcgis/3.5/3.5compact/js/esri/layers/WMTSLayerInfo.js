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
define("esri/layers/WMTSLayerInfo",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5=_1(null,{declaredClass:"esri.layers.WMTSLayerInfo",identifier:null,tileMatrixSet:null,format:null,style:null,tileInfo:null,title:null,fullExtent:null,initialExtent:null,description:null,constructor:function(_6){if(_6){this.title=_6.title;this.tileMatrixSet=_6.tileMatrixSet;this.format=_6.format;this.style=_6.style;this.tileInfo=_6.tileInfo;this.fullExtent=_6.fullExtent;this.initialExtent=_6.initialExtent;this.identifier=_6.identifier;this.description=_6.description;}}});if(_3("extend-esri")){_2.setObject("layers.WMTSLayerInfo",_5,_4);}return _5;});