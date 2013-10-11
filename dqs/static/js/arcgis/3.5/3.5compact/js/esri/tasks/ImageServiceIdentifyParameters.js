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
define("esri/tasks/ImageServiceIdentifyParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/lang","esri/geometry/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.tasks.ImageServiceIdentifyParameters",geometry:null,mosaicRule:null,pixelSizeX:null,pixelSizeY:null,pixelSize:null,returnGeometry:false,timeExtent:null,toJson:function(_9){var g=_9&&_9["geometry"]||this.geometry,_a={geometry:g,returnGeometry:this.returnGeometry,mosaicRule:this.mosaicRule?_3.toJson(this.mosaicRule.toJson()):null};if(g){_a.geometryType=_7.getJsonType(g);}var _b=this.timeExtent;_a.time=_b?_b.toJson().join(","):null;if(_6.isDefined(this.pixelSizeX)&&_6.isDefined(this.pixelSizeY)){_a.pixelSize=_3.toJson({x:parseFloat(this.pixelSizeX),y:parseFloat(this.pixelSizeY)});}else{if(this.pixelSize){_a.pixelSize=this.pixelSize?_3.toJson(this.pixelSize.toJson()):null;}}return _a;}});if(_4("extend-esri")){_2.setObject("tasks.ImageServiceIdentifyParameters",_8,_5);}return _8;});