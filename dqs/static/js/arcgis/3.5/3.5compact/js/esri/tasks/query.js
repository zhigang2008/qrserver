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
define("esri/tasks/query",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/json","dojo/has","esri/kernel","esri/geometry/jsonUtils","esri/tasks/SpatialRelationship","dojo/has!extend-esri?esri/tasks/QueryTask","dojo/has!extend-esri?esri/tasks/RelationshipQuery","dojo/has!extend-esri?esri/tasks/StatisticDefinition"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.tasks.Query",constructor:function(){this.spatialRelationship=_9.SPATIAL_REL_INTERSECTS;},text:null,where:"",geometry:null,groupByFieldsForStatistics:null,objectIds:null,returnGeometry:false,orderByFields:null,outSpatialReference:null,outFields:null,outStatistics:null,timeExtent:null,relationParam:null,pixelSize:null,toJson:function(_a){var _b={text:this.text,where:this.where,returnGeometry:this.returnGeometry,spatialRel:this.spatialRelationship,maxAllowableOffset:this.maxAllowableOffset,geometryPrecision:this.geometryPrecision},g=_a&&_a["geometry"]||this.geometry,_c=this.objectIds,_d=this.outFields,_e=this.outSpatialReference,_f=this.groupByFieldsForStatistics,_10=this.orderByFields,_11=this.outStatistics;if(g){_b.geometry=g;_b.geometryType=_7.getJsonType(g);_b.inSR=g.spatialReference.wkid||_4.toJson(g.spatialReference.toJson());}if(_c){_b.objectIds=_c.join(",");}if(_d){_b.outFields=_d.join(",");}if(_f){_b.groupByFieldsForStatistics=_f.join(",");}if(_10){_b.orderByFields=_10.join(",");}if(_11){var _12=[];_3.forEach(_11,function(_13,idx){_12.push(_13.toJson());});_b.outStatistics=_4.toJson(_12);}if(_e!==null){_b.outSR=_e.wkid||_4.toJson(_e.toJson());}else{if(g){_b.outSR=g.spatialReference.wkid||_4.toJson(g.spatialReference.toJson());}}var _14=this.timeExtent;_b.time=_14?_14.toJson().join(","):null;var _15=this.relationParam;if(_15&&this.spatialRelationship===_9.SPATIAL_REL_RELATION){_b.relationParam=_15;}_b.pixelSize=this.pixelSize?_4.toJson(this.pixelSize.toJson()):null;_b._ts=this._ts;return _b;}});_2.mixin(_9,_8);if(_5("extend-esri")){_2.setObject("tasks.Query",_9,_6);}return _9;});