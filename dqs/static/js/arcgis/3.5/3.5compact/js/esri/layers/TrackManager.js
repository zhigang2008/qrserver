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
define("esri/layers/TrackManager",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/graphic","esri/geometry/Polyline","esri/layers/GraphicsLayer"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9=_1(null,{declaredClass:"esri.layers._TrackManager",constructor:function(_a){this.layer=_a;this.trackMap={};},initialize:function(_b){this.map=_b;var _c=this.layer,_d=_c.renderer.trackRenderer;if(_d&&(_c.geometryType==="esriGeometryPoint")){var _e=(this.container=new _8._GraphicsLayer({id:_c.id+"_tracks",_child:true}));_e.loaded=true;_e.onLoad(_e);_e._setMap(_b,_c._div);_e.setRenderer(_d);}},addFeatures:function(_f){var _10,_11=this.trackMap,_12=this.layer,_13=_12._trackIdField;_3.forEach(_f,function(_14){var _15=_14.attributes;_10=_15[_13];var ary=(_11[_10]=(_11[_10]||[]));ary.push(_14);});var _16=_12._startTimeField,_17=_12.objectIdField;var _18=function(a,b){var _19=a.attributes[_16],_1a=b.attributes[_16];if(_19===_1a){return (a.attributes[_17]<b.attributes[_17])?-1:1;}else{return (_19<_1a)?-1:1;}};for(_10 in _11){_11[_10].sort(_18);}},drawTracks:function(){var _1b=this.container;if(!_1b){return;}var _1c=this.trackMap,sr=this.map.spatialReference,_1d,ary,_1e,i,_1f,_20=this.layer._trackIdField,_21;for(_1d in _1c){ary=_1c[_1d];_1e=[];for(i=ary.length-1;i>=0;i--){_1f=ary[i].geometry;if(_1f){_1e.push([_1f.x,_1f.y]);}}_21={};_21[_20]=_1d;if(_1e.length>0){_1b.add(new _6(new _7({paths:[_1e],spatialReference:sr}),null,_21));}}},moveLatestToFront:function(){_3.forEach(this.getLatestObservations(),function(_22){var _23=_22._shape;_23&&_23._moveToFront();this._repaint(_22,null,true);},this.layer);},getLatestObservations:function(){var _24=[];if(!this.layer.renderer.latestObservationRenderer){return _24;}var _25=this.trackMap,_26;for(_26 in _25){var ary=_25[_26];_24.push(ary[ary.length-1]);}return _24;},clearTracks:function(){var _27=this.getLatestObservations();this.trackMap={};var _28=this.container;if(_28){_28.clear();}_3.forEach(_27,function(_29){this._repaint(_29,null,true);},this.layer);},isLatestObservation:function(_2a){var _2b=this.layer._trackIdField;var _2c=this.trackMap[_2a.attributes[_2b]];if(_2c){return (_2c[_2c.length-1]===_2a);}return false;},destroy:function(){var _2d=this.container;if(_2d){_2d.clear();_2d._unsetMap(this.map,this.layer._div);this.container=null;}this.map=null;this.layer=null;this.trackMap=null;}});if(_4("extend-esri")){_2.setObject("layers._TrackManager",_9,_5);}return _9;});