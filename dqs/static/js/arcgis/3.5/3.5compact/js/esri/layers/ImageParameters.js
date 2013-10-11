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
define("esri/layers/ImageParameters",["dojo/_base/kernel","dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/lang","esri/layerUtils"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_2(null,{declaredClass:"esri.layers.ImageParameters",constructor:function(){this.layerDefinitions=[];},bbox:null,extent:null,width:null,height:null,dpi:null,format:null,imageSpatialReference:null,layerOption:null,layerIds:null,transparent:null,timeExtent:null,layerTimeOptions:null,toJson:function(_a){if(this.bbox){_1.deprecated(this.declaredClass+" : Property 'bbox' deprecated. Use property 'extent'.");}var bb=this.bbox||this.extent;bb=bb&&_a&&bb._normalize(true);var _b=this.layerOption,_c=bb?(bb.spatialReference.wkid||_4.toJson(bb.spatialReference.toJson())):null,_d=this.imageSpatialReference,_e={dpi:this.dpi,format:this.format,transparent:this.transparent,size:(this.width!==null&&this.height!==null?this.width+","+this.height:null),bbox:(bb?(bb.xmin+","+bb.ymin+","+bb.xmax+","+bb.ymax):null),bboxSR:_c,layers:(_b?_b+":"+this.layerIds.join(","):null),imageSR:(_d?(_d.wkid||_4.toJson(_d.toJson())):_c)};_e.layerDefs=_8._serializeLayerDefinitions(this.layerDefinitions);var _f=this.timeExtent;_e.time=_f?_f.toJson().join(","):null;_e.layerTimeOptions=_8._serializeTimeOptions(this.layerTimeOptions);return _7.filter(_e,function(_10){if(_10!==null){return true;}});}});_3.mixin(_9,{LAYER_OPTION_SHOW:"show",LAYER_OPTION_HIDE:"hide",LAYER_OPTION_INCLUDE:"include",LAYER_OPTION_EXCLUDE:"exclude"});if(_5("extend-esri")){_3.setObject("layers.ImageParameters",_9,_6);}return _9;});