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
define("esri/tasks/GPResultImageLayer",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","dojo/io-query","esri/kernel","esri/layers/ArcGISDynamicMapServiceLayer"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(_7,{declaredClass:"esri.tasks._GPResultImageLayer",constructor:function(_9,_a){if(_a&&_a.imageParameters&&_a.imageParameters.extent){this.initialExtent=(this.fullExtent=_a.imageParameters.extent);this.spatialReference=this.initialExtent.spatialReference;}this.getImageUrl=_2.hitch(this,this.getImageUrl);this.loaded=true;this.onLoad(this);},getImageUrl:function(_b,_c,_d,_e){var _f=this._url.path+"?",_10=this._params,sr=_b.spatialReference.wkid;_e(_f+_5.objectToQuery(_2.mixin(_10,{f:"image",bbox:_3.toJson(_b.toJson()),bboxSR:sr,imageSR:sr,size:_c+","+_d})));}});if(_4("extend-esri")){_2.setObject("tasks._GPResultImageLayer",_8,_6);}return _8;});