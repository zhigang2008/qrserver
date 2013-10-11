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
define("esri/layers/MosaicRule",["dojo/_base/declare","dojo/_base/lang","dojo/_base/json","dojo/has","esri/kernel","esri/lang","esri/geometry/Point"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.layers.MosaicRule",method:null,where:null,sortField:null,sortValue:null,ascending:false,lockRasterIds:null,viewpoint:null,objectIds:null,operation:null,constructor:function(_9){if(!_2.isObject(_9)){return;}_2.mixin(this,_9);if(_9.mosaicMethod){this.method=_9.mosaicMethod;}if(this.method&&this.method.toLowerCase().substring(0,4)!=="esri"){this.method=this._getMethodEnum(this.method);}if(_9.mosaicOperation){this.operation=_9.mosaicOperation;}if(this.operation&&this.operation.toUpperCase().substring(0,3)!=="MT_"){this.operation=this._getOperatorEnum(this.operation);}if(_9.fids){this.objectIds=_9.fids;}if(_9.viewpoint){this.viewpoint=new _7(_9.viewpoint);}},toJson:function(){var _a={mosaicMethod:this.method,where:this.where,sortField:this.sortField,sortValue:this.sortValue?_3.toJson(this.sortValue):null,ascending:this.ascending,lockRasterIds:this.lockRasterIds,viewpoint:this.viewpoint?this.viewpoint.toJson():null,fids:this.objectIds,mosaicOperation:this.operation};return _6.filter(_a,function(_b){if(_b!==null){return true;}});},_getMethodEnum:function(_c){if(!_c){return;}var _d=_8.METHOD_NONE;switch(_c.toLowerCase()){case "byattribute":_d=_8.METHOD_ATTRIBUTE;break;case "center":_d=_8.METHOD_CENTER;break;case "lockraster":_d=_8.METHOD_LOCKRASTER;break;case "nadir":_d=_8.METHOD_NADIR;break;case "northwest":_d=_8.METHOD_NORTHWEST;break;case "seamline":_d=_8.METHOD_SEAMLINE;break;case "viewpoint":_d=_8.METHOD_VIEWPOINT;break;}return _d;},_getOperatorEnum:function(_e){if(!_e){return;}switch(_e.toLowerCase()){case "first":return _8.OPERATION_FIRST;case "last":return _8.OPERATION_LAST;case "max":return _8.OPERATION_MAX;case "min":return _8.OPERATION_MIN;case "blend":return _8.OPERATION_BLEND;case "mean":return _8.OPERATION_MEAN;}}});_2.mixin(_8,{METHOD_NONE:"esriMosaicNone",METHOD_CENTER:"esriMosaicCenter",METHOD_NADIR:"esriMosaicNadir",METHOD_VIEWPOINT:"esriMosaicViewpoint",METHOD_ATTRIBUTE:"esriMosaicAttribute",METHOD_LOCKRASTER:"esriMosaicLockRaster",METHOD_NORTHWEST:"esriMosaicNorthwest",METHOD_SEAMLINE:"esriMosaicSeamline",OPERATION_FIRST:"MT_FIRST",OPERATION_LAST:"MT_LAST",OPERATION_MIN:"MT_MIN",OPERATION_MAX:"MT_MAX",OPERATION_MEAN:"MT_MEAN",OPERATION_BLEND:"MT_BLEND"});if(_4("extend-esri")){_2.setObject("layers.MosaicRule",_8,_5);}return _8;});