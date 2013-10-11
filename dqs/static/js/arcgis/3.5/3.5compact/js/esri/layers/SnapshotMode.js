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
define("esri/layers/SnapshotMode",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/SpatialReference","esri/tasks/query","esri/layers/RenderMode"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1([_7],{declaredClass:"esri.layers._SnapshotMode",constructor:function(_9){this.featureLayer=_9;this._featureMap={};this._drawFeatures=_2.hitch(this,this._drawFeatures);this._queryErrorHandler=_2.hitch(this,this._queryErrorHandler);},startup:function(){if(this.featureLayer._collection){this._applyTimeFilter();}else{this._fetchAll();}},propertyChangeHandler:function(_a){if(this._init){if(_a){if(this.featureLayer._collection){console.log("FeatureLayer: layer created by value (from a feature collection) does not support definition expressions and time definitions. Layer id = "+this.featureLayer.id);}else{this._fetchAll();}}else{this._applyTimeFilter();}}},drawFeature:function(_b){var _c=this.featureLayer,_d=_c.objectIdField,_e=_b.attributes[_d];this._addFeatureIIf(_e,_b);this._incRefCount(_e);},resume:function(){this.propertyChangeHandler(0);},refresh:function(){var _f=this.featureLayer;if(_f._collection){_f._fireUpdateStart();_f._refresh(true);_f._fireUpdateEnd();}else{this._fetchAll();}},_getRequestId:function(_10){var id="_"+_10.name+_10.layerId+_10._ulid;return id.replace(/[^a-zA-Z0-9\_]+/g,"_");},_fetchAll:function(){var _11=this.featureLayer;if(_11._collection){return;}_11._fireUpdateStart();this._clearIIf();this._sendRequest();},_sendRequest:function(){var map=this.map,_12=this.featureLayer,_13=_12.getDefinitionExpression();var _14=new _6();_14.outFields=_12.getOutFields();_14.where=_13||"1=1";_14.returnGeometry=true;_14.outSpatialReference=new _5(map.spatialReference.toJson());_14.timeExtent=_12.getTimeDefinition();_14.maxAllowableOffset=_12._maxOffset;if(_12._ts){_14._ts=(new Date()).getTime();}var _15;if(_12._usePatch){_15=this._getRequestId(_12);this._cancelPendingRequest(null,_15);}_12._task.execute(_14,this._drawFeatures,this._queryErrorHandler,_15);},_drawFeatures:function(_16){this._purgeRequests();var _17=_16.features,_18=this.featureLayer,_19=_18.objectIdField,i,len=_17.length,_1a,oid;for(i=0;i<len;i++){_1a=_17[i];oid=_1a.attributes[_19];this._addFeatureIIf(oid,_1a);this._incRefCount(oid);}this._applyTimeFilter(true);_18._fireUpdateEnd(null,_16.exceededTransferLimit?{queryLimitExceeded:true}:null);if(_16.exceededTransferLimit){_18.onQueryLimitExceeded();}},_queryErrorHandler:function(err){this._purgeRequests();var _1b=this.featureLayer;_1b._errorHandler(err);_1b._fireUpdateEnd(err);}});if(_3("extend-esri")){_2.setObject("layers._SnapshotMode",_8,_4);}return _8;});