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
define("esri/tasks/IdentifyParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/layerUtils","esri/geometry/jsonUtils","esri/geometry/scaleUtils"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a=_1(null,{declaredClass:"esri.tasks.IdentifyParameters",constructor:function(){this.layerOption=_a.LAYER_OPTION_TOP;},geometry:null,spatialReference:null,layerIds:null,tolerance:null,returnGeometry:false,mapExtent:null,width:400,height:400,dpi:96,layerDefinitions:null,timeExtent:null,layerTimeOptions:null,dynamicLayerInfos:null,toJson:function(_b){var g=_b&&_b["geometry"]||this.geometry,_c=this.mapExtent,sr=this.spatialReference,_d=this.layerIds,_e={geometry:g,tolerance:this.tolerance,returnGeometry:this.returnGeometry,mapExtent:_c,imageDisplay:this.width+","+this.height+","+this.dpi,maxAllowableOffset:this.maxAllowableOffset};if(g){_e.geometryType=_8.getJsonType(g);}if(sr!==null){_e.sr=sr.wkid||_4.toJson(sr.toJson());}else{if(g){_e.sr=g.spatialReference.wkid||_4.toJson(g.spatialReference.toJson());}else{if(_c){_e.sr=_c.spatialReference.wkid||_4.toJson(_c.spatialReference.toJson());}}}_e.layers=this.layerOption;if(_d){_e.layers+=":"+_d.join(",");}_e.layerDefs=_7._serializeLayerDefinitions(this.layerDefinitions);var _f=this.timeExtent;_e.time=_f?_f.toJson().join(","):null;_e.layerTimeOptions=_7._serializeTimeOptions(this.layerTimeOptions);if(this.dynamicLayerInfos&&this.dynamicLayerInfos.length>0){var _10,_11={extent:_c,width:this.width,spatialReference:_c.spatialReference},_12=_9.getScale(_11),_13=_7._getLayersForScale(_12,this.dynamicLayerInfos),_14=[];_3.forEach(this.dynamicLayerInfos,function(_15){if(!_15.subLayerIds){var _16=_15.id;if((!this.layerIds||(this.layerIds&&_3.indexOf(this.layerIds,_16)!==-1))&&_3.indexOf(_13,_16)!==-1){var _17={id:_16};_17.source=_15.source&&_15.source.toJson();var _18;if(this.layerDefinitions&&this.layerDefinitions[_16]){_18=this.layerDefinitions[_16];}if(_18){_17.definitionExpression=_18;}var _19;if(this.layerTimeOptions&&this.layerTimeOptions[_16]){_19=this.layerTimeOptions[_16];}if(_19){_17.layerTimeOptions=_19.toJson();}_14.push(_17);}}},this);_10=_4.toJson(_14);if(_10==="[]"){_10="[{}]";}_e.dynamicLayers=_10;}return _e;}});_2.mixin(_a,{LAYER_OPTION_TOP:"top",LAYER_OPTION_VISIBLE:"visible",LAYER_OPTION_ALL:"all"});if(_5("extend-esri")){_2.setObject("tasks.IdentifyParameters",_a,_6);}return _a;});