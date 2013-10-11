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
define("esri/layers/RenderMode",["dojo/_base/kernel","dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","dojo/io/script","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_2(null,{declaredClass:"esri.layers._RenderMode",constructor:function(){this._prefix="jsonp_"+(_1._scopeName||"dojo")+"IoScript";},initialize:function(_9){this.map=_9;this._init=true;},startup:function(){},propertyChangeHandler:function(_a){},destroy:function(){this._init=false;},drawFeature:function(_b){},suspend:function(){},resume:function(){},refresh:function(){},_incRefCount:function(_c){var _d=this._featureMap[_c];if(_d){_d._count++;}},_decRefCount:function(_e){var _f=this._featureMap[_e];if(_f){_f._count--;}},_getFeature:function(oid){return this._featureMap[oid];},_addFeatureIIf:function(oid,_10){var _11=this._featureMap,_12=_11[oid],_13=this.featureLayer;if(!_12){_11[oid]=_10;_13._add(_10);_10._count=0;}return _12||_10;},_removeFeatureIIf:function(oid){var _14=this._featureMap[oid],_15=this.featureLayer;if(_14){if(_14._count){return;}delete this._featureMap[oid];_15._remove(_14);}return _14;},_clearIIf:function(){var i,_16=this.featureLayer,_17=_16.graphics,_18=_16._selectedFeatures,_19=_16.objectIdField;for(i=_17.length-1;i>=0;i--){var _1a=_17[i];var oid=_1a.attributes[_19];if(oid in _18){_1a._count=1;continue;}_1a._count=0;this._removeFeatureIIf(oid);}},_isPending:function(id){var dfd=_6[this._prefix+id];return dfd?true:false;},_cancelPendingRequest:function(dfd,id){dfd=dfd||_6[this._prefix+id];if(dfd){try{dfd.cancel();_6._validCheck(dfd);}catch(e){}}},_purgeRequests:function(){_6._validCheck(null);},_toggleVisibility:function(_1b){var _1c=this.featureLayer,_1d=_1c.graphics,_1e=_1b?"show":"hide",i,len=_1d.length;_1b=_1b&&_1c._ager;for(i=0;i<len;i++){var _1f=_1d[i];_1f[_1e]();if(_1b){_1c._repaint(_1f);}}},_applyTimeFilter:function(_20){var _21=this.featureLayer;if(!_21.timeInfo||_21.suspended){return;}if(!_20){_21._fireUpdateStart();}var _22=_21._trackManager;if(_22){_22.clearTracks();}var _23=_21.getTimeDefinition(),_24=_21._getOffsettedTE(_21._mapTimeExtent);if(_24){_24=_21._getTimeOverlap(_23,_24);if(_24){var _25=_21._filterByTime(_21.graphics,_24.startTime,_24.endTime);if(_22){_22.addFeatures(_25.match);}_4.forEach(_25.match,function(_26){var _27=_26._shape;if(!_26.visible){_26.show();_27=_26._shape;_27&&_27._moveToFront();}if(_21._ager&&_27){_21._repaint(_26);}});_4.forEach(_25.noMatch,function(_28){if(_28.visible){_28.hide();}});}else{this._toggleVisibility(false);}}else{if(_22){_22.addFeatures(_21.graphics);}this._toggleVisibility(true);}if(_22){_22.moveLatestToFront();_22.drawTracks();}if(!_20){_21._fireUpdateEnd();}}});if(_5("extend-esri")){_3.setObject("layers._RenderMode",_8,_7);}return _8;});