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
define("esri/layers/ArcGISMapServiceLayer",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/has","esri/kernel","esri/lang","esri/request","esri/SpatialReference","esri/geometry/Extent","esri/layers/LayerInfo"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1(null,{declaredClass:"esri.layers.ArcGISMapServiceLayer",constructor:function(_c,_d){this.layerInfos=[];var _e=(this._params={}),_f=this._url.query?this._url.query.token:null;if(_f){_e.token=_f;}},_load:function(){_7({url:this._url.path,content:_2.mixin({f:"json"},this._params),callbackParamName:"callback",load:this._initLayer,error:this._errorHandler});},spatialReference:null,initialExtent:null,fullExtent:null,description:null,units:null,_initLayer:function(_10,io){try{this._findCredential();var ssl=(this.credential&&this.credential.ssl)||(_10&&_10._ssl);if(ssl){this._useSSL();}this.description=_10.description;this.copyright=_10.copyrightText;this.spatialReference=_10.spatialReference&&new _8(_10.spatialReference);this.initialExtent=_10.initialExtent&&new _9(_10.initialExtent);this.fullExtent=_10.fullExtent&&new _9(_10.fullExtent);this.units=_10.units;this.maxRecordCount=_10.maxRecordCount;this.maxImageHeight=_10.maxImageHeight;this.maxImageWidth=_10.maxImageWidth;this.supportsDynamicLayers=_10.supportsDynamicLayers;var _11=(this.layerInfos=[]),_12=_10.layers,dvl=(this._defaultVisibleLayers=[]);_3.forEach(_12,function(lyr,i){_11[i]=new _a(lyr);if(lyr.defaultVisibility){dvl.push(lyr.id);}});if(!this.visibleLayers){this.visibleLayers=dvl;}this.version=_10.currentVersion;if(!this.version){var ver;if("capabilities" in _10||"tables" in _10){ver=10;}else{if("supportedImageFormatTypes" in _10){ver=9.31;}else{ver=9.3;}}this.version=ver;}this.capabilities=_10.capabilities;if(_6.isDefined(_10.minScale)&&!this._hasMin){this.setMinScale(_10.minScale);}if(_6.isDefined(_10.maxScale)&&!this._hasMax){this.setMaxScale(_10.maxScale);}}catch(e){this._errorHandler(e);}}});if(_4("extend-esri")){_2.setObject("layers.ArcGISMapServiceLayer",_b,_5);}return _b;});