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
define("esri/layers/ImageServiceParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/lang"],function(_1,_2,_3,_4,_5,_6){var _7=_1(null,{declaredClass:"esri.layers.ImageServiceParameters",extent:null,width:null,height:null,imageSpatialReference:null,format:null,interpolation:null,compressionQuality:null,bandIds:null,timeExtent:null,mosaicRule:null,renderingRule:null,noData:null,toJson:function(_8){var _9=this.bbox||this.extent;_9=_9&&_8&&_9._normalize(true);var _a=_9?(_9.spatialReference.wkid||_3.toJson(_9.spatialReference.toJson())):null,_b=this.imageSpatialReference,_c={bbox:_9?(_9.xmin+","+_9.ymin+","+_9.xmax+","+_9.ymax):null,bboxSR:_a,size:(this.width!==null&&this.height!==null?this.width+","+this.height:null),imageSR:(_b?(_b.wkid||_3.toJson(_b.toJson())):_a),format:this.format,interpolation:this.interpolation,compressionQuality:this.compressionQuality,bandIds:this.bandIds?this.bandIds.join(","):null,mosaicRule:this.mosaicRule?_3.toJson(this.mosaicRule.toJson()):null,renderingRule:this.renderingRule?_3.toJson(this.renderingRule.toJson()):null,noData:this.noData};var _d=this.timeExtent;_c.time=_d?_d.toJson().join(","):null;return _6.filter(_c,function(_e){if(_e!==null){return true;}});}});_2.mixin(_7,{INTERPOLATION_BILINEAR:"RSP_BilinearInterpolation",INTERPOLATION_CUBICCONVOLUTION:"RSP_CubicConvolution",INTERPOLATION_MAJORITY:"RSP_Majority",INTERPOLATION_NEARESTNEIGHBOR:"RSP_NearestNeighbor"});if(_4("extend-esri")){_2.setObject("layers.ImageServiceParameters",_7,_5);}return _7;});