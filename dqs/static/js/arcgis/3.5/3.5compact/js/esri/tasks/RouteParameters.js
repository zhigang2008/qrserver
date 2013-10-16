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
define("esri/tasks/RouteParameters",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/lang","esri/graphicsUtils","esri/tasks/NATypes"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.tasks.RouteParameters",accumulateAttributes:null,attributeParameterValues:null,barriers:null,directionsLanguage:null,directionsLengthUnits:null,directionsOutputType:null,directionsStyleName:null,directionsTimeAttribute:null,doNotLocateOnRestrictedElements:true,findBestSequence:null,ignoreInvalidLocations:null,impedanceAttribute:null,outputLines:null,outputGeometryPrecision:null,outputGeometryPrecisionUnits:null,outSpatialReference:null,polygonBarriers:null,polylineBarriers:null,preserveFirstStop:null,preserveLastStop:null,restrictionAttributes:null,restrictUTurns:null,returnBarriers:false,returnDirections:false,returnPolygonBarriers:false,returnPolylineBarriers:false,returnRoutes:true,returnStops:false,startTime:null,stops:null,useHierarchy:null,useTimeWindows:null,toJson:function(_a){var _b={returnDirections:this.returnDirections,returnRoutes:this.returnRoutes,returnStops:this.returnStops,returnBarriers:this.returnBarriers,returnPolygonBarriers:this.returnPolygonBarriers,returnPolylineBarriers:this.returnPolylineBarriers,attributeParameterValues:this.attributeParameterValues&&_3.toJson(this.attributeParameterValues),outSR:this.outSpatialReference?(this.outSpatialReference.wkid||_3.toJson(this.outSpatialReference.toJson())):null,outputLines:this.outputLines,findBestSequence:this.findBestSequence,preserveFirstStop:this.preserveFirstStop,preserveLastStop:this.preserveLastStop,useTimeWindows:this.useTimeWindows,startTime:this.startTime?this.startTime.getTime():null,accumulateAttributeNames:this.accumulateAttributes?this.accumulateAttributes.join(","):null,ignoreInvalidLocations:this.ignoreInvalidLocations,impedanceAttributeName:this.impedanceAttribute,restrictionAttributeNames:this.restrictionAttributes?this.restrictionAttributes.join(","):null,restrictUTurns:this.restrictUTurns,useHierarchy:this.useHierarchy,directionsLanguage:this.directionsLanguage,outputGeometryPrecision:this.outputGeometryPrecision,outputGeometryPrecisionUnits:this.outputGeometryPrecisionUnits,directionsLengthUnits:_8.LengthUnit[this.directionsLengthUnits],directionsTimeAttributeName:this.directionsTimeAttribute,directionsStyleName:this.directionsStyleName},_c=this.stops;if(_c.declaredClass==="esri.tasks.FeatureSet"&&_c.features.length>0){_b.stops=_3.toJson({type:"features",features:_7._encodeGraphics(_c.features,_a&&_a["stops.features"]),doNotLocateOnRestrictedElements:this.doNotLocateOnRestrictedElements});}else{if(_c.declaredClass==="esri.tasks.DataLayer"){_b.stops=_c;}else{if(_c.declaredClass==="esri.tasks.DataFile"){_b.stops=_3.toJson({type:"features",url:_c.url,doNotLocateOnRestrictedElements:this.doNotLocateOnRestrictedElements});}}}if(this.directionsOutputType){switch(this.directionsOutputType.toLowerCase()){case "complete":_b.directionsOutputType="esriDOTComplete";break;case "complete-no-events":_b.directionsOutputType="esriDOTCompleteNoEvents";break;case "instructions-only":_b.directionsOutputType="esriDOTInstructionsOnly";break;case "standard":_b.directionsOutputType="esriDOTStandard";break;case "summary-only":_b.directionsOutputType="esriDOTSummaryOnly";break;default:_b.directionsOutputType=this.directionsOutputType;}}var _d=function(_e,_f){if(!_e){return null;}if(_e.declaredClass==="esri.tasks.FeatureSet"){if(_e.features.length>0){return _3.toJson({type:"features",features:_7._encodeGraphics(_e.features,_a&&_a[_f])});}else{return null;}}else{if(_e.declaredClass==="esri.tasks.DataLayer"){return _e;}else{if(_e.declaredClass==="esri.tasks.DataFile"){return _3.toJson({type:"features",url:_e.url});}}}return _3.toJson(_e);};if(this.barriers){_b.barriers=_d(this.barriers,"barriers.features");}if(this.polygonBarriers){_b.polygonBarriers=_d(this.polygonBarriers,"polygonBarriers.features");}if(this.polylineBarriers){_b.polylineBarriers=_d(this.polylineBarriers,"polylineBarriers.features");}return _6.filter(_b,function(_10){if(_10!==null){return true;}});}});if(_4("extend-esri")){_2.setObject("tasks.RouteParameters",_9,_5);}return _9;});